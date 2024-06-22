package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/MaximilianHagelstam/hncli/client"
)

func main() {
	postClient := client.New(&http.Client{Timeout: time.Second * 2})

	ids, err := postClient.GetTopPostÍDs()
	if err != nil {
		fmt.Println(err)
	}

	posts := []client.Post{}
	topTenIDs := (*ids)[:10]

	for _, id := range topTenIDs {
		post, err := postClient.GetPostByID(id)
		if err != nil {
			fmt.Println(err)
		}

		posts = append(posts, *post)
	}

	fmt.Printf("%+v\n", posts)
}
