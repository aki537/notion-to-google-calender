package api

import (
	"fmt"
	"testing"
)

func TestNewGoogleCalendar(t *testing.T) {
	a, err := NewGoogleCalendar()
	if err != nil {
		t.Fatal(err)
	}

	b, err := a.FetchDateEvents("2022-03-23")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(b)
}
