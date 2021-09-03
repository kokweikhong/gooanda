package gooanda

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type connection struct {
	endpoint string
	method   string
	token    string
	data     []byte
}

func (co *connection) connect() []byte {
	client := &http.Client{Timeout: 2 * time.Second}
	var buffer bytes.Buffer
	buffer.WriteString("Bearer ")
	buffer.WriteString(co.token)
	auth := buffer.String()
	req, err := http.NewRequest(co.method, co.endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "v20-golang/0.1")
	req.Header.Set("Authorization", auth)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}
