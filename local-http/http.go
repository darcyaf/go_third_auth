package local_http

import (
	"context"
	"io/ioutil"
	"net/http"
)
import url "net/url"

func Get(ctx context.Context, rawurl string, params url.Values) ([]byte, error) {
	url2, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	url2.RawQuery = params.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", url2.String(), nil)
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}