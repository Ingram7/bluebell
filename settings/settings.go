package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

func Init() (err error) {
	// viper.SetConfigFile("./config.toml")  // 或者
	viper.SetConfigName("config") // 并不需要后缀
	viper.AddConfigPath(".")
	viper.AddConfigPath("./settings")
	if err = viper.ReadInConfig(); err != nil {
		return errors.WithMessage(err, "Init with read config failed")
	}

	if err = viper.Unmarshal(Conf); err != nil {
		return errors.WithMessage(err, "viper.Unmarshal failed")
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config file changed:", in.Name)
		if err = viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		}
	})

	return
}

type AppConfig struct {
	StartTime    string `mapstructure:"start_time"`
	MachineID    int64  `mapstructure:"machine_id"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
	*LogConfig   `mapstructure:"log"`
}

type MysqlConfig struct {
	DSN string `mapstructure:"dsn"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       string `mapstructure:"db"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}
