package worker

import (
	"encoding/json"
	"fmt"
	"os"
	"stream-radar/internal/http"
	"stream-radar/worker/types"
)

var (
	kickClientId     string
	kickClientSecret string
)

func getKickToken() (types.TokenResponse, error) {
	kickClientId = os.Getenv("KICK_CLIENT_ID")
	kickClientSecret = os.Getenv("KICK_CLIENT_SECRET")
	resp, err := http.Post[types.TokenResponse](http.PostOptions{
		BaseOptions: http.BaseOptions{
			Url: fmt.Sprint("https://id.kick.com/oauth/token"),
			Headers: map[string]string{
				"Content-Type": "application/x-www-form-urlencoded",
			},
		},
		QueryParams: nil,
		Body: map[string]string{
			"client_id":     kickClientId,
			"client_secret": kickClientSecret,
			"grant_type":    "client_credentials",
		},
	})

	return resp, err
}

func searchKickLive(channelName string, token string) (SearchChannelResposne, error) {
	kickClientId = os.Getenv("KICK_CLIENT_ID")
	kickClientSecret = os.Getenv("KICK_CLIENT_SECRET")
	resp, err := http.Get[SearchChannelResposne](http.GetOptions{
		BaseOptions: http.BaseOptions{
			Url: "https://api.kick.com/public/v1/channels",
			Headers: map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", token),
				"Client-Id":     kickClientId,
			},
		},
		QueryParams: map[string]string{
			"slug": channelName,
		},
	})
	return resp, err
}

func TestKick(channelName string) {
	resp, err1 := getKickToken()

	if err1 != nil {
		return
	}
	respChannel, _ := searchKickLive(channelName, resp.AccessToken)

	json.NewEncoder(os.Stdout).Encode(respChannel)
	//fmt.Printf("%v", respChannel.Data)
	//fmt.Println(err)
}
