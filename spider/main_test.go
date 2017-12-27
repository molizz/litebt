package main

import (
	"testing"

	"github.com/molisoft/litebt/spider/spider"
	"github.com/molisoft/litebt/web/model"
)

func TestCreateFile(t *testing.T) {
	files := []spider.File{{
		Path:   []string{"hello", "fuck"},
		Length: 1024,
	}}
	bt := spider.BitTorrent{
		InfoHash: "hash",
		Name:     "hhh",
		Files:    files,
		Length:   1024,
	}

	file, err := CreateFile(&bt)
	if err != nil {
		t.Error(err)
		return
	}
	if file.ID > 0 {
		t.Log("good!", file.ID)
		model.Db.Unscoped().Delete(file)
	}
}
