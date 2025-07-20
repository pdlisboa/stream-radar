package worker

import (
	"encoding/json"
	"fmt"
	"os"
	"stream-radar/internal/http"
	"stream-radar/worker/types"
)

type SearchChannelResposne struct {
	Data       []interface{} `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

var (
	twClientId     string
	twClientSecret string
)

func getTwToken() (types.TokenResponse, error) {
	twClientId = os.Getenv("TW_CLIENT_ID")
	twClientSecret = os.Getenv("TW_CLIENT_SECRET")
	resp, err := http.Post[types.TokenResponse](http.PostOptions{
		BaseOptions: http.BaseOptions{
			Url: fmt.Sprint("https://id.twitch.tv/oauth2/token?"),
			Headers: map[string]string{
				"Content-Type": "application/x-www-form-urlencoded",
			},
		},
		QueryParams: nil,
		Body: map[string]string{
			"client_id":     twClientId,
			"client_secret": twClientSecret,
			"grant_type":    "client_credentials",
		},
	})

	return resp, err
}

func searchChannel(channelName string, token string) (SearchChannelResposne, error) {
	twClientId = os.Getenv("TW_CLIENT_ID")
	twClientSecret = os.Getenv("TW_CLIENT_SECRET")
	resp, err := http.Get[SearchChannelResposne](http.GetOptions{
		BaseOptions: http.BaseOptions{
			Url: fmt.Sprint("https://api.twitch.tv/helix/search/channels"),
			Headers: map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", token),
				"Client-Id":     twClientId,
			},
		},
		QueryParams: map[string]string{
			"query": channelName,
		},
	})
	return resp, err
}

func TestTwitch(channelName string) {
	resp, err1 := getTwToken()

	if err1 != nil {
		return
	}
	respChannel, _ := searchChannel(channelName, resp.AccessToken)

	json.NewEncoder(os.Stdout).Encode(respChannel)
	//fmt.Printf("%v", respChannel)
	//fmt.Println(err)
}
