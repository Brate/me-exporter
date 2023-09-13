package Me

import (
	"github.com/stretchr/testify/assert"
	"go_prometheus/config"
	"testing"
)

func TestNewMe4DiskGroups(t *testing.T) {

	p := NewMe4DiskGroups(me.baseUrl + "/disk-groups")

	assert.Equal(t, "disk-group", p[0].ObjectName)
	assert.Equal(t, "/meta/disk-groups", p[0].Meta)
}
