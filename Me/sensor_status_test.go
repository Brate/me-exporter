package Me

import (
	"github.com/stretchr/testify/assert"
	"go_prometheus/config"
	"testing"
)

func TestNewMe4Sensors(t *testing.T) {

	p := NewMe4Sensors(me.baseUrl + "/sensor-status")

	assert.Equal(t, "sensor", p[0].ObjectName)
	assert.Equal(t, "/meta/sensors", p[0].Meta)
}
