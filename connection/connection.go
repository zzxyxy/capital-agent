package connection

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

var log *zap.Logger

func GetConnection(identifier, password, url, xcapapikey string) connection {
	return connection{
		identifier, password, url, xcapapikey,
	}
}

type connection struct {
	Identifier string
	Password   string
	Url        string
	Xcapapikey string
}

func (o connection) Open() {
	log, _ = zap.NewDevelopment()

	logindata := LoginJson{
		Identifier: o.Identifier,
		Password:   o.Password,
	}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(logindata)

	req, _ := http.NewRequest(
		http.MethodPost,
		o.Url+"/api/v1/session",
		&buf,
	)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-CAP-API-KEY", o.Xcapapikey)

	client := &http.Client{}
	resp, _ := client.Do(req)

	log.Info(fmt.Sprintf("Connection status is %s", resp.Status))
}
