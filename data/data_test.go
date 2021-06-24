package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	r := sum(1, 2)

	assert.Equal(t, 3, r)
}
