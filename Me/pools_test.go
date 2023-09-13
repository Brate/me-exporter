package Me

import (
	"github.com/stretchr/testify/assert"
	"go_prometheus/config"
	"testing"
)

func TestNewMe4Pools(t *testing.T) {

	p := NewMe4Pools(me.baseUrl + "/pools")

	assert.Equal(t, "pools", p[0].ObjectName)
	assert.Equal(t, "/meta/pools", p[0].Meta)
}
