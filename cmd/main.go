package main

import (
	"fmt"

	"github.com/MaximilianHagelstam/hncli/news"
)

func main() {
	ids, err := news.GetTopPostÍds()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ids)
}
