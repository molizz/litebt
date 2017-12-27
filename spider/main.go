package main

import (
	"fmt"

	"github.com/molisoft/litebt/spider/spider"
	"github.com/molisoft/litebt/web/model"
)

func CreateFile(t *spider.BitTorrent) (*model.File, error) {
	var files []model.SubFile
	for _, f := range t.Files {
		files = append(files, model.SubFile{f.Path, f.Length})
	}
	file := model.NewFile(t.InfoHash, t.Name, files, t.Length)

	db := model.Db.Create(file)
	if db.Error != nil {
		return nil, db.Error
	}
	return file, nil
}

func main() {

	// 存储到数据库
	newTorrent := func(t *spider.BitTorrent) {
		file, err := CreateFile(t)
		if err != nil {
			fmt.Println("[error ] create file err ", err)
			return
		}
		fmt.Println("[success] create file #", file.ID)
	}

	fmt.Println("Start spider...")

	go spider.RunSpider(newTorrent)
	select {}
}
