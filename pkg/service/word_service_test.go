package service

import (
	"testing"

	"github.com/jtreeves/wordiest-api/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestWordService(t *testing.T) {
	ws := NewWordService()

	t.Run("Do", func(t *testing.T) {
		expected := model.Response{
			Data: model.Word{
				ID:     1,
				Value:  "word",
				Length: 4,
				Letter: "W",
			},
		}

		actual := ws.Do()

		assert.Equal(t, expected, actual)
	})
}
