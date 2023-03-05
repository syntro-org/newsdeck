package config

import "os"

type Config struct {
	Port    string
	ApiKey  string
	AppMode string
}

const DEF_PORT = "8080"
const CHANGE_ME = "change me"
const DEBUG_MODE = "debug"
const RELEASE_MODE = "release"

func New() Config {
	return Config{
		Port:    getValue("PORT", DEF_PORT),
		ApiKey:  getValue("API_KEY", CHANGE_ME),
		AppMode: getValue("APP_MODE", DEBUG_MODE),
	}
}

func getValue(key, def string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return def
}
