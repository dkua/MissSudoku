package bot

import (
	"bytes"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/dkua/go-sudoku"
	"regexp"
	"strconv"
	"strings"
)

func parseTimeline(timeline []anaconda.Tweet, tweets [][]string) [][]string {
	for _, tweet := range timeline {
		tweets = append(tweets, []string{strconv.FormatInt(tweet.Id, 10), parseTweet(tweet)})
	}
	return tweets
}

func parseTweet(tweet anaconda.Tweet) string {
	text, err := parseText(tweet.Text)
	if err != nil {
		return fmt.Sprintf("%v @%v", err, tweet.User.ScreenName)
	} else {
		return fmt.Sprintf("%v@%v", text, tweet.User.ScreenName)
	}
}

func parseText(tweet_text string) (string, error) {

	// Build puzzle from tweet, puzzle is either lines of 9 units separated by newlines
	// or a single 81 char string
	var buffer bytes.Buffer
	text_array := strings.Fields(tweet_text)
	regex := regexp.MustCompile(`[0-9]*\.*`)
	for _, text := range text_array {
		// Check validity of the text
		text = strings.Join(regex.FindAllString(text, -1), "")
		if len(text) == 9 || len(text) == 81 {
			buffer.WriteString(text)
		}
	}
	puzzle := buffer.String()

	if len(puzzle) == 81 {
		solution, err := sudoku.Solve(puzzle)
		if err != nil {
			return "", fmt.Errorf("Sorry couldn't solve this puzzle.")
		}
		return sudoku.Display(solution, err), nil
	}
	return "", fmt.Errorf("Sorry couldn't parse this tweet")
}
