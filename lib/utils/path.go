package utils

import (
	"os"
	"path"
)

// 返回当前程序的所在目录
func CurrentPath() string {
	return path.Dir(os.Args[0])
}
