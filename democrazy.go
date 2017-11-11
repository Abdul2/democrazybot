package main

import (
"fmt"
"github.com/dghubble/go-twitter/twitter"
"github.com/dghubble/oauth1"
"os"
)

func main() {

	consumerKey := os.Getenv("consumerKey")
	consumerSecret := os.Getenv("consumerSecret")
	accessToken := os.Getenv("accessToken")
	accessSecret := os.Getenv("accessSecret")

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Search tweets to retweet
	searchParams := &twitter.SearchTweetParams{
		Query:      "#Puidgemont ",
		Count:      5,
		ResultType: "recent",
		Lang:       "es",
	}

	searchResult, _, _ := client.Search.Tweets(searchParams)

	// Retweet
	for _, tweet := range searchResult.Statuses {
		tweet_id := tweet.ID
		client.Statuses.Retweet(tweet_id, &twitter.StatusRetweetParams{})

		t := tweet.Text + "IT is a democrazy"

		fmt.Printf("RETWEETED: %+v\n", t)
	}
}