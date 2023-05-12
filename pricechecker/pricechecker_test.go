package pricechecker

import (
	"strconv"
	"testing"
)

func TestBuildLookupURL(t *testing.T) {
	r := AppQuery{ID: "782438457"}
	url, err := r.buildLookupURL()

	if err != nil {
		t.Error(err)
	}

	if url != "https://itunes.apple.com/lookup?country=kr&id=782438457" {
		t.Error("No match for url")
	}

	r = AppQuery{ID: "782438457", Country: "us"}

	url, err = r.buildLookupURL()

	if err != nil {
		t.Error(err)
	}

	if url != "https://itunes.apple.com/lookup?country=us&id=782438457" {
		t.Error("No match for url")
	}
}

func TestLookupApp(t *testing.T) {
	// 정상
	r := AppQuery{ID: "782438457"}
	ra, err := r.LookupApp()

	if err != nil {
		t.Error(err)
	}

	if ra.Name != "Hoplite" {
		t.Errorf("Name is not matched [%v]", ra.Name)
	}

	if strconv.FormatInt(ra.ID, 10) != r.ID {
		t.Errorf("ID is not matched [%v]", ra.ID)
	}

	// 없는 country코드
	r = AppQuery{ID: "782438457", Country: "en"}
	_, err = r.LookupApp()

	if err == nil {
		t.Error("Should be Error")
	}

	t.Log(err)

	// 잘못된 ID
	r = AppQuery{ID: "7824384571"}
	_, err = r.LookupApp()

	if err == nil {
		t.Error("Should be Error")
	}

	t.Log(err)
}
