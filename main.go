package main

import (
	connection "capital-agent/connection"

	"go.uber.org/zap"
)

const IDENTIFIER = "zxyxy@protonmail.com"
const PASSWORD = "Db65y1AvDVyqXTDzv_2ECmtL-e38Gs"
const XCAPAPIKEY = "q0qubQHyJvJ9Wgax"
const URL = "https://demo-api-capital.backend-capital.com"

var log *zap.Logger

func main() {
	log, _ = zap.NewDevelopment()

	conn := connection.GetConnection(
		IDENTIFIER,
		PASSWORD,
		URL,
		XCAPAPIKEY,
	)
	conn.Open()
}
