package main

import (
	"fmt"
	"zap-example/utils"

	"go.uber.org/zap"
)

func main() {
	utils.InitializeLogger()
	utils.Logger.Info("Hello World")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			utils.Logger.Debug(fmt.Sprintf("checking for value %d", i), zap.String("value", " even"))
		} else {
			utils.Logger.Debug(fmt.Sprintf("checking for value %d", i), zap.String("value", "odd"))
		}
	}
}
