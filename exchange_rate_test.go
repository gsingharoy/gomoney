package gomoney

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEurExchangeRate(t *testing.T) {
	t.Log("When the endpoint is working")
	_, err := newEurExchangeRate()
	assert.Nil(t, err)
}
