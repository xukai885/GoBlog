package modules

type AppUser struct {
	UserID   uint64 `json:"user_id" db:"user_id"`
	Username string `json:"username" db:"username" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}

type AddAppUser struct {
	UserID     uint64 `json:"user_id" db:"user_id"`
	Username   string `json:"username" db:"username" binding:"required"`
	Password   string `json:"password" db:"password" binding:"required"`
	RePassword string `json:"rePassword" binding:"required,eqfield=Password"`
}
