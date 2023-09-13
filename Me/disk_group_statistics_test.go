package Me

import (
	"github.com/stretchr/testify/assert"
	"go_prometheus/config"
	"testing"
)

func TestNewMe4DiskGroupStatistics(t *testing.T) {

	p := NewMe4DiskGroupStatistics(me.baseUrl + "/disk-group-statistics")

	assert.Equal(t, "disk-group-statistics", p[0].ObjectName)
	assert.Equal(t, "/meta/disk-group-statistics", p[0].Meta)

}
