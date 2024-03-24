package modules

import "time"

type Images struct {
	Id         int64     `json:"id" db:"id"`
	CreateTime time.Time `json:"createtime" db:"createtime"`
	Name       string    `json:"name" db:"name"`
	Path       string    `json:"path" db:"path"`
	ImageUrl   string    `json:"imagesUrl" db:"imagesUrl"`
	ImageSize  int64     `json:"size" db:"size"`
}
