package json

import (
	"encoding/json"
	"fmt"
	"io"
)

const (
	FailedDecode = "%s serializer.json.Decode"
	FailedEncode = "%s serializer.json.Encode"
)

type Serializer struct {
}

func (s *Serializer) Decode(input []byte, dst interface{}) error {
	if err := json.Unmarshal(input, &dst); err != nil {
		err = fmt.Errorf(FailedDecode, err.Error())
		return err
	}

	return nil
}

func (s *Serializer) DecodeIoReader(input io.Reader, dst interface{}) error {
	if err := json.NewDecoder(input).Decode(&dst); err != nil {
		err = fmt.Errorf(FailedDecode, err.Error())
		return err
	}

	return nil
}

func (s *Serializer) Encode(input interface{}) ([]byte, error) {
	data, err := json.Marshal(input)
	if err != nil {
		err = fmt.Errorf(FailedEncode, err.Error())
		return []byte{}, err
	}

	return data, nil
}
