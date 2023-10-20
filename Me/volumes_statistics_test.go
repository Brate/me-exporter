package Me

import (
	"github.com/stretchr/testify/assert"
	"go_prometheus/config"
	"testing"
)

func TestNewMe4ShowVolumeStatics(t *testing.T) {

	p := NewMe4ShowVolumeStatics(me.baseUrl + "/show/volume-statistics")

	assert.Equal(t, "volume-statistics", p[0].ObjectName)
	assert.Equal(t, "/meta/volume-statistics", p[0].Meta)

}
