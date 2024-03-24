package snowflake

import (
	"GoBlog/settings"
	"fmt"
	"time"

	"github.com/sony/sonyflake"
)

var (
	sonyFlake     *sonyflake.Sonyflake
	sonyMachineID uint16
)

func getMachineID() (uint16, error) {
	return sonyMachineID, nil
}

// 传入当前机器id
func Init(cfg *settings.Appconfig) (err error) {
	sonyMachineID = cfg.MachineID
	t, _ := time.Parse("2006-01-02", cfg.StartTime)
	settings := sonyflake.Settings{
		StartTime: t,
		MachineID: getMachineID,
	}
	sonyFlake = sonyflake.NewSonyflake(settings)
	return
}

// GetID 返回生成的id值
func GetID() (id uint64, err error) {
	if sonyFlake == nil {
		err = fmt.Errorf("sony flake noe inited")
		return
	}
	id, err = sonyFlake.NextID()
	return
}
