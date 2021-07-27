package gateway

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendMessage(t *testing.T) {
	client := NewLineNotityClient()
	err := client.SendMessage("これはテストです。")
	assert.Nil(t, err)
}
