package main

import (
	"NotionToGoogleCalender/api"
	"fmt"
	"time"
)

const (
	timeFormat = "2006-01-02"
)

// GetNotionCalender は、notionカレンダーから情報を取得して構造体を返します。
func GetNotionCalender(start, end string) ([]*AddCalender, error) {
	addCalenderList := []*AddCalender{}
	notionList := []*api.Result{}
	notion, err := api.New()
	if err != nil {
		return nil, err
	}

	rangeList := GetSplitRange(start, end, []*Range{})
	// 2ヶ月ごとにページのリストを取得
	for _, item := range rangeList {
		reqBody, err := api.NewCalenderReqBody(item.Start, item.End)
		if err != nil {
			return nil, err
		}
		res, err := notion.FetchPageList(reqBody)
		if err != nil {
			return nil, err
		}
		notionList = append(notionList, res.Results...)
	}

	// ページの子ブロックを取得して構造体に追加
	for _, item := range notionList {
		childblock, err := notion.FetchChildrenBlock(item.ID)
		if err != nil {
			return nil, err
		}
		// 子ブロックのリストを変換して追加
		addCalender := toAddCalender(childblock.Results, item.Properties.Date.Date.Start)
		addCalenderList = append(addCalenderList, addCalender)
	}
	return addCalenderList, nil
}

// GetSplitRangeは、取得したい日付範囲を2ヶ月ごとに分割して再帰的に取得します
func GetSplitRange(start, end string, list []*Range) []*Range {
	startDate, _ := time.Parse(timeFormat, start)
	endDate, _ := time.Parse(timeFormat, end)

	// 開始日より2ヶ月プラスした日時
	addStartDate := startDate.AddDate(0, 2, 0)
	//  開始日より2ヶ月プラスした日時 < 終了日となっているか
	ok := addStartDate.Before(endDate)
	if ok {
		addStart := addStartDate.Format(timeFormat)
		list = append(list, &Range{Start: start, End: addStart})
		list = GetSplitRange(addStart, end, list)
	} else {
		list = append(list, &Range{Start: start, End: end})
	}
	return list
}

// toAddCalender は、取得した子ブロックをgoogleカレンダー登録用の構造体に変換します。
func toAddCalender(list []*api.Result, date string) *AddCalender {
	result := &AddCalender{}

	// 追加する日付
	result.Date = date
	// タイトル
	dateTime, _ := time.Parse(timeFormat, date)
	result.Title = fmt.Sprintf("Diary %d-%d-%d", dateTime.Year(), dateTime.Month(), dateTime.Day())
	// 本文
	body := ""
	for _, item := range list {
		for _, text := range item.Paragraph.RichText {
			body += text.PlainText + "\n"
		}
	}
	result.Body = body

	return result
}
