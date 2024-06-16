package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type Post struct {
	ID    int
	Title string
	URL   string
	By    string
	Score int
}

const (
	GetPostsUrl    = "https://hacker-news.firebaseio.com/v0/topstories.json"
	GetPostByIDUrl = "https://hacker-news.firebaseio.com/v0/item/%s.json"
)

type Api struct {
	client *http.Client
}

func New(c *http.Client) *Api {
	return &Api{client: c}
}

func (a *Api) GetTopPostÍds() (*[]int, error) {
	req, err := http.NewRequest(http.MethodGet, GetPostsUrl, nil)
	if err != nil {
		return nil, err
	}

	res, err := a.client.Do(req)
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
