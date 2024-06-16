package news

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type Post struct {
	ID    int
	Title string
	URL   string
	By    string
	Score int
}

func GetTopPostÍds() (IDs *[]int, err error) {
	url := "https://hacker-news.firebaseio.com/v0/topstories.json"

	client := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	postIds := []int{}
	if err := json.Unmarshal(body, &postIds); err != nil {
		return nil, err
	}

	return &postIds, nil
}
