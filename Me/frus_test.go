package Me

import (
	"github.com/stretchr/testify/assert"
	"go_prometheus/config"
	"testing"
)

func TestNewMe4Frus(t *testing.T) {

	p := NewMe4Frus(me.baseUrl + "/frus")

	assert.Equal(t, "fru", p[0].ObjectName)
	assert.Equal(t, "/meta/enclosure-fru", p[0].Meta)
}
