package config

import (
	"os"

	root "skynet/pkg"
)

func GetConfig() *root.Config {
	return &root.Config{
		Mongo: &root.MongoConfig{
			Ip:     envOrDefaultString("skynet:mongo:ip", "127.0.0.1:27017"),
			DbName: envOrDefaultString("skynet:mongo:dbName", "skynet")},
		Server: &root.ServerConfig{Port: envOrDefaultString("skynet:server:port", ":8080")}}
}

func envOrDefaultString(envVar string, defaultValue string) string {
	value := os.Getenv(envVar)
	if value == "" {
		return defaultValue
	}

	return value
}
