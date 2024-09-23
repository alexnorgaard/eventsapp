package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Create private data struct to hold config options.
type config struct {
	Postgres struct {
		Host     string
		Port     string
		Database string
		User     string
		Password string
	}
	Google_geocoding_api struct {
		Api_key string
	}
	S3 struct {
		Access_key          string
		Secret_key          string
		Bucket_name_banners string
	}
}

// Read the config file from the current directory and marshal
// into the conf config struct.
func GetConf() *config {
	viper.AddConfigPath(".")
	viper.AddConfigPath("../../")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("%v", err)
	}
	conf := &config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}

	return conf
}
