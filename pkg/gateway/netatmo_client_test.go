package gateway

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCarbonDioxideConcentration(t *testing.T) {
	client := NewNetatmoClient()
	_, err := client.GetCarbonDioxideConcentration()
	assert.Nil(t, err)
}
