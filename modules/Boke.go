package modules

import "time"

type Boke struct {
	Id           int64     `json:"id" db:"id"`
	CreateTime   time.Time `json:"createtime" db:"createtime"`
	TitleName    string    `json:"titlename" db:"titlename"`
	FileName     string    `json:"filename" db:"filename"`
	Md5FileName  string    `json:"md5filename" db:"md5filename"`
	FilePath     string    `json:"filepath" db:"filepath"`
	Introduction string    `json:"introduction" db:"introduction"`
	Type         int       `json:"type" db:"type"`
}

type BokeType struct {
	Id       int64  `json:"id" db:"id"`
	Typename string `json:"typename" db:"typename"`
	Num      int64  `json:"count_sum" db:"count_sum"`
}

type BokeInfo struct {
	TitleName  string    `json:"titlename"`
	Content    string    `json:"content"`
	CreateTime time.Time `json:"createtime"`
}

type TimeArchive struct {
	Data string `json:"data" db:"yearmounth"`
	Sum  string `json:"sum" db:"sum"`
}
