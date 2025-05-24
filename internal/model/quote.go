package model

type Quote struct {
	ID     int    `json:"-"`
	UUID   string `json:"id"`
	Author string `json:"author"`
	Text   string `json:"quote"`
}
