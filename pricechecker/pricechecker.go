package pricechecker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type AppStore struct{}

type RequireInfo struct {
	ID      string
	Country string
}

type appStoreResponse struct {
	Count   float64         `json:"resultCount"`
	Results []RegisteredApp `json:"results"`
}

type RegisteredApp struct {
	AppUrl      string  `json:"trackViewUrl"`
	Name        string  `json:"trackName"`
	ID          int64   `json:"trackId"`
	Price       float32 `json:"price"`
	Rating      float32 `json:"averageUserRatingForCurrentVersion"`
	ReleaseDate string  `json:"currentVersionReleaseDate"`
}

type AppStoreError struct {
	ErrMsg string
}

func (ase AppStoreError) Error() string {
	return fmt.Sprintf("Error Message : %v", ase.ErrMsg)
}

func (a AppStore) buildLookupURL(r RequireInfo) (string, error) {
	u, _ := url.Parse("https://itunes.apple.com/lookup")
	q := u.Query()

	if r.ID == "" {
		return "", AppStoreError{ErrMsg: "No ID presented"}
	}

	q.Set("id", r.ID)

	if r.Country != "" {
		q.Set("country", r.Country)
	} else {
		q.Set("country", "kr")
	}

	u.RawQuery = q.Encode()

	return u.String(), nil
}

func (a AppStore) LookupApp(r RequireInfo) (RegisteredApp, error) {
	url, err := a.buildLookupURL(r)
	if err != nil {
		return RegisteredApp{}, err
	}

	result, err := http.Get(url)

	if err != nil {
		return RegisteredApp{}, err
	}

	defer result.Body.Close()

	response, err := ioutil.ReadAll(result.Body)

	if err != nil {
		return RegisteredApp{}, err
	}

	var appResults appStoreResponse

	err = json.Unmarshal(response, &appResults)

	if err != nil {
		return RegisteredApp{}, err
	}

	if len(appResults.Results) > 0 {
		return appResults.Results[0], nil
	}
	return RegisteredApp{}, AppStoreError{ErrMsg: "No results for the id"}
}
