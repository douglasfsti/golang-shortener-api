package serializer

import (
	"github.com/douglasfsti/golang-shortener-api/internal/serializer/json"
	"io"
)

type Serializer interface {
	Decode(input []byte, dst interface{}) error
	DecodeIoReader(input io.Reader, dst interface{}) error
	Encode(input interface{}) ([]byte, error)
}

func NewSerializer() Serializer {
	return &json.Serializer{}
}
