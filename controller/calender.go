package controller

import (
	"NotionToGoogleCalender/api"
	"log"

	"google.golang.org/api/calendar/v3"
)

// PutGoogleCalender は、notionのカレンダーをgooglecalenderに登録します
// 登録済みの場合は更新します。
func PutGoogleCalender(addCalenderList []*AddCalender) error {
	calender, err := api.NewGoogleCalendar()
	if err != nil {
		return err
	}

	// notionのカレンダーをgooglecalenderに登録
	for _, item := range addCalenderList {
		log.Printf("カレンダー登録日：%s", item.Date)
		event := &calendar.Event{
			Summary:     item.Title,
			Description: item.Body,
			Start: &calendar.EventDateTime{
				Date: item.Date,
			},
			End: &calendar.EventDateTime{
				Date: item.Date,
			},
			ColorId: "2",
			Reminders: &calendar.EventReminders{
				UseDefault: false,
			},
		}
		// その日のイベントを取得し、同じタイトルがある場合は更新、無い場合は新しく挿入する
		dateList, err := calender.FetchDateEvents(item.Date)
		if err != nil {
			return err
		}
		eventId, ok := getEventID(dateList.Items, item.Title)
		if ok {
			log.Printf("カレンダー挿入 %s", event.Summary)
			err = calender.UpdateEvent(eventId, event)
			if err != nil {
				return err
			}
		} else {
			log.Printf("カレンダー更新 %s", event.Summary)
			err = calender.InsertEvent(event)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// getEventID は、引数に渡したタイトルがイベント内にある場合にeventIDを返却します
func getEventID(eventItems []*calendar.Event, addTitle string) (string, bool) {
	for _, item := range eventItems {
		if item.Summary == addTitle {
			return item.Id, true
		}
	}
	return "", false
}
