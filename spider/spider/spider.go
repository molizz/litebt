package spider

import (
	"encoding/hex"
	"fmt"

	"github.com/shiyanhui/dht"
)

type File struct {
	Path   []string `json:"path"`
	Length int      `json:"length"`
}

type BitTorrent struct {
	InfoHash string `json:"infohash"`
	Name     string `json:"name"`
	Files    []File `json:"files,omitempty"`
	Length   int    `json:"length,omitempty"`
}

func RunSpider(callback func(torrent *BitTorrent)) {
	//go func() {
	//	http.ListenAndServe(":6060", nil)
	//}()

	w := dht.NewWire(65536, 1024, 256)
	go func() {
		for resp := range w.Response() {
			metadata, err := dht.Decode(resp.MetadataInfo)
			if err != nil {
				continue
			}

			info := metadata.(map[string]interface{})

			if _, ok := info["name"]; !ok {
				continue
			}

			bt := &BitTorrent{
				InfoHash: hex.EncodeToString(resp.InfoHash),
				Name:     info["name"].(string),
			}

			if v, ok := info["files"]; ok {
				files := v.([]interface{})
				bt.Files = make([]File, len(files))

				for i, item := range files {
					f := item.(map[string]interface{})

					file := File{
						Length: f["length"].(int),
					}
					for _, path := range f["path"].([]interface{}) {
						file.Path = append(file.Path, fmt.Sprintf("%s", path))
					}
					bt.Files[i] = file
				}
			} else if _, ok := info["length"]; ok {
				bt.Length = info["length"].(int)
			}

			callback(bt)
		}
	}()
	go w.Run()

	config := dht.NewCrawlConfig()
	config.OnAnnouncePeer = func(infoHash, ip string, port int) {
		w.Request([]byte(infoHash), ip, port)
	}
	config.BlackListMaxSize = config.BlackListMaxSize * 10 // 会占用更多的内存
	config.MaxNodes = config.MaxNodes * 10
	d := dht.New(config)

	d.Run()
}
