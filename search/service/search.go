package service

import (
	"github.com/go-ego/riot"
	"github.com/go-ego/riot/types"
	"github.com/molisoft/litebt/lib/utils"
)

var (
	searcher *riot.Engine
)

// 添加索引
//
func AddIndex(content string, index uint64) {
	data := types.DocIndexData{Content: content}
	searcher.IndexDoc(index, data)
}

// 搜索
//
func Search(key string, page int, max int) types.SearchResp {
	page = utils.Max(page, 1)               // 至少为1
	max = utils.Min(utils.Max(max, 1), 100) // 至少为1，且不能大于100

	req := types.SearchReq{
		Text: key,
		RankOpts: &types.RankOpts{
			OutputOffset: page * max,
			MaxOutputs:   max,
		},
	}
	return searcher.Search(req)
}

// 立即刷新缓存（阻塞至完成）
//
func Flush() {
	searcher.FlushIndex()
}

func RunSearcher() {
	searcher = riot.New()
}
