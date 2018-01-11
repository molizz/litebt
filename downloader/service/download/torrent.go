package download

import (
	"github.com/anacrolix/dht"
	"github.com/anacrolix/torrent"
	"github.com/kataras/go-errors"
	"time"
)

type ProgressFunc func(completed int64, total int64)

type Torrent struct {
	magnet  string
	client  *torrent.Client
	torrent *torrent.Torrent

	progressFunc ProgressFunc
}

func (this *Torrent) Start() (err error) {
	go func() {
		<-this.torrent.GotInfo()
		this.torrent.DownloadAll()
	}()

	go func() {
		<-this.torrent.GotInfo()
		for {
			completed := this.torrent.BytesCompleted()
			total := this.torrent.Info().TotalLength()

			if this.progressFunc != nil {
				this.progressFunc(completed, total)
			}
			if completed == total { // 下载完成
				return
			}
			time.Sleep(3 * time.Second)
		}
	}()

	if this.client.WaitAll() {
		return nil
	}
	return errors.New("download fail")
}

func (this *Torrent) AddMagnet(magnetUri string) (err error) {
	this.magnet = magnetUri
	this.torrent, err = this.client.AddMagnet(magnetUri)
	if err != nil {
		return err
	}
}

func NewTorrent() (*Torrent, error) {
	config := torrent.Config{
		DHTConfig: dht.ServerConfig{
			StartingNodes: dht.GlobalBootstrapAddrs,
		},
	}
	c, err := torrent.NewClient(&config)
	if err != nil {
		return nil, err
	}

	return &Torrent{
		client: c,
	}, nil
}
