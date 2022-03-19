package main

import (
	"fmt"

	"github.com/abdolrhman/simple-go-lang-rest-api/pkg/config"
	"github.com/abdolrhman/simple-go-lang-rest-api/pkg/database"
	"github.com/abdolrhman/simple-go-lang-rest-api/pkg/logger"
	"github.com/abdolrhman/simple-go-lang-rest-api/pkg/router"
	"github.com/spf13/viper"
)

func main() {
	if err := config.Setup(); err != nil {
		logger.Fatalf("config.Setup() error: %s", err)
	}

	if err := database.Setup(); err != nil {
		logger.Fatalf("database.Setup() error: %s", err)
	}

	db := database.GetDB()
	r := router.Setup(db)

	host := "127.0.0.1"
	if h := viper.GetString("server.host"); h != "" {
		host = h
	}
	logger.Infof("Server is starting at %s:%s", host, viper.GetString("server.port"))
	fmt.Printf("Server is starting at %s:%s Check logs for details.", host, viper.GetString("server.port"))
	logger.Fatalf("%v", r.Run(host+":"+viper.GetString("server.port")))
}
