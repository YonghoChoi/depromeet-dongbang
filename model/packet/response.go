package packet

import (
	"encoding/json"
)

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (o *Response) MarshalJson() (string, error) {
	jsonBytes, err := json.Marshal(o)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

func (o *Response) UnmarshalJson(data string) error {
	return json.Unmarshal([]byte(data), o)
}
