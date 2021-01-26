package go_unsplash

import (
	"fmt"
	"net/http"
	"net/url"
)

// PhotoResult represents the response of a photo search
type PhotoResult struct {
	Total      int     `json:"total"`
	TotalPages int     `json:"total_pages"`
	Results    []Photo `json:"results"`
}

// Photo represents a photo object
type Photo struct {
	ID             string `json:"id"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
	Description    string `json:"description"`
	AltDescription string `json:"alt_description"`
	Urls           Urls   `json:"urls"`
	Links          Links  `json:"links"`
	User           User   `json:"user"`
}

type PhotosCollection struct {
	client *Client
	url    string
}

func (col *PhotosCollection) Search(params interface{}) ([]Photo, *http.Response, []byte, error) {
	var r PhotoResult
	col.client.req.URL, _ = url.Parse(fmt.Sprintf("%s%s%s", col.client.baseURL, "/search/photos", params))
	col.client.req.Method = "GET"
	resp, body, err := col.client.executeAPICall(&r)
	return r.Results, resp, body, err
}

// UserResult represents the response of a user search
type UserResult struct {
	Total      int    `json:"total"`
	TotalPages int    `json:"total_pages"`
	Results    []User `json:"results"`
}

// User represents a user object
type User struct {
	ID                string            `json:"id"`
	CreatedAt         string            `json:"created_at"`
	UpdatedAt         string            `json:"updated_at"`
	Username          string            `json:"username"`
	Name              string            `json:"name"`
	Location          string            `json:"location"`
	PortfolioURL      string            `json:"portfolio_url"`
	Bio               string            `json:"bio"`
	InstagramUsername string            `json:"instagram_username"`
	Links             Links             `json:"links"`
	ProfileImage      map[string]string `json:"profile_image"`
}

type UsersCollection struct {
	client *Client
	url    string
}

func (col *UsersCollection) Search(params interface{}) ([]User, *http.Response, []byte, error) {
	var r UserResult
	col.client.req.URL, _ = url.Parse(fmt.Sprintf("%s%s%s", col.client.baseURL, "/search/users", params))
	col.client.req.Method = "GET"
	resp, body, err := col.client.executeAPICall(&r)
	return r.Results, resp, body, err
}

// CollectionResult represents the response of a collection search
type CollectionResult struct {
	Total      int          `json:"total"`
	TotalPages int          `json:"total_pages"`
	Results    []Collection `json:"results"`
}

// Collection represents a collection object
type Collection struct {
	ID          string            `json:"id"`
	PublishedAt string            `json:"published_at"`
	UpdatedAt   string            `json:"updated_at"`
	Slug        string            `json:"slug"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Links       map[string]string `json:"links"`
}

type CollectionsCollection struct {
	client *Client
	url    string
}

func (col *CollectionsCollection) Search(params interface{}) ([]Collection, *http.Response, []byte, error) {
	var r CollectionResult
	col.client.req.URL, _ = url.Parse(fmt.Sprintf("%s%s%s", col.client.baseURL, "/search/collections", params))
	col.client.req.Method = "GET"
	resp, body, err := col.client.executeAPICall(&r)
	return r.Results, resp, body, err
}

// Topic represents a topic object
type Topic struct {
	ID          string            `json:"id"`
	PublishedAt string            `json:"published_at"`
	UpdatedAt   string            `json:"updated_at"`
	Slug        string            `json:"slug"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Links       map[string]string `json:"links"`
}

type TopicsCollection struct {
	client *Client
	url    string
}

// Stats represents a stats object
type Stats struct {
	Photos             int `json:"photos"`
	Downloads          int `json:"downloads"`
	Views              int `json:"views"`
	Likes              int `json:"likes"`
	Photographers      int `json:"photographers"`
	Pixels             int `json:"pixels"`
	DownloadsPerSecond int `json:"downloads_per_second"`
	ViewsPerSecond     int `json:"views_per_second"`
	Developers         int `json:"developers"`
	Applications       int `json:"applications"`
	Requests           int `json:"requests"`
}

type Urls struct {
	Raw     string `json:"raw"`
	Full    string `json:"full"`
	Regular string `json:"regular"`
	Small   string `json:"small"`
	Thumb   string `json:"thumb"`
}

type Links struct {
	Self      string `json:"self"`
	Html      string `json:"html"`
	Photos    string `json:"photos"`
	Likes     string `json:"likes"`
	Portfolio string `json:"portfolio"`
	Following string `json:"following"`
	Followers string `json:"followers"`
}
