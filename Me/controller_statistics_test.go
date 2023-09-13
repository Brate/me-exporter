package Me

import (
	"github.com/stretchr/testify/assert"
	"go_prometheus/config"
	"testing"
)

func TestNewMe4ShowControllerStatistics(t *testing.T) {
	p := NewMe4ControllerStatistics(me.baseUrl + "/show/controller-statistics")

	assert.Equal(t, "controller-statistics", p[0].ObjectName)
	assert.Equal(t, "/meta/controller-statistics", p[0].Meta)
}
