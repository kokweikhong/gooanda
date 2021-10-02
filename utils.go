package gooanda

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func urlAddQuery(endpoint string, querys interface{}) (string, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return "", fmt.Errorf("failed to parse url from %v, %v", endpoint, err)
	}
	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return "", fmt.Errorf("failed to parse query from url %v, %v", endpoint, err)
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
		return nil, fmt.Errorf("failed to marshal data at %T, %v", querys, err)
	}
	err = json.Unmarshal(data, &mapQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal querys %s to %T, %v", string(data), mapQuery, err)
	}
	return mapQuery, nil
}
