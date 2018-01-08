package service

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/molisoft/litebt/lib"
)

func SearchHandler(ctx *gin.Context) {
	key := ctx.Param("key")
	page, err := strconv.Atoi(ctx.Param("page"))
	if err != nil {
		page = 1
	}
	max, err := strconv.Atoi(ctx.Param("max"))
	if err != nil {
		max = 20
	}
	resp := Search(key, page, max)

	ctx.JSON(200, gin.H{
		"status": "ok",
		"result": resp.Docs,
	})
}

func AddIndexHandler(ctx *gin.Context) {
	content := ctx.Param("content")
	index, err := strconv.Atoi(ctx.Param("index"))
	if err != nil {
		ctx.JSON(200, gin.H{"status": "fail"})
		return
	}
	AddIndex(content, uint64(index))
	ctx.JSON(200, gin.H{"status": "ok"})
}

func RunHttp() error {
	engine := gin.Default()
	engine.POST("/search", SearchHandler)
	engine.POST("/add_index", AddIndexHandler)
	err := engine.Run(fmt.Sprintf(":%d", lib.Cfg.Searcher.Port))
	if err != nil {
		fmt.Println(err)
	}
}
