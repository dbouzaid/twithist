package hist

import (
	"encoding/json"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/dbouzaid/twithist/constants"
	"github.com/dbouzaid/twithist/util"
	"net/http"
	"net/url"
	"time"
)

// TODO: Comment Code

func LoadHist(w http.ResponseWriter, req *http.Request) {
	// Setup the Anaconda Library for the Twitter API with keys and tokens
	anaconda.SetConsumerKey(constants.ConsumerKey)
	anaconda.SetConsumerSecret(constants.ConsumerSecret)
	api := anaconda.NewTwitterApi(constants.AccessToken, constants.AccessTokenSecret)

	//
	path := req.URL.Path[1:]
	name := util.GetSlicedPathAt(1, path)

	//
	currentTime := time.Now().UTC()
	timeFrame := currentTime.Add(-24 * time.Hour)
	// Find the user
	user, err := api.GetUsersLookup(name, nil)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		queries := createQueries(user[0])
		tweets, _ := api.GetUserTimeline(queries)
		validTweets := getPastTweets(tweets, timeFrame)
		tweetMap := mapTweets(validTweets, timeFrame)
		twitHist := TwitterHistogram{TweetsTimeMap: tweetMap}
		jsonHash, _ := json.Marshal(twitHist)
		w.Write(jsonHash)
	}
}

func getPastTweets(tweets []anaconda.Tweet, timeFrame time.Time) (validTweets []anaconda.Tweet) {
	for _, tweet := range tweets {
		tweetTime, err := tweet.CreatedAtTime()
		if err != nil {
			fmt.Println("Error: ", err)
		} else if tweetTime.After(timeFrame) {
			validTweets = append(validTweets, tweet)
		}
	}
	return
}

func mapTweets(tweets []anaconda.Tweet, timeFrame time.Time) (tweetMap map[int]int) {
	tweetMap = initialiseMap()
	for _, tweet := range tweets {
		tweetTime, err := tweet.CreatedAtTime()
		if err != nil {
			fmt.Println("Error: ", err)
		} else {
			key := int(tweetTime.Sub(timeFrame).Hours())
			tweetMap[key]++
		}
	}
	return
}

func initialiseMap() (m map[int]int) {
	m = make(map[int]int)
	for i := 0; i <= 24; i++ {
		m[i] = 0
	}
	return
}

func createQueries(user anaconda.User) (urlVals url.Values) {
	urlVals = url.Values{}
	urlVals.Add(constants.UserId, user.IdStr)
	urlVals.Add(constants.Count, "200")
	urlVals.Add(constants.ExcludeReplies, "true")
	urlVals.Add(constants.IncludeRts, "false")
	return
}

type TwitterHistogram struct {
	TweetsTimeMap map[int]int `json:"tweets"`
}
