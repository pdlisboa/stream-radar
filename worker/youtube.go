package worker

import (
	"encoding/json"
	"fmt"
	"os"
	"stream-radar/internal/http"
)

type YTSearchResponse struct {
	Items []struct {
		ID struct {
			VideoID string `json:"videoId"`
		} `json:"id"`
		Snippet struct {
			LiveBroadcastContent string `json:"liveBroadcastContent"`
			Title                string `json:"title"`
			ChannelTitle         string `json:"channelTitle"`
		} `json:"snippet"`
	} `json:"items"`
}

var ytApiKey string

func searchLiveYt(channelName string) (YTSearchResponse, error) {
	ytApiKey = os.Getenv("YT_API_KEY")
	url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/search")
	resp, err := http.Get[YTSearchResponse](http.GetOptions{
		BaseOptions: http.BaseOptions{
			Url: url,
		},
		QueryParams: map[string]string{
			"part":      "snippet",
			"q":         channelName,
			"type":      "video",
			"eventType": "live",
			"key":       ytApiKey,
		},
	})
	return resp, err
}

func TestYt(channelName string) {
	respChannel, _ := searchLiveYt(channelName)
	json.NewEncoder(os.Stdout).Encode(respChannel)

	//fmt.Printf(a)
	//fmt.Println(err)
}
