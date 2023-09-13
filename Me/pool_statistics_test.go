package Me

import (
	"github.com/stretchr/testify/assert"
	"go_prometheus/config"
	"testing"
)

func TestNewMe4PoolStatistics(t *testing.T) {

	p := NewMe4PoolStatistics(me.baseUrl + "/pool-statistics")

	assert.Equal(t, "pool-statistics", p[0].ObjectName)
	assert.Equal(t, "/meta/pool-statistics", p[0].Meta)
}
