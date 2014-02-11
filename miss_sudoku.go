package main

import (
	"configuration"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/dkua/go-sudoku"
	"net/url"
)

// Secrets
var CONSUMER_KEY = ""
var CONSUMER_TOKEN = ""
var ACCESS_KEY = ""
var ACCESS_TOKEN = ""

func main() {
	api := configuration.GetTwitterApi()
	puzzles := getPuzzles(*api)
	for id, puzzle := range puzzles {
		solved_puzzle := sudoku.Display(sudoku.Solve(puzzle))
		fmt.Println(id)
		fmt.Println(solved_puzzle)
	}
}

func getPuzzles(api anaconda.TwitterApi) map[string]string {
	puzzles := make(map[string]string, 0)

	values := url.Values{}
	values.Set("exclude_replies", "true")

	timeline, err := api.GetHomeTimeline(values)
	fmt.Println(timeline, err)

	return puzzles
}
