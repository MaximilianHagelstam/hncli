package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
	By    string `json:"by"`
	Score int    `json:"score"`
}

const (
	GetPostIDsUrl  = "https://hacker-news.firebaseio.com/v0/topstories.json"
	GetPostByIDUrl = "https://hacker-news.firebaseio.com/v0/item/%d.json"
)

type Client struct {
	client *http.Client
}

func New(c *http.Client) *Client {
	return &Client{client: c}
}

func (c *Client) GetTopPostÍDs() (*[]int, error) {
	req, err := http.NewRequest(http.MethodGet, GetPostIDsUrl, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(req)
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

	postIDs := []int{}
	if err := json.Unmarshal(body, &postIDs); err != nil {
		return nil, err
	}

	return &postIDs, nil
}

func (c *Client) GetPostByID(ID int) (*Post, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(GetPostByIDUrl, ID), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(req)
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

	post := Post{}
	if err := json.Unmarshal(body, &post); err != nil {
		return nil, err
	}

	return &post, nil
}
