package main

import (
	"fmt"
	"go-video/base/config"
	"go-video/base/initialize"
	"log"

	"go.uber.org/zap"
)

var (
	port = 5621
)

func main() {
	cfg, err := config.LoadConfig("./base/config")
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	initialize.InitialLogger()

	db, err := initialize.InitDB(cfg.Database)
	if err != nil {
		log.Fatalf("Could not load db: %v", err)
	}

	r := initialize.Routers(db, cfg)

	zap.S().Infof("启动服务器， 端口：%d", port)

	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		zap.S().Error("启动服务器失败 ", err.Error())
	}
}
