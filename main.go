package main

const (
	// startとendに指定された間のカレンダーを登録する
	Start = "2021-01-01"
	End   = "2022-01-01"
)

func main() {
	// notionの情報を取得
	// notionList := GetNotionCalender(Start, End)

	// 取得した情報をgoogleカレンダーに登録
	// err := PutGoogleCalender(notionList)
	// if err != nil {
	// 	log.Fatalf("Failed PutGoogleCalender: %v", err)
	// }
}
