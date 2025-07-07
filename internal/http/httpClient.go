package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/exp/maps"
	"io"
	"net/http"
	"net/url"
	"time"
)

type BaseOptions struct {
	Url     string
	Headers map[string]string
}

type GetOptions struct {
	BaseOptions
	QueryParams map[string]string
}

type PostOptions struct {
	BaseOptions
	QueryParams map[string]string
	Body        any
}

func applyHeaders(req *http.Request, headers map[string]string) {
	req.Header.Add("Accept", `application/json`)
	if headers != nil && len(headers) > 0 {
		for _, header := range maps.Keys(headers) {
			req.Header.Add(header, headers[header])
		}
	}
}

func generateBodyPost() {

}

func getClient() *http.Client {
	c := http.Client{Timeout: time.Duration(10) * time.Second}
	return &c
}

func Get[T any](options GetOptions) (T, error) {

	client := getClient()

	var respBody T

	urlString, _ := url.Parse(options.Url)
	params := url.Values{}
	for key, value := range options.QueryParams {
		params.Set(key, value)
	}
	urlString.RawQuery = params.Encode()

	req, err := http.NewRequest(http.MethodGet, urlString.String(), nil)
	if err != nil {
		fmt.Printf("error %s", err)
		return respBody, err
	}
	applyHeaders(req, options.Headers)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("error %s", err)
		return respBody, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("error %s", err)
		return respBody, err
	}

	err = json.Unmarshal(body, &respBody)

	if err != nil {
		fmt.Printf("error %s", err)
		return respBody, err
	}

	return respBody, err
}

func Post[T any](options PostOptions) (T, error) {
	var respBody T
	var bodyBytes []byte
	var err error

	contentType := options.Headers["Content-Type"]
	switch contentType {
	case "application/x-www-form-urlencoded":
		values := url.Values{}
		if mapBody, ok := options.Body.(map[string]string); ok {
			for key, value := range mapBody {
				values.Set(key, value)
			}
		} else if v, ok := options.Body.(url.Values); ok {
			values = v
		}
		bodyBytes = []byte(values.Encode())
	default:
		bodyBytes, err = json.Marshal(options.Body)
		if err != nil {
			return respBody, err
		}
	}

	payload := bytes.NewBuffer(bodyBytes)

	client := getClient()
	req, err := http.NewRequest(http.MethodPost, options.Url, payload)

	if err != nil {
		fmt.Printf("error %s", err)
		return respBody, err
	}
	applyHeaders(req, options.Headers)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("error %s", err)
		return respBody, err
	}
	defer resp.Body.Close()

	bodyResp, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("error %s", err)
		return respBody, err
	}

	err = json.Unmarshal(bodyResp, &respBody)

	if err != nil {
		fmt.Printf("error %s", err)
		return respBody, err
	}

	return respBody, err
}
