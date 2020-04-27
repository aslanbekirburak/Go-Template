package main

import (
	controllers "common_dashboard_backend/apps/Dashboard/controllers"
	"common_dashboard_backend/common/gateways/redis"
	. "common_dashboard_backend/common/logger"
	common "common_dashboard_backend/common/projectArch/interactors"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Specification struct {
	RedisHostComm string
	RedisHostProd string
	RedisPort     string

	LogFile string
	Debug   bool
}

type Environment struct {
	Env string
}

var s Specification
var e Environment

var logFile *os.File

func main() {
	var err error

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	// viper.SetConfigName("config.yml")

	// Init configuration
	viper.AddConfigPath("../../")
	// viper.AddConfigPath("./")
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {
		// Handle errors reading the config file
		log.Fatal(err)
	}

	s.RedisHostProd = viper.GetString("app.redis.host.prod")
	s.RedisPort = viper.GetString("app.redis.port")

	s.LogFile = viper.GetString("app.redis.log.file")
	s.Debug = viper.GetBool("app.redis.log.debug")

	viper.WatchConfig()
	log.Println("APP: Common Dashboard Backend")
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name, "APP: Common Dashboard Backend")
	})

	// Init logging
	logFile, err = os.OpenFile(s.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	InitLogger(logFile, logFile, logFile, s.Debug)

	//redis initialization

	redis.Init(s.RedisHostComm, s.RedisHostProd, s.RedisPort)
	common.RedisCommStorage = redis.GetRedisCommStorage()
	defer redis.Close()

	// controllers.StartApplicationBackend(s.MediaPath)
	controllers.StartApplicationBackend()
}
