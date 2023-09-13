package Me

import (
	"github.com/stretchr/testify/assert"
	"go_prometheus/config"
	"testing"
)

func TestNewMe4TiersInfo(t *testing.T) {

	p := NewMe4TiersInfo(me.baseUrl + "/tiers")

	assert.Equal(t, "tiers", p[0].ObjectName)
	assert.Equal(t, "/meta/tiers", p[0].Meta)
}
