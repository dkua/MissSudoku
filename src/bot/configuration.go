package bot

import (
	"encoding/json"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"os"
)

type Configuration struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

func GetTwitterApi() *anaconda.TwitterApi {
	file, _ := os.Open("settings.json")
	decoder := json.NewDecoder(file)
	settings := &Configuration{}
	decoder.Decode(&settings)
	fmt.Println(settings.AccessToken)
	fmt.Println(settings.AccessSecret)

  anaconda.SetConsumerKey(settings.ConsumerKey)
  anaconda.SetConsumerSecret(settings.ConsumerSecret)
	api := anaconda.NewTwitterApi(settings.AccessToken, settings.AccessSecret)

	return api
}
