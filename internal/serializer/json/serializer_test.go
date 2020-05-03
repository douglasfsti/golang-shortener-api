package json

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSerializer(t *testing.T) {
	type person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	p := &person{
		Name: "doug",
		Age:  27,
	}

	serializer := &Serializer{}

	t.Run("should encode struct without error", func(t *testing.T) {
		data, err := serializer.Encode(p)

		assert.NoError(t, err)
		assert.True(t, len(data) > 0)
	})

	t.Run("should decode []byte without error", func(t *testing.T) {
		data, err := serializer.Encode(p)

		assert.NoError(t, err)
		assert.True(t, len(data) > 0)

		var output person
		err = serializer.Decode(data, &output)
		assert.NoError(t, err)
		assert.Equal(t, p.Name, output.Name)
		assert.Equal(t, p.Age, output.Age)
	})
}
