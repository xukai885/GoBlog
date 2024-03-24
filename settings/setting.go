package settings

import (
	"errors"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 全局变量，保存程序的所有配置信息
var Conf = new(Appconfig)

type Appconfig struct {
	AppName               string `mapstructure:"app_name"`
	Port                  int    `mapstructure:"port"`
	StartTime             string `mapstructure:"start_time"`
	MachineID             uint16 `mapstructure:"machine_id"`
	RegistrationsSwitch   bool   `mapstructure:"registration_switch"`
	YiYanSyncMysqlToRedis bool   `mapstructure:"yiyan_sync_mysql_to_redis"`
	YiYanReadToRedis      bool   `mapstructure:"yiyan_read_to_redis"`
	*FileConfig           `mapstructure:"file"`
	*LogConfig            `mapstructure:"log"`
	*MysqlConfig          `mapstructure:"mysql"`
	*RedisConfig          `mapstructure:"redis"`
}
type FileConfig struct {
	MdfilePath string `mapstructure:"mdfile_path"`
	ImagePath  string `mapstructure:"image_path"`
	CtUrl      string `mapstructure:"ct_url"`
}

type LogConfig struct {
	Mode      string `mapstructure:"mode"`
	Lever     string `mapstructure:"lever"`
	Filename  string `mapstructure:"filename"`
	Maxsize   int    `mapstructure:"max_size"`
	Maxage    int    `mapstructure:"max_age"`
	Maxbackup int    `mapstructure:"max_backup"`
}

type MysqlConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Dbname       string `mapstructure:"dbname"`
	Maxconns     int    `mapstructure:"max_conns"`
	Maxidleconns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	PoolSize int    `mapstructure:"poolsize"`
	Db       int    `mapstructure:"db"`
}

func Init() (err error) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		return errors.New("缺少环境变量 APP_ENV 请配置")
	}

	configfile := fmt.Sprintf("config-%s.yaml", env)

	viper.SetConfigFile(configfile)
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("err viper.readinconfig %v\n", err)
		return
	}

	//	 把读取到的配置信息反序列化到Conf变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper unmarshal err %v \n", err)
	}

	//配置热加载
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("配置文件修改了...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("vipre unmarshal err:%v \n", err)
		}
	})
	return
}
