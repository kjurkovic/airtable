package util

import "encoding/json"

func ToJson(obj any) (string, error) {
	buf, err := json.Marshal(obj)

	if err != nil {
		return "", err
	}

	return string(buf), nil
}
