package Me

import (
	"github.com/stretchr/testify/assert"
	"go_prometheus/config"
	"testing"
)

func TestNewMe4Disks(t *testing.T) {

	p := NewMe4Disks(me.baseUrl + "/disks")

	assert.Equal(t, "drive", p[0].ObjectName)
	assert.Equal(t, "/meta/drives", p[0].Meta)
}
