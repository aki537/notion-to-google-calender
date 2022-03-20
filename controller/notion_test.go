package controller

import (
	"fmt"
	"testing"
)

func TestGetNotionCalender(t *testing.T) {
	result, err := GetNotionCalender("2021-11-01", "2022-06-30")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}
