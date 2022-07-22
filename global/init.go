package global

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

const defaultConfigFile = "config.yaml"

func init() {
	initConfig()
	initRedis()
}

func initConfig() {
	v := viper.New()
	v.SetConfigFile(defaultConfigFile)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&GVA_CONFIG); err != nil {
		fmt.Println(err)
	}
}

func initRedis() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     GVA_CONFIG.Redis.Addr,
		Password: GVA_CONFIG.Redis.Password, // no password set
		DB:       GVA_CONFIG.Redis.DB,       // use default DB
	})

	GVA_REDIS = rdb

}
