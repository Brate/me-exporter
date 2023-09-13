package Me

import (
	"github.com/stretchr/testify/assert"
	"go_prometheus/config"
	"testing"
)

func TestNewMe4Enclosures(t *testing.T) {

	p := NewMe4Enclosures(me.baseUrl + "/enclosures")

	assert.Equal(t, "enclosures", p[0].ObjectName)
	assert.Equal(t, "/meta/enclosures", p[0].Meta)
}
