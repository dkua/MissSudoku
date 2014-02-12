package bot

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"strings"
)

type Puzzle struct {
	Id         int64
	Created    int64
	Updated    int64
	Puzzle     int64
	TweetId    string
	SolutionId int64
	ScreenName string
}

type Solution struct {
	Id       int64
	Created  int64
	Updated  int64
	Solution string
	PuzzleId int64
}

func parseTimeline(timeline []anaconda.Tweet) map[string][]string {
	fmt.Println(len(timeline))
	tweets := make(map[string][]string, 0)
	for _, tweet := range timeline {
		tweets[tweet.IdStr] = parseTweet(tweet)
	}
	return tweets
}

func parseTweet(tweet anaconda.Tweet) []string {
	parsed_tweet := make([]string, 2)
	parsed_tweet[0] = tweet.User.ScreenName
	tweet_array := strings.Fields(tweet.Text)
	for _, puzzle := range tweet_array {
		if len(puzzle) == 81 {
			parsed_tweet[1] = puzzle
		} else {
			parsed_tweet[1] = ""
		}
	}
	return parsed_tweet
}
