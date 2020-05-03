package serializer

import (
	"github.com/douglasfsti/golang-shortener-api/internal/serializer/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewJson(t *testing.T) {
	t.Run("should return an instance of json serializer", func(t *testing.T) {
		serializer := NewSerializer()

		assert.IsType(t, &json.Serializer{}, serializer)
	})
}
