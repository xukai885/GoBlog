package modules

import "time"

type Share struct {
	Sum   int         `json:"sum"`
	Share []ShareInfo `json:"shareinfo"`
}

type ShareInfo struct {
	Date  time.Time `json:"date" db:"date"`
	Url   string    `json:"url" db:"url"`
	Title string    `json:"title" db:"title"`
	Type  string    `json:"type" db:"type"`
}
