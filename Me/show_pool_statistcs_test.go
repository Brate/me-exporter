package Me

import (
	"github.com/stretchr/testify/assert"
	"go_prometheus/config"
	"testing"
)

func TestNewMe4ShowPoolStatistics(t *testing.T) {

	p := NewMe4ShowPoolStatistics(me.baseUrl + "/show/pool-statistics")

	assert.Equal(t, "/meta/pool-statistics", p[0].Meta)
	assert.Equal(t, "pool-statistics", p[0].ObjectName)
}
