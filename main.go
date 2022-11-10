package main

import (
	"go.uber.org/zap"
)

var log *zap.Logger

func main() {
	log, _ := zap.NewDevelopment()

	log.Info("This is our first log message using zap!")
}
