package bot

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
)

func GetPuzzles(api anaconda.TwitterApi, max_id, since_id string) map[string]string {
	values := url.Values{}
	values.Set("max_id", string(max_id))
	values.Set("since_id", string(since_id))

	timeline, err := api.GetMentionsTimeline(values)
	if err != nil {
		fmt.Println(err)
	}
	return parseTimeline(timeline)
}
