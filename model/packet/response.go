package packet

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Response struct {
	ID ID
}

type ResponseError struct {
	Err error
}

func (err ResponseError) MarshalJSON() ([]byte, error) {
	if err.Err == nil {
		return []byte("null"), nil
	}

	return []byte(fmt.Sprintf("\"%v\"", err.Err)), nil
}

func (err *ResponseError) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, v); err != nil {
		return err
	}

	if v == nil {
		err.Err = nil
		return nil
	}

	switch tv := v.(type) {
	case string:
		err.Err = errors.New(tv)
		return nil
	default:
		return errors.New("ResponseError unmarshal failed")
	}
}
