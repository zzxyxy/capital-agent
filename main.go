package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

var log *zap.Logger

type LoginJson struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

func main() {
	log, _ := zap.NewDevelopment()

	logindata := LoginJson{
		Identifier: "zxyxy@protonmail.com",
		Password:   "Db65y1AvDVyqXTDzv_2ECmtL-e38Gs",
	}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(logindata)

	req, _ := http.NewRequest(
		http.MethodPost,
		"https://demo-api-capital.backend-capital.com/api/v1/session",
		&buf,
	)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-CAP-API-KEY", "q0qubQHyJvJ9Wgax")

	client := &http.Client{}
	resp, _ := client.Do(req)

	log.Info(resp.Status)
	log.Info("This is our first log message using zap!")
}
