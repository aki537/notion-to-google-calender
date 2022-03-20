package api

/*
	Notionから取得できるリスト・ブロックの構造体
*/
type ResList struct {
	Object     string    `json:"object"`
	Results    []*Result `json:"results"`
	NextCursor string    `json:"next_cursor"`
	HasMore    bool      `json:"has_more"`
	Type       string    `json:"type"`
}
type Result struct {
	Object         string       `json:"object"`
	ID             string       `json:"id"`
	CreatedTime    string       `json:"created_time"`
	LastEditedTime string       `json:"last_edited_time"`
	CreatedBy      CreatedBy    `json:"created_by"`
	LastEditedBy   LastEditedBy `json:"last_edited_by"`
	Cover          interface{}  `json:"cover"`
	Icon           interface{}  `json:"icon"`
	Parent         Parent       `json:"parent"`
	Archived       bool         `json:"archived"`
	Properties     Properties   `json:"properties"`
	URL            string       `json:"url"`
	HasChildren    bool         `json:"has_children"`
	Type           string       `json:"type"`
	Paragraph      Paragraph    `json:"paragraph"`
}
type Icon struct {
	Type  string `json:"type"`
	Emoji string `json:"emoji"`
}
type CreatedBy struct {
	Object string `json:"object"`
	ID     string `json:"id"`
}
type LastEditedBy struct {
	Object string `json:"object"`
	ID     string `json:"id"`
}
type Text struct {
	Content string      `json:"content"`
	Link    interface{} `json:"link"`
}
type Annotations struct {
	Bold          bool   `json:"bold"`
	Italic        bool   `json:"italic"`
	Strikethrough bool   `json:"strikethrough"`
	Underline     bool   `json:"underline"`
	Code          bool   `json:"code"`
	Color         string `json:"color"`
}
type Title struct {
	Type        string      `json:"type"`
	Text        Text        `json:"text"`
	Annotations Annotations `json:"annotations"`
	PlainText   string      `json:"plain_text"`
	Href        interface{} `json:"href"`
}

type Tags struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Type        string        `json:"type"`
	MultiSelect []MultiSelect `json:"multi_select"`
}

type MultiSelect struct {
	Options []interface{} `json:"options"`
}
type RichText struct {
	Type        string      `json:"type"`
	Text        Text        `json:"text"`
	Annotations Annotations `json:"annotations"`
	PlainText   string      `json:"plain_text"`
	Href        interface{} `json:"href"`
}
type Property struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Type     string     `json:"type"`
	RichText []RichText `json:"rich_text"`
}
type Date struct {
	ID   string       `json:"id"`
	Name string       `json:"name"`
	Type string       `json:"type"`
	Date ChildrenDate `json:"date"`
}
type ChildrenDate struct {
	Start    string `json:"start"`
	End      string `json:"end"`
	TimeZone string `json:"time_zone"`
}
type Name struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Title []Title `json:"title"`
}
type Properties struct {
	Tags     Tags     `json:"Tags"`
	Property Property `json:"Property"`
	Date     Date     `json:"Date"`
	Name     Name     `json:"Name"`
}
type Parent struct {
	Type   string `json:"type"`
	PageID string `json:"page_id"`
}
type Paragraph struct {
	RichText []RichText `json:"rich_text"`
	Color    string     `json:"color"`
}

/*
	NotionAPIにリクエストする際のボディ
*/
type ReqBody struct {
	Filter Filter  `json:"filter,omitempty"`
	Sorts  []Sorts `json:"sorts,omitempty"`
}
type Or struct {
	Property string     `json:"property,omitempty"`
	Date     FilterDate `json:"date,omitempty"`
}
type Filter struct {
	Or  []Or  `json:"or,omitempty"`
	And []And `json:"and,omitempty"`
}
type Sorts struct {
	Property  string `json:"property,omitempty"`
	Direction string `json:"direction,omitempty"`
}
type And struct {
	Property string     `json:"property,omitempty"`
	Date     FilterDate `json:"date,omitempty"`
}
type FilterDate struct {
	After  string `json:"after,omitempty"`
	Before string `json:"before,omitempty"`
}
