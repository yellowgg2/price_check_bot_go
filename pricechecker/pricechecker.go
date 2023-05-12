package pricechecker

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type AppQuery struct {
	ID      string
	Country string
}

type appStoreResponse struct {
	Count   float64         `json:"resultCount"`
	Results []RegisteredApp `json:"results"`
	Err     string          `json:"errorMessage"`
}

type RegisteredApp struct {
	AppUrl      string  `json:"trackViewUrl"`
	Name        string  `json:"trackName"`
	ID          int64   `json:"trackId"`
	Price       float32 `json:"price"`
	Rating      float32 `json:"averageUserRatingForCurrentVersion"`
	ReleaseDate string  `json:"currentVersionReleaseDate"`
}

func (r AppQuery) buildLookupURL() (string, error) {
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

func (r AppQuery) LookupApp() (RegisteredApp, error) {
	url, err := r.buildLookupURL()
	if err != nil {
		return RegisteredApp{}, err
	}

	result, err := http.Get(url)

	if err != nil {
		return RegisteredApp{}, err
	}

	defer result.Body.Close()

	response, err := io.ReadAll(result.Body)

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

	errMsg := "No results for the id"
	if appResults.Err != "" {
		errMsg = appResults.Err
	}

	return RegisteredApp{}, AppStoreError{ErrMsg: errMsg, ID: r.ID}
}
