package foo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFool(t *testing.T) {
	assert.Equal(t, "Foo", Fool(), "they should be equal")
}