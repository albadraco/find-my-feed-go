package main

import (
	//"encoding/json"
	//"fmt"
	//"os"

	log "github.com/sirupsen/logrus"
	"github.com/mmcdole/gofeed"

	"github.com/albadraco/find-my-feed-go/pkg/cfg"
	"github.com/albadraco/find-my-feed-go/pkg/types"
	"github.com/albadraco/find-my-feed-go/pkg/utils"
)

func main() {

	mycfg := cfg.Load()

	interestedIn, err := utils.CollectInterested(mycfg.DestinationPaths)
	if err != nil {
		log.Println("No Interests: ", err)
	}

	if mycfg.Debug {
		for _, dirname := range interestedIn {
			log.Println(dirname)
		}
	}
	
	var selections types.Myfeedselections
	getdata := true
	if getdata {
		fp := gofeed.NewParser()

		for _, feed := range mycfg.Feeds {
			log.Warnf(">>>>> FEED: %s", feed.FeedURL)
			f, _ := fp.ParseURL(feed.FeedURL, mycfg.Header)
			log.Warnf(" <<<<<< Count: %d", len(f.Items))
			if f.Items != nil {
				for _, entry := range f.Items {
					log.Debugf("DEBUG: %v", entry)
					yes, details := utils.AmInterested(entry.Title, interestedIn)
					if yes {
						if entry.Enclosures != nil && len(entry.Enclosures) > 0 {
							details.URL = entry.Enclosures[0].URL
							selections.Selected = append(selections.Selected, details)
						} else if len(entry.Link) > 0 {
							details.URL = entry.Link
							selections.Selected = append(selections.Selected, details)
						} else {
							selections.Skipped = append(selections.Skipped, entry.Title)
						}
					} else {
						selections.NoInterest = append(selections.NoInterest, entry.Title)
					}
					if utils.IsSeasonOne(entry.Title) {
						selections.S01E01 = append(selections.S01E01, entry.Title)
					}
					if Ok, _ := utils.HasSeason(entry.Title); Ok {
						selections.NoSeason = append(selections.NoSeason, entry.Title)
					}
				}
			} else {
				log.Println("***** Nothing returned.")
				log.Println(f)
			}
		}

		log.Println("----------------------")
		log.Printf("Selected...[%3d]:", len(selections.Selected))
		for _, sel := range selections.Selected {
			log.Printf("sel.Name...: %s", sel.Name)
			log.Printf("sel.Path...: %s", sel.Path)
			log.Printf("sel.URL....: %s", sel.URL)
		}
		if mycfg.Debug {
			for i, sk := range selections.Skipped {
				log.Printf("Skipped....[%3d]: %s", i, sk)
			}
			for i, sk := range selections.Unknown {
				log.Printf("Unknown....[%3d]: %s", i, sk)
			}
			for i, sk := range selections.NoSeason {
				log.Printf("NoSeason...[%3d]: %s", i, sk)
			}
			for i, sk := range selections.NoInterest {
				log.Printf("NoInterest.[%3d]: %s", i, sk)
			}
		}
		for i, s := range selections.S01E01 {
			log.Printf("S01E01.....[%3d]: %s", i, s)
		}
		log.Println("----------------------")
	}
}
