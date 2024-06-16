package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/MaximilianHagelstam/hncli/api"
)

func main() {
	client := http.Client{
		Timeout: time.Second * 2,
	}

	api := api.New(&client)

	ids, err := api.GetTopPostÍds()
	if err != nil {
		fmt.Println(err)
	}

	post, err := api.GetPostByID((*ids)[1])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(post.Title, post.URL)
}
