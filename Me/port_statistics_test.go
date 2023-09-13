package Me

import (
	"github.com/stretchr/testify/assert"
	"go_prometheus/config"
	"testing"
)

func TestNewMe4PortStatistics(t *testing.T) {

	p := NewMe4PortStatistics(me.baseUrl + "/port-statistics")

	assert.Equal(t, "host-port-statistics", p[0].ObjectName)
	assert.Equal(t, "/meta/host-port-statistics", p[0].Meta)
}
