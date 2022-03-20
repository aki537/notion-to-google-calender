package main

import (
	"fmt"
	"log"
)

const (
	// startとendに指定された間のカレンダーを登録する
	Start = "2021-01-01"
	End   = "2022-01-01"
)

func main() {
	// notionの情報を取得
	notionList, err := GetNotionCalender(Start, End)
	if err != nil {
		log.Fatalf("Failed GetNotionCalender: %v", err)
	}
	fmt.Println(notionList)

	// 取得した情報をgoogleカレンダーに登録
	// err := PutGoogleCalender(notionList)
	// if err != nil {
	// 	log.Fatalf("Failed PutGoogleCalender: %v", err)
	// }
}
