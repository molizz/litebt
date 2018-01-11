package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/molisoft/litebt/lib"
)

func downloadHandler(ctx *gin.Context) {

}

func statusHandler(ctx *gin.Context) {

}

func cancelHandler(ctx *gin.Context) {

}

func RunHttp() {
	engine := gin.New()
	engine.POST("/download", downloadHandler)
	engine.POST("/status", statusHandler)
	engine.POST("/cancel", cancelHandler)

	engine.Run(fmt.Sprintf(":%d", lib.Cfg.Downloader.Port))
}
