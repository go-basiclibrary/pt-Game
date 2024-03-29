package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"log"
)

var Conf = new(Config)

func Init(path string) (*Config, error) {
	//默认使用dev.yaml配置文件
	if path == "" {
		viper.AddConfigPath("./Conf")
		viper.SetConfigName("dev")
		viper.SetConfigType("yaml")
	} else {
		//parse指定yaml
		viper.SetConfigFile(fmt.Sprintf("Conf/%s", path))
	}

	//加载配置
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	//解析
	if err := viper.Unmarshal(Conf); err != nil {
		return nil, errors.Wrapf(err, "faild to load config")
	}

	//开启配置热加载
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Println("config file changed:", in.Name)
	})
	return Conf, nil
}

func GetConfig() *Config {
	return Conf
}
