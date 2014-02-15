package bot

import (
	"bytes"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/dkua/go-sudoku"
	"regexp"
	"strings"
)

func parseTimeline(timeline []anaconda.Tweet) map[string]string {
	tweets := make(map[string]string, 0)
	for _, tweet := range timeline {
		tweets[tweet.IdStr] = parseTweet(tweet)
	}
	return tweets
}

func parseTweet(tweet anaconda.Tweet) string {
	text, err := parseText(tweet.Text)
	if err != nil {
		return fmt.Sprintf("%v", err)
	} else {
		var buffer bytes.Buffer
		buffer.WriteString(fmt.Sprintf("%v@%v", text, tweet.User.ScreenName))
		return buffer.String()
	}
}

func parseText(tweet_text string) (string, error) {

	// Build puzzle from tweet, puzzle is either lines of 9 units separated by newlines
	// or a single 81 char string
	var buffer bytes.Buffer
	text_array := strings.Fields(tweet_text)
	for _, text := range text_array {
		if len(text) == 9 || len(text) == 81 {
			buffer.WriteString(text)
		}
	}

	// Check validity of the puzzle and return, else return error
	regex := regexp.MustCompile(`[0-9]*\.*`)
	puzzle := strings.Join(regex.FindAllString(buffer.String(), -1), "")
	if len(puzzle) == 81 {
		solution, err := sudoku.Solve(puzzle)
		if err != nil {
			return "", fmt.Errorf("Sorry couldn't solve this puzzle.")
		}
		return sudoku.Display(solution, err), nil
	}
	return "", fmt.Errorf("Sorry couldn't parse this tweet")
}
