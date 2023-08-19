package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindrome(t *testing.T) {
	assert.True(t, isPalindrome("1100110000110011"))
	assert.False(t, isPalindrome("1110110000110011"))
}
