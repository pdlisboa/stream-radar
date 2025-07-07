package worker

import (
	"fmt"
	"os"
	"stream-radar/internal/http"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

type SearchChannelResposne struct {
	Data       []interface{} `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

var (
	clientId     string
	clientSecret string
)

func getToken() (TokenResponse, error) {
	clientId = os.Getenv("TW_CLIENT_ID")
	clientSecret = os.Getenv("TW_CLIENT_SECRET")
	resp, err := http.Post[TokenResponse](http.PostOptions{
		BaseOptions: http.BaseOptions{
			Url: fmt.Sprint("https://id.twitch.tv/oauth2/token?"),
			Headers: map[string]string{
				"Content-Type": "application/x-www-form-urlencoded",
			},
		},
		QueryParams: nil,
		Body: map[string]string{
			"client_id":     clientId,
			"client_secret": clientSecret,
			"grant_type":    "client_credentials",
		},
	})

	return resp, err
}

func searchChannel(channelName string, token string) (SearchChannelResposne, error) {
	clientId = os.Getenv("TW_CLIENT_ID")
	clientSecret = os.Getenv("TW_CLIENT_SECRET")
	resp, err := http.Get[SearchChannelResposne](http.GetOptions{
		BaseOptions: http.BaseOptions{
			Url: fmt.Sprint("https://api.twitch.tv/helix/search/channels"),
			Headers: map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", token),
				"Client-Id":     clientId,
			},
		},
		QueryParams: map[string]string{
			"query": channelName,
		},
	})
	return resp, err
}

func Test() {
	resp, err1 := getToken()

	if err1 != nil {
		return
	}
	fmt.Println(resp.AccessToken)
	respChannel, err := searchChannel("minerva", resp.AccessToken)

	fmt.Println(respChannel)
	fmt.Println(err)
}
