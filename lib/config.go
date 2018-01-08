package lib

import (
	"fmt"
	"github.com/kataras/go-errors"
	"os"

	"github.com/jinzhu/configor"
	"github.com/molisoft/litebt/lib/utils"
)

type Config struct {
	Searcher struct {
		Port int `json:"port"`
	}
	Db struct {
		DbName   string `json:"dbname"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
}

var Cfg *Config

func init() {
	if Cfg != nil {
		return
	}
	Cfg = new(Config)

	var configPath = configPathFormDir(utils.CurrentPath())
	if _, err := os.Stat(configPath); err != nil {
		configPath = configPathFormDir(os.Getenv("PWD"))
	}

	err := configor.Load(Cfg, configPath)
	if err != nil {
		panic(errors.New("没有找到配置文件config.yml （" + configPath + ")"))
	}
}

func configPathFormDir(dir string) string {
	return fmt.Sprintf("%s/config/config.yml", dir)
}
