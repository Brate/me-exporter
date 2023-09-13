package Me

import (
	"github.com/stretchr/testify/assert"
	"go_prometheus/config"
	"testing"
)

func TestNewMe4TierStatistics(t *testing.T) {

	p := NewMe4TierStatistics(me.baseUrl + "/tier-statistics")

	assert.Equal(t, "tier-statistics", p[0].ObjectName)
	assert.Equal(t, "/meta/tier-statistics", p[0].Meta)
}
