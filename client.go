package go_unsplash

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Options struct {
	BaseAPIURL string
	AccessKey  string
	SecretKey  string
}

type Client struct {
	http    *http.Client
	req     *http.Request
	options *Options
	baseURL string
}

func NewClient(options *Options) *Client {
	r := http.Request{}
	r.Header = make(map[string][]string)
	r.Header.Add("Authorization", "Client-ID "+options.AccessKey)

	return &Client{
		http:    &http.Client{Timeout: time.Second * 10},
		req:     &r,
		options: options,
		baseURL: options.BaseAPIURL,
	}
}

func (client *Client) Photos() *PhotosCollection {
	return &PhotosCollection{
		client: client,
		url:    fmt.Sprintf("%v/%v", client.baseURL, "search/photos"),
	}
}

func (client *Client) Collections() *CollectionsCollection {
	return &CollectionsCollection{
		client: client,
		url:    fmt.Sprintf("%v/%v", client.baseURL, "search/collections"),
	}
}

func (client *Client) Topics() *TopicsCollection {
	return &TopicsCollection{
		client: client,
		url:    fmt.Sprintf("%v/%v", client.baseURL, "topics"),
	}
}

func (client *Client) Users() *UsersCollection {
	return &UsersCollection{
		client: client,
		url:    fmt.Sprintf("%v/%v", client.baseURL, "search/users"),
	}
}

func (client *Client) executeAPICall(result interface{}) (*http.Response, []byte, error) {
	response, errGet := client.http.Do(client.req)
	if errGet != nil {
		return nil, nil, errGet
	}

	b, errRead := ioutil.ReadAll(response.Body)
	if errRead != nil {
		return nil, nil, errRead
	}

	errUnm := json.Unmarshal(b, result)
	if errUnm != nil {
		return nil, nil, errUnm
	}

	return response, b, nil
}
