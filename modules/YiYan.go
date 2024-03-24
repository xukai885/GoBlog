package modules

type YiYan struct {
	Id         int    `json:"id" db:"id"`
	Hitokoto   string `json:"hitokoto" db:"hitokoto"`
	FromSource string `json:"from_source" db:"from_source"`
	FromWho    string `json:"from_who" db:"from_who"`
}
