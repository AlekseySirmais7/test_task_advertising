package config

import (
	"github.com/spf13/viper"
	"log"
)

type DataBaseConfig struct {
	Host            string `json:"host"`
	Port            int    `json:"port"`
	Database        string `json:"database"`
	User            string `json:"user"`
	Password        string `json:"password"`
	ConnectionCount int    `json:"connectionCount"`
}

func GetDBConfig() (DataBaseConfig, error) {
	v := viper.New()
	config := DataBaseConfig{}

	v.SetConfigName("config")
	v.AddConfigPath(".")

	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %v", err)
	}

	config.User = v.GetString("db.user")
	config.Password = v.GetString("db.password")
	config.Host = v.GetString("db.host")
	config.Database = v.GetString("db.database")
	config.Port = v.GetInt("db.port")
	config.ConnectionCount = v.GetInt("db.connectionCount")

	err = v.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Decoding problem: , %v", err)
	}

	return config, err
}

func GetTestConfig() (DataBaseConfig, error) {
	v := viper.New()
	config := DataBaseConfig{}

	v.SetConfigName("config")
	v.AddConfigPath(".")
	v.AddConfigPath("../../../.")
	v.AddConfigPath("../../../../build/")

	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %v", err)
	}

	config.User = v.GetString("test.user")
	config.Password = v.GetString("test.password")
	config.Host = v.GetString("test.host")
	config.Database = v.GetString("test.database")
	config.Port = v.GetInt("test.port")
	config.ConnectionCount = v.GetInt("test.connectionCount")

	err = v.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Decoding problem: , %v", err)
	}

	return config, err
}
