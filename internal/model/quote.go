package model

type Quote struct {
    ID     int    `json:"id"`
    Author string `json:"author"`
    Text   string `json:"text"`
}
