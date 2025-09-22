package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	assert.Equal(t, 5, result, "2 + 3 should equal 5")
}

func TestMultiply(t *testing.T) {
	result := Multiply(3, 4)
	assert.Equal(t, 12, result, "3 * 4 should equal 12")
}
