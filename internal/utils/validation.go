package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

const TheCatAPIURL = "https://api.thecatapi.com/v1/breeds/search?name="

type Breed struct {
	Name string `json:"name"`
}

func ValidateBreed(breed string) error {
	res, err := http.Get(TheCatAPIURL + url.QueryEscape(breed))
	if err != nil {
		return fmt.Errorf("failed to fetch breed data: %v", err)
	}
	defer res.Body.Close()

	var breeds []map[string]interface{}
	if err = json.NewDecoder(res.Body).Decode(&breeds); err != nil {
		return fmt.Errorf("failed to decode breed data: %v", err)
	}

	if len(breeds) == 0 {
		return errors.New("invalid breed")
	}

	return nil
}
