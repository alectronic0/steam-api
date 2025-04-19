package httpclient

import (
	"encoding/json"
	"net/http"
)

func Get[Res any](reqURL string) (*Res, error) {
	resp, err := http.Get(reqURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result Res
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
