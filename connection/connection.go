package connection

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"io/ioutil"

	"go.uber.org/zap"
)

var log *zap.Logger

func GetConnection(identifier, password, url, xcapapikey string) (connection, error) {
	connection := connection{
		identifier, password, url, xcapapikey,
	}
	err := connection.Open()
	return connection, err
}

type connection struct {
	Identifier string
	Password   string
	Url        string
	Xcapapikey string
}

func (o connection) Open() error {
	log, err := zap.NewDevelopment()

	if err != nil {
		return errors.New("error while open connection")
	}

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
	resp, err := client.Do(req)

	if err != nil {
		return errors.New("error while reading session reply body")
	}

	var session_info Session_reply
	bodybytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return errors.New("cant unmarshall json")
	}

	err = json.Unmarshal(bodybytes, &session_info)

	if err != nil {
		return errors.New("cant unmarshall json")
	}

	log.Info(fmt.Sprintf("Connection status is %s", resp.Status))
	return nil
}
