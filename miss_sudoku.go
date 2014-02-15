package main

import (
	"./bot"
	"fmt"
	"net/url"
	"time"
)

func main() {
	api := bot.GetTwitterApi()
	for {
		SINCE_ID := bot.GetSinceId(*api)
		puzzles := bot.GetPuzzles(*api, SINCE_ID)
		values := url.Values{}
		for id, text := range puzzles {
			values.Set("in_reply_to_status_id", string(id))
			tweet, err := api.PostTweet(text, values)
			if err != nil {
				fmt.Printf("%v\n%v", "Something went wrong", err)
			}
			fmt.Println(tweet)
		}
		time.Sleep(60 * time.Second)
	}
}
