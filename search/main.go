package search

import (
	"fmt"

	"github.com/molisoft/litebt/search/service"
)

func main() {

	service.RunSearcher()
	err := service.RunHttp()
	if err != nil {
		fmt.Println(err)
	}
}
