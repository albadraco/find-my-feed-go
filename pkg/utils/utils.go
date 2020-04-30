package utils

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"	
	"strings"

	"github.com/albadraco/find-my-feed-go/pkg/types"
)

// SeasonOneExpression is this a season 1 episode.
var SeasonOneExpression = "(^.*)([Ss]01[Ee]01)(.*$)"

// SeasonExpressions is this a season expression.
var SeasonExpressions = []string{ 
	"(^.*)([Ss][0-9][0-9][Ee][0-9][0-9])(.*$)", 
	"(^.*)([0-9]x[0-9])|([0-9][0-9]x[0-9][0-9])(.*$)",
	"(^.*)([0-9]?[0-9]x[0-9]?[0-9])(.*$)",
}

// Alphabetic sort type
type Alphabetic []types.MyInterests

// Len length
func (list Alphabetic) Len() int { return len(list) }

// Swap  ya switch it.
func (list Alphabetic) Swap(i, j int) { list[i], list[j] = list[j], list[i] }

// Less is it less?
func (list Alphabetic) Less(i, j int) bool {
    var si string = list[i].Name
    var sj string = list[j].Name
    var silower = strings.ToLower(si)
    var sjlower = strings.ToLower(sj)
    if silower == sjlower {
        return si < sj
    }
    return silower < sjlower
}

// CollectInterested  collects from file system lists of path names that are shows that is interesting.
func CollectInterested( paths []string) (curInterests []types.MyInterests, err error) {
	for _, p := range paths {
		files, err := ioutil.ReadDir(p)
		if err != nil {
			fmt.Println("READ DIR ERROR: ", err)
		}
		for _, file := range files {
			if file.Mode().IsDir() {
				interest := types.MyInterests{
					Name: file.Name(),
					Path: p,
				}
				curInterests = append(curInterests, interest)
			}	
		}
	}
	sort.Sort(Alphabetic(curInterests))
	return curInterests, err
}

// HasSeason definition
func HasSeason(title string) ( Yes bool) {
	var seasonone = regexp.MustCompile(SeasonExpressions[0])
	var seasontwo = regexp.MustCompile(SeasonExpressions[1])
	var seasonthree = regexp.MustCompile(SeasonExpressions[2])
	return seasonone.MatchString(title) || seasontwo.MatchString(title) || seasonthree.MatchString(title)
}

// IsSeasonOne  is it a new season?
func IsSeasonOne(title string) ( Yes bool ) {
	var seasonone = regexp.MustCompile(SeasonOneExpression)
	return seasonone.MatchString(title)
}

// AmInterested  do i want this one or not
func AmInterested( title string, myinterests []types.MyInterests ) ( Yes bool, Details types.MyInterests ) {
	tl := strings.ToLower(title)

	Yes = false
	for _, interest := range myinterests {
		il := strings.ToLower(interest.Name)
		//fmt.Printf("AmInterested: %s -> %s\n", il, tl)
		if strings.Contains(tl, il) {
			Yes = true
			Details = interest
			break
		}
	}
	return Yes, Details
}