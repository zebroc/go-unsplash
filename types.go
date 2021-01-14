package go_unsplash

// Photo represents a photo object
type Photo struct {
	ID             string            `json:"id"`
	CreatedAt      string            `json:"created_at"`
	UpdatedAt      string            `json:"updated_at"`
	Width          int               `json:"width"`
	Height         int               `json:"height"`
	Description    string            `json:"description"`
	AltDescription string            `json:"alt_escription"`
	Urls           map[string]string `json:"urls"`
	User           User              `json:"user"`
}

// User represents a user object
type User struct {
	ID           string            `json:"id"`
	CreatedAt    string            `json:"created_at"`
	UpdatedAt    string            `json:"updated_at"`
	Username     string            `json:"username"`
	Name         string            `json:"name"`
	Links        map[string]string `json:"links"`
	ProfileImage map[string]string `json:"profile_image"`
}

// Collection represents a collection object
type Collection struct {
	ID          string            `json:"id"`
	PublishedAt string            `json:"published_at"`
	UpdatedAt    string            `json:"updated_at"`
	Slug        string            `json:"slug"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Links       map[string]string `json:"links"`
}

// Topic represents a topic object
type Topic struct {
	ID        string `json:"id"`
	PublishedAt string            `json:"published_at"`
	UpdatedAt    string            `json:"updated_at"`
	Slug        string            `json:"slug"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Links       map[string]string `json:"links"`
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
