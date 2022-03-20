package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

const (
	calendarId = "primary"
	timeFormat = "2006-01-02"
)

type GoogleCalendar struct {
	Srv *calendar.Service
}

// NewGoogleCalendar は,GoogleCalendarと接続する構造体を返却
func NewGoogleCalendar() (*GoogleCalendar, error) {
	ctx := context.Background()
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		return nil, err
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		return nil, err
	}
	client := getClient(config)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, err
	}

	return &GoogleCalendar{Srv: srv}, nil
}

// FetchDateEvents は、指定した日付に登録されている予定を取得します
func (g *GoogleCalendar) FetchDateEvents(date string) (*calendar.Events, error) {
	dateTime, _ := time.Parse(timeFormat, date)
	dateTimeMin := dateTime.AddDate(0, 0, -1)
	dateTimeMax := dateTime.AddDate(0, 0, 1)
	dateMinStr := dateTimeMin.Format(time.RFC3339)
	dateMaxStr := dateTimeMax.Format(time.RFC3339)

	events, err := g.Srv.Events.List(calendarId).ShowDeleted(false).
		SingleEvents(true).TimeMin(dateMinStr).TimeMax(dateMaxStr).OrderBy("startTime").Do()
	if err != nil {
		return nil, err
	}
	return events, nil
}

// InsertEvent は、GoogleCalendarにイベントを追加します
func (g *GoogleCalendar) InsertEvent(event *calendar.Event) error {
	_, err := g.Srv.Events.Insert(calendarId, event).Do()
	if err != nil {
		return err
	}
	return nil
}

// UpdateEvent は、GoogleCalendarに登録済みのイベントを更新します
func (g *GoogleCalendar) UpdateEvent(eventId string, event *calendar.Event) error {
	_, err := g.Srv.Events.Update(calendarId, eventId, event).Do()
	if err != nil {
		return err
	}
	return nil
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
