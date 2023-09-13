package Me

import (
	"github.com/stretchr/testify/assert"
	"go_prometheus/config"
	"testing"
)

func TestNewMe4DiskStatistics(t *testing.T) {

	p := NewMe4DiskStatistics(me.baseUrl + "/disk-statistics")

	assert.Equal(t, "disk-statistics", p[0].ObjectName)
	assert.Equal(t, "/meta/disk-statistics", p[0].Meta)
}
