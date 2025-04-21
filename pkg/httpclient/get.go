package httpclient

import (
	"encoding/json"
	"net/http"
)

func NillableGet[Res any](reqURL string) (*Res, error) {
	println(reqURL)
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

func Get[Res any](reqURL string) (Res, error) {
	println(reqURL)
	var result Res
	resp, err := http.Get(reqURL)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return result, err
	}

	return result, nil
}
