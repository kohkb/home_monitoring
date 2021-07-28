package main

import (
	"log"
	"strconv"

	"github.com/kohkb/home_monitoring/pkg/gateway"
)

func main() {
	var Logger *log.Logger
	netatmo_client := gateway.NewNetatmoClient()
	co2_value, err := netatmo_client.GetCarbonDioxideConcentration()

	if err != nil {
		Logger.Fatal(err)
	}
	// TODO ログを設定する
	if co2_value > 1000 {
		line_notify_client := gateway.NewLineNotityClient()
		msg := "現在のCO2濃度は" + strconv.Itoa(co2_value) + "ppmです。換気してください。"
		line_notify_client.SendMessage(msg)
	}
}
