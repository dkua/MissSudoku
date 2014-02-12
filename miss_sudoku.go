package main

import (
	"bot"
	"fmt"
	"net/url"
)

func main() {
	api := bot.GetTwitterApi()
	puzzles := bot.GetPuzzles(*api)
  values := url.Values{}
	for _, puzzle := range puzzles {
    values.Set("in_reply_to_status_id", puzzle[0])
    tweet, err := api.PostTweet(puzzle[1], values)
    if (err != nil) {
        fmt.Printf("%v\n%v", "Something went wrong", err)
    }
    fmt.Println(tweet)
	}
}
