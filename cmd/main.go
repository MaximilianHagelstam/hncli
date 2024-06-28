package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/MaximilianHagelstam/hncli/client"
)

const PageLength = 10

func main() {
	postClient := client.New(&http.Client{Timeout: time.Second * 2})

	ids, err := postClient.GetTopPostÍDs()
	if err != nil {
		fmt.Println(err)
	}

	topIDs := (*ids)[:PageLength]
	c := make(chan client.Post, len(topIDs))

	for _, id := range topIDs {
		go func() {
			post, _ := postClient.GetPostByID(id)
			c <- *post
		}()
	}

	posts := []client.Post{}
	for range topIDs {
		posts = append(posts, <-c)
	}

	fmt.Printf("%+v\n", posts)
}
