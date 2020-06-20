package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVectorString(t *testing.T) {
	vector := Vector{-1.2, 3, 4.56}
	assert.Equal(t, "(-1.20, 3.00, 4.56)", fmt.Sprintf("%v", vector))
}
