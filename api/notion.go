package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const (
	baseURL = "https://api.notion.com/v1"
)

type Notion struct {
	DatabaseID string
	SecretKey  string
}

// 環境変数をセットしたNotion構造体を返却
func NewNotion() (*Notion, error) {
	err := godotenv.Load("./.env")
	if err != nil {
		return nil, err
	}
	secretKey := os.Getenv("NOTION_SECRET_KEY")
	databaseId := os.Getenv("NOTION_DATABASE_ID")

	return &Notion{SecretKey: secretKey, DatabaseID: databaseId}, nil
}

// ヘッダをセットしたリクエストを返却
func (n *Notion) NewRequest(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Notion-Version", "2022-02-22")
	req.Header.Set("Authorization", n.SecretKey)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

// FetchPageList は、ページのリストを取得
func (n *Notion) FetchPageList(reqBody io.Reader) (*ResList, error) {
	ReqURL := fmt.Sprintf("%s/databases/%s/query", baseURL, n.DatabaseID)
	req, err := n.NewRequest("POST", ReqURL, reqBody)
	if err != nil {
		return nil, err
	}
	result, err := FetchResToJSON(req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FetchChildrenBlock は、ページの子ブロックを取得
func (n *Notion) FetchChildrenBlock(pageID string) (*ResList, error) {
	ReqURL := fmt.Sprintf("%s/blocks/%s/children", baseURL, pageID)
	req, err := n.NewRequest("GET", ReqURL, nil)
	if err != nil {
		return nil, err
	}
	result, err := FetchResToJSON(req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// apiから取得したレスポンスをJSONにunmarshalして返却
func FetchResToJSON(req *http.Request) (*ResList, error) {
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	res := ResList{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// カレンダー用のリクエストボディを返却
// start以降、end以前の日付を昇順として取得できます
func NewCalenderReqBody(start, end string) (io.Reader, error) {
	reqBody := &ReqBody{
		Filter: Filter{
			And: []And{
				{
					Property: "Date",
					Date: FilterDate{
						Before: end,
					},
				},
				{
					Property: "Date",
					Date: FilterDate{
						After: start,
					},
				},
			},
		},
		Sorts: []Sorts{
			{
				Property:  "Date",
				Direction: "ascending",
			},
		},
	}
	reqJSON, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	result := bytes.NewBuffer(reqJSON)

	return result, nil
}
