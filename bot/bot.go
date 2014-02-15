package bot

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
)

func GetPuzzles(api anaconda.TwitterApi, since_id int64) map[int64]string {
	values := url.Values{}
	values.Set("since_id", string(since_id))
	values.Set("count", "50")

	var max_id int64 = 0
	tweets := make(map[int64]string)
	for true {
		timeline, err := api.GetMentionsTimeline(values)
		if err != nil {
			fmt.Println(err)
		}
		if len(timeline) == 0 {
			break
		}

		for id, text := range parseTimeline(timeline) {
			tweets[id] = text
		}
		if max_id == 0 {
			max_id = timeline[len(timeline)-1].Id - 1
			values.Set("max_id", string(max_id))
		}
	}
	return tweets
}

func GetSinceId(api anaconda.TwitterApi) int64 {
	values := url.Values{}
	values.Set("screen_name", "misssudoku")
	values.Set("count", "1")

	user_timeline, err := api.GetUserTimeline(values)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	last_tweet := user_timeline[0]
	return last_tweet.InReplyToStatusID
}
