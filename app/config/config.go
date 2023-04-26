package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/joho/godotenv"
	"github.com/lovelyrrg51/go_backend/app/logger"
	"github.com/mitchellh/mapstructure"
)

type Config struct {
	AppPort string `mapstructure:"PORT"`
	ENV     string `mapstructure:"ENV"`

	DbHost        string `mapstructure:"DB_HOST"`
	DbPort        string `mapstructure:"DB_PORT"`
	DbUser        string `mapstructure:"DB_USER"`
	DbPass        string `mapstructure:"DB_PASS"`
	DbName        string `mapstructure:"DB_NAME"`
	DbMaxIdle     string `mapstructure:"DB_MAX_IDLE"`
	DbMaxOpenCon  string `mapstructure:"DB_MAX_OPEN_CON"`
	DbMaxLifetime string `mapstructure:"DB_MAX_LIFETIME"`

	JWTSecretKey string `mapstructure:"JWT_SECRET_KEY"`
}

var Cfg *Config

func (c *Config) IsProduction() bool {
	return c.ENV == "PRODUCTION"
}

func getBasePath() string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Fprintf(os.Stderr, "Unable to identify current directory (needed to load .env)")
		os.Exit(1)
	}
	return filepath.Dir(file)
}

func getConfig() *Config {

	//path := getBasePath()
	godotenv.Load()

	var conf *Config
	envMap := make(map[string]interface{})
	for _, env := range os.Environ() {
		key := env[:strings.Index(env, "=")]
		value := env[strings.Index(env, "=")+1:]
		envMap[key] = value
	}

	err := mapstructure.Decode(envMap, &conf)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	return conf
}

func init() {
	Cfg = getConfig()
}
