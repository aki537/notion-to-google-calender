package controller

import "testing"

func TestPutGoogleCalender(t *testing.T) {
	addCalenderList := []*AddCalender{
		{
			Date:  "2022-03-21",
			Title: "テスト20220321",
			Body:  "テスト更新1",
		},
		{
			Date:  "2022-03-22",
			Title: "テスト20220322",
			Body:  "テスト更新3333333333",
		},
	}
	err := PutGoogleCalender(addCalenderList)
	if err != nil {
		t.Fatal(err)
	}
}
