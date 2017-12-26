package model

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

type SubFile struct {
	Path   []string
	Length int
}

type File struct {
	gorm.Model

	Hash     string `gorm:"size:40;unique_index"`
	Name     string
	SubFiles string `gorm:"type:text"`
	Length   int
}

func NewFile(hash string, name string, files []SubFile, length int) *File {
	subFilesRaw, err := json.Marshal(files)
	if err != nil {
		return nil
	}
	return &File{
		Hash:     hash,
		Name:     name,
		SubFiles: string(subFilesRaw),
		Length:   length,
	}
}
