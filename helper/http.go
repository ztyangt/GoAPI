package helper

import (
	"io"
	"net/http"

	"github.com/spf13/cast"
)

type CurlResult struct {
	Data       any
	Body       string
	StatusCode int
	Header     http.Header
	Err        error
}

func GET(url string, params map[string]any, headers map[string]any) (result CurlResult) {

	request, _ := http.NewRequest("GET", url, nil)

	if params != nil {
		query := request.URL.Query()
		for key, val := range params {
			query.Add(key, cast.ToString(val))
		}
		request.URL.RawQuery = query.Encode()
	}

	if headers != nil {
		for key, val := range headers {
			request.Header.Set(key, cast.ToString(val))
		}
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		result.Err = err
		return
	}

	result.StatusCode = response.StatusCode
	result.Header = response.Header

	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			result.Err = err
			return
		}
		result.Body = string(body)
		result.Data = Json.Decode(string(body))
	}

	return
}
