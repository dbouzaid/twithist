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

// LoadHist is used when the user first loads the histogram endpoint of the web application.
// Writes out a response to the user with a JSON hash of the number of tweets they made in the
// past 24 hours, all split up based on how many tweets per hour within the 24 hours were tweeted.
func LoadHist(w http.ResponseWriter, req *http.Request) {
	// Setup the Anaconda Library for the Twitter API with keys and tokens
	anaconda.SetConsumerKey(constants.ConsumerKey)
	anaconda.SetConsumerSecret(constants.ConsumerSecret)
	api := anaconda.NewTwitterApi(constants.AccessToken, constants.AccessTokenSecret)

	// Retrieve the current path and slice it to retrieve the twitter username to search for
	path := req.URL.Path[1:]
	name := util.GetSlicedPathAt(1, path)

	// Retrieve the current time in UTC
	currentTime := time.Now().UTC()

	// Create the time frame of 24 hours
	timeFrame := currentTime.Add(-24 * time.Hour)

	// Find the user based on named supplied by the endpoint
	user, err := api.GetUsersLookup(name, nil)

	// Check if a user was found
	if err != nil {
		// No user found with the name provided at the end point
		w.Write([]byte("No user found for specified name"))
	} else {
		// Create the queries to retrieve tweets by the user
		queries := createQueries(user[0])

		// Use the queries to retrieve tweets by the user
		tweets, _ := api.GetUserTimeline(queries)

		// Get the valid tweets within the given time frame
		validTweets := getValidTweets(tweets, timeFrame)

		// Create a map of the valid tweets
		tweetMap := mapTweets(validTweets, timeFrame)

		// Store the map in the struct ready to be marshaled
		twitHist := twitterHistogram{TweetsTimeMap: tweetMap}

		// Create the JSON hash to be returned to user
		jsonHash, _ := json.Marshal(twitHist)

		// Write the response with the JSON hash to the user
		w.Write(jsonHash)
	}
}


// getValidTweets is used to 
func getValidTweets(tweets []anaconda.Tweet, timeFrame time.Time) (validTweets []anaconda.Tweet) {
	for _, tweet := range tweets {
		tweetTime, err := tweet.CreatedAtTime()
		if err != nil {
			fmt.Println("Error1: ", err)
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
			fmt.Println("Error2: ", err)
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
	urlVals.Add(constants.UserID, user.IdStr)
	urlVals.Add(constants.Count, "200")
	urlVals.Add(constants.ExcludeReplies, "true")
	urlVals.Add(constants.IncludeRts, "false")
	return
}

type twitterHistogram struct {
	TweetsTimeMap map[int]int `json:"tweets"`
}
