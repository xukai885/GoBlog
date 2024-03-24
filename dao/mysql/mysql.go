package mysql

import (
	"GoBlog/settings"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init(cfg *settings.MysqlConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Dbname,
	)

	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("mysql connect err %v \n", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(cfg.Maxconns)
	db.SetMaxIdleConns(cfg.Maxidleconns)
	return
}
func Close() {
	_ = db.Close()
}
