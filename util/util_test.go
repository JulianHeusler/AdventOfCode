package util_test

import (
	"adventofcode/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInt(t *testing.T) {
	assert.Equal(t, 1, util.GetInt("1"))
	assert.Equal(t, 123, util.GetInt("123"))
	assert.Equal(t, 0, util.GetInt("a"))
	assert.Equal(t, 0, util.GetInt(""))
}
