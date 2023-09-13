package Me

import (
	"github.com/stretchr/testify/assert"
	"go_prometheus/config"
	"testing"
)

func TestNewMe4Ports(t *testing.T) {

	p := NewMe4Ports(me.baseUrl + "/ports")

	assert.Equal(t, "ports", p[0].ObjectName)
	assert.Equal(t, "/meta/port", p[0].Meta)
}
