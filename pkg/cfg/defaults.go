package cfg

import (
	//"encoding/json"
	//"fmt"
	//"os"

	"github.com/mmcdole/gofeed"
	
	"github.com/albadraco/find-my-feed-go/pkg/types"
	//"../../pkg/utils"
)

// Defaultfeeds  a set of working defaults from windows.
var defaultfeeds = types.Myfeedconfig {
	Debug: false,
	Feeds: []types.Myfeedinfo {
		{
			Etag: "None",
			FeedType: "RSS",
			FeedURL: "https://eztv.io/ezrss.xml",
		},
		{
			Etag: "None",
			FeedType: "ATOM",
			FeedURL: "https://rarbg.to/rssdd_magnet.php?categories=41",
		},
		{
			Etag: "None",
			FeedType: "ATOM",
			FeedURL: "http://showrss.info/other/all.rss",
		},
	},
	DestinationPaths: []string{
		"\\\\bluemoon\\storage\\Media\\upstairs\\tvshows",
		"\\\\bluemoon\\storage\\Media\\TVShows",
		//"/mnt/raid6/Media/upstairs/tvshows",
		//"/mnt/raid6/Media/TVShows",
	},
	Header: gofeed.Header{
		Name: "If-None-Match",
		Value: "None",
	},
}
