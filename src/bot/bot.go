package bot

import (
	"bytes"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/dkua/go-sudoku"
	"net/url"
)

func GetPuzzles(api anaconda.TwitterApi) [][]string {
	puzzles := make([][]string, 0)

	values := url.Values{}
	//	values.Set("max_id", string(max_id))
	//	values.Set("since_id", string(since_id))

	timeline, err := api.GetMentionsTimeline(values)
	if err != nil {
		fmt.Println(err)
	}
	tweets := parseTimeline(timeline)
	for id, tweet := range tweets {
		result := make([]string, 2)
		screenname := tweet[0]
		text := tweet[1]
		result[0] = id
		if text == "" {
			result[1] = "Sorry couldn't parse your tweet."
		} else {
			var buffer bytes.Buffer
			buffer.WriteString(fmt.Sprintf("%v@%v", sudoku.Display(sudoku.Solve(text)), screenname))
			result[1] = buffer.String()
		}
		puzzles = append(puzzles, result)
	}
	return puzzles
}
