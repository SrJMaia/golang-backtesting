package round

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRound(t *testing.T) {
	r := round()

	assert.Equal(t, "3.00", r)
}
