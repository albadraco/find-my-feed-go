package utils

import (
	"io/ioutil"
	"regexp"
	"sort"	
	"strings"

	log "github.com/sirupsen/logrus"


	"github.com/albadraco/find-my-feed-go/pkg/types"
)

// SeasonOneExpression is this a season 1 episode.
var SeasonOneExpression = "(^.*)([Ss]01[Ee]01)(.*$)"

// SeasonExpressions is this a season expression.
	//"(^.*)([s][0-9][0-9][e][0-9][0-9])(.*$)", 
	//"(^.*)([0-9]x[0-9])|([0-9][0-9]x[0-9][0-9])(.*$)",
	//"(^.*)(\\d{1,2}x\\d{1,2})(.*$)",
var SeasonExpressions = []string{ 
	"(^.*)(s\\d+e\\d+)(.*$)", 
	"(^.*)(\\d+x\\d+)(.*$)",
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
			log.Errorf("READ DIR ERROR: ", err)
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

func getshow(re string, title string) (show string) {
	pattern := regexp.MustCompile(re)
	idx := pattern.FindAllSubmatchIndex([]byte(title), -1)
	log.Debugf("getshow: IDX: %v", idx)
	
	for _, loc := range idx {
		show = string(title[loc[2]:loc[3]])
	}

	return show
}

// HasSeason definition
func HasSeason(title string) ( Yes bool, Name string) {
	var seasonone = regexp.MustCompile(SeasonExpressions[0])
	var seasontwo = regexp.MustCompile(SeasonExpressions[1])
	Yes = false
	//Yes = seasonone.MatchString(title) || seasontwo.MatchString(title)
	if seasonone.MatchString(title) {
		Yes = true
		Name  = getshow(SeasonExpressions[0], title)
	} else if seasontwo.MatchString(title) {
		Yes = true
		Name  = getshow(SeasonExpressions[1], title)
	}
	return Yes, Name
}

// IsSeasonOne  is it a new season?
func IsSeasonOne(title string) ( Yes bool ) {
	var seasonone = regexp.MustCompile(SeasonOneExpression)
	return seasonone.MatchString(title)
}

// AmInterested  do i want this one or not
func AmInterested( title string, myinterests []types.MyInterests ) ( Yes bool, Details types.MyInterests ) {
	replacer := strings.NewReplacer(" ", "", ".", "", "(", "", ")", "", "[", "", "]", "")
	tl := strings.ToLower(title)
	_, tl = HasSeason(tl)
	tl = replacer.Replace(tl)

	Yes = false
	for _, interest := range myinterests {
		il := strings.ToLower(interest.Name)
		il = replacer.Replace(il)
		log.Debugf("AmInterested: %s -> %s", tl, il)
		if tl == il {
			log.Debugf(" **** FOUND ONE ****")
			Yes = true
			Details = interest
			break
		}
	}
	return Yes, Details
}
