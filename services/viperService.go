package services

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetViper(config_string string) (string_map map[string]interface{}, err error) {
	fmt.Println("start to get  viper")
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.SetDefault("application.port", 8080)

	if err = viper.ReadInConfig(); err != nil {
		return string_map, err
	}

	string_map = viper.GetStringMap(config_string)
	return string_map, nil
}
