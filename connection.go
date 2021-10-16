package gooanda

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type connection struct {
	endpoint string
	method   string
	token    string
	data     []byte
}

func (co *connection) connect() ([]byte, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	var buffer bytes.Buffer
	buffer.WriteString("Bearer ")
	buffer.WriteString(co.token)
	auth := buffer.String()
	req, err := http.NewRequest(co.method, co.endpoint, bytes.NewBuffer(co.data))
	if err != nil {
		return nil, fmt.Errorf("failed to request api from %v, %v", co.endpoint, err)
	}
	// req.Header.Set("User-Agent", "v20-golang/0.1")
	req.Header.Set("Authorization", auth)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to request api after set token, %v", err)
	}
	defer resp.Body.Close()
	var body []byte
	if strings.Contains(co.endpoint, "stream") {
		body, err = streamApiConnect(resp)
		if err != nil {
			return nil, err
		}
	} else {
		body, err = restApiConnect(resp)
		if err != nil {
			return nil, err
		}
	}
	return body, nil
}

func restApiConnect(resp *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func streamApiConnect(resp *http.Response) ([]byte, error) {
	reader := bufio.NewReader(resp.Body)
	body, err := reader.ReadBytes('\n')
	if err != nil {
		return nil, err
	}
	return body, nil
}
