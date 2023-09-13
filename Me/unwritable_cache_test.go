package Me

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMe4UnwritableCache(t *testing.T) {

	p := NewMe4UnwritableCache(me.baseUrl + "/unwritable-cache")

	assert.Equal(t, "unwritable-system-cache", p[0].ObjectName)
	assert.Equal(t, "/meta/unwritable-cache", p[0].Meta)
}
