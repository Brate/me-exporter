package Me

import (
	"github.com/stretchr/testify/assert"
	"go_prometheus/config"
	"testing"
)

func TestNewMe4Fans(t *testing.T) {

	p := NewMe4Fans(me.baseUrl + "/fans")

	assert.Equal(t, "fan-details", p[0].ObjectName)
	assert.Equal(t, "/meta/fan", p[0].Meta)
}
