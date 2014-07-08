package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/dkua/MissSudoku/bot"
)

func main() {
	api := bot.GetTwitterApi()
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil) // Used to bind Heroku port
	for {
		since_id := bot.GetSinceId(*api)
		solutions := bot.GetSolutions(*api, since_id)
		values := url.Values{}
		for i := len(solutions) - 1; i >= 0; i-- {
			solution := solutions[i]
			values.Set("in_reply_to_status_id", solution[0])
			tweet, err := api.PostTweet(solution[1], values)
			if err != nil {
				fmt.Printf("%v\n%v", "Something went wrong", err)
			}
			fmt.Println(tweet)
		}
		time.Sleep(360 * time.Second)
	}
}
