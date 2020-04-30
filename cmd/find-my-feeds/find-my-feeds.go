package main

import (
	//"encoding/json"
	"fmt"
	//"os"

	"github.com/mmcdole/gofeed"

	"github.com/albadaco/find-my-feed-go/pkg/cfg"
	"github.com/albadaco/find-my-feed-go/pkg/types"
	"github.com/albadaco/find-my-feed-go/pkg/utils"
)

func main() {

	mycfg := cfg.Load()

	interestedIn, err := utils.CollectInterested(mycfg.DestinationPaths)
	if err != nil {
		fmt.Println("No Interests: ", err)
	}

	if mycfg.Debug {
		for _, dirname := range interestedIn {
			fmt.Println(dirname)
		}
	}

	var selections types.Myfeedselections
	getdata := true
	if getdata {
		fp := gofeed.NewParser()

		for _, feed := range mycfg.Feeds {
			fmt.Printf(">>>>> FEED: %s", feed.FeedURL)
			f, _ := fp.ParseURL(feed.FeedURL, mycfg.Header)
			fmt.Printf(" <<<<<< Count: %d\n", len(f.Items))
			if f.Items != nil {
				for _, entry := range f.Items {
					if mycfg.Debug {
						fmt.Printf("DEBUG: %v\n\n", entry)
					}
					yes, details := utils.AmInterested(entry.Title, interestedIn)
					if yes {
						if entry.Enclosures != nil && len(entry.Enclosures) > 0 {
							details.URL = entry.Enclosures[0].URL
							selections.Selected = append(selections.Selected, details)
						} else if len(entry.Link) > 0 {
							details.URL = entry.Link
							selections.Selected = append(selections.Selected, details)
						} else {
							selections.Skipped = append(selections.Skipped, entry.Title+"\n")
						}
					} else {
						selections.NoInterest = append(selections.NoInterest, entry.Title+"\n")
					}
					if utils.IsSeasonOne(entry.Title) {
						selections.S01E01 = append(selections.S01E01, entry.Title+"\n")
					}
					if !utils.HasSeason(entry.Title) {
						selections.NoSeason = append(selections.NoSeason, entry.Title+"\n")
					}
				}
			} else {
				fmt.Println("***** Nothing returned.")
				fmt.Println(f)
			}
		}

		fmt.Println("----------------------")
		fmt.Printf("Selected...[%3d]:\n", len(selections.Selected))
		for _, sel := range selections.Selected {
			fmt.Printf("\tsel.Name...: %s\n", sel.Name)
			fmt.Printf("\tsel.Path...: %s\n", sel.Path)
			fmt.Printf("\tsel.URL....: %s\n", sel.URL)
		}
		fmt.Printf("Skipped....[%3d]: %v\n\n", len(selections.Skipped), selections.Skipped)
		fmt.Printf("Unknown....[%3d]: %v\n\n", len(selections.Unknown), selections.Unknown)
		fmt.Printf("NoSeason...[%3d]: %v\n\n", len(selections.NoSeason), selections.NoSeason)
		fmt.Printf("NoInterest.[%3d]: %v\n\n", len(selections.NoInterest), selections.NoInterest)
		fmt.Printf("S01E01.....[%3d]: %v\n\n", len(selections.S01E01), selections.S01E01)
		fmt.Println("----------------------")
	}
}
