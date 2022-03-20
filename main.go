package main

import (
	"NotionToGoogleCalender/controller"
	"log"
)

const (
	// startとendに指定された間のカレンダーを登録する
	Start = "2021-10-01"
	End   = "2022-03-21"
)

func main() {
	// notionの情報を取得
	notionList, err := controller.GetNotionCalender(Start, End)
	if err != nil {
		log.Fatalf("Failed GetNotionCalender: %v", err)
	}

	// 取得した情報をgoogleカレンダーに登録
	err = controller.PutGoogleCalender(notionList)
	if err != nil {
		log.Fatalf("Failed PutGoogleCalender: %v", err)
	}
}
