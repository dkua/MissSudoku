package bot

import (
	"os"

	"github.com/ChimeraCoder/anaconda"
)

type Configuration struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

func GetTwitterApi() *anaconda.TwitterApi {
	var config Configuration
	config.ConsumerKey = os.Getenv("CONSUMER_KEY")
	config.ConsumerSecret = os.Getenv("CONSUMER_SECRET")
	config.AccessToken = os.Getenv("ACCESS_TOKEN")
	config.AccessSecret = os.Getenv("ACCESS_SECRET")

	anaconda.SetConsumerKey(config.ConsumerKey)
	anaconda.SetConsumerSecret(config.ConsumerSecret)
	api := anaconda.NewTwitterApi(config.AccessToken, config.AccessSecret)

	return api
}
