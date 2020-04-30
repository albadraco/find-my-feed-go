package types

import (
	"github.com/mmcdole/gofeed"
)

// Myfeedconfig a struct
type Myfeedconfig struct {
	Debug            bool          `json:"debug,omitempty"`
	Feeds            []Myfeedinfo  `json:"feeds,omitempty"`
	DestinationPaths []string      `json:"destinations,omitempty"`
	Header           gofeed.Header `json:"header,omitempty"`
}

// MyInterests a struct
type MyInterests struct {
	URL  string `json:"url,omitempty"`
	Name string `json:"name,omitempty"`
	Path string `json:"path,omitempty"`
}

// Myshow a struct
type Myshow struct {
	Name string
	Last string
	Size int
}

// Myfeedinfo a struct
type Myfeedinfo struct {
	Etag     string `json:"etag,omitempty"`
	FeedType string `json:"type,omitempty"`
	FeedURL  string `json:"url,omitempty"`
}

// Myfeedselections a struct
type Myfeedselections struct {
	Selected   []MyInterests `json:"selected,omitempty"`
	Skipped    []string      `json:"skipped,omitempty"`
	Unknown    []string      `json:"unknown,omitempty"`
	NoSeason   []string      `json:"noseason,omitempty"`
	NoInterest []string      `json:"nointerest,omitempty"`
	S01E01     []string      `json:"s01e01,omitempty"`
}
