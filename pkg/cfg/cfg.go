package cfg

import (
	//"encoding/json"
	//"fmt"
	//"os"

	//"github.com/mmcdole/gofeed"
	
	"github.com/albadraco/find-my-feed-go/pkg/types"
	//"../../pkg/utils"
)

// Load  create configuration
func Load() *types.Myfeedconfig {
	return &defaultfeeds
}