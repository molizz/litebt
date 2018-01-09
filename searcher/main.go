package searcher

import (
	"fmt"

	"github.com/molisoft/litebt/searcher/service"
)

func main() {

	service.RunSearcher()
	err := service.RunHttp()
	if err != nil {
		fmt.Println(err)
	}
}
