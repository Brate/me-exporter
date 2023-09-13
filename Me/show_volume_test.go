package Me

import (
	"github.com/stretchr/testify/assert"
	"go_prometheus/config"
	"testing"
)

func TestNewMe4ShowVolume(t *testing.T) {

	p := NewMe4ShowVolume(me.baseUrl + "/show/volumes")

	assert.Equal(t, "volume", p[0].ObjectName)
	assert.Equal(t, "/meta/volumes", p[0].Meta)
}
