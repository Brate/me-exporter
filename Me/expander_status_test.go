package Me

import (
	"github.com/stretchr/testify/assert"
	"go_prometheus/config"
	"testing"
)

func TestNewMe4ExpanderStatus(t *testing.T) {

	p := NewMe4ExpanderStatus(me.baseUrl + "/expander-status")

	assert.Equal(t, "enclosure-id", p[0].ObjectName)
	assert.Equal(t, "/meta/sas-status-controller-a", p[0].Meta)
}
