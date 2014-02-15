package main

import (
	"fmt"
	"github.com/dkua/MissSudoku/bot"
	"net/url"
	"time"
)

func main() {
	for {
		api := bot.GetTwitterApi()
		puzzles := bot.GetPuzzles(*api)
		values := url.Values{}
		for id, text := range puzzles {
			values.Set("in_reply_to_status_id", id)
			tweet, err := api.PostTweet(text, values)
			if err != nil {
				fmt.Printf("%v\n%v", "Something went wrong", err)
			}
			fmt.Println(tweet)
		}
		time.Sleep(60 * time.Second)
	}
}
