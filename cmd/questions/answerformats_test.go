package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatValidation(t *testing.T) {
	assert.True(t, IsValid(BooleanFormat, "true"))
	assert.False(t, IsValid(BooleanFormat, "potato"))

	assert.True(t, IsValid(TextFormat, "32-1-0321j0djwa0dd2901"))

	assert.False(t, IsValid(IntegerFormat, "potato"))
	assert.False(t, IsValid(IntegerFormat, "1.5"))
	assert.True(t, IsValid(IntegerFormat, "10"))
}
