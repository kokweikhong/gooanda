package gooanda

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func urlAddQuery(endpoint string, querys interface{}) (string, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return "", err
	}
	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return "", err
	}
	qrys, err := convertQuerys(querys)
	if err != nil {
		return "", err
	}
	for k, v := range qrys {
		q.Add(k, fmt.Sprintf("%v", v))
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}

func convertQuerys(querys interface{}) (map[string]interface{}, error) {
	mapQuery := make(map[string]interface{})
	data, err := json.Marshal(querys)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &mapQuery)
	if err != nil {
		return nil, err
	}
	return mapQuery, nil
}
