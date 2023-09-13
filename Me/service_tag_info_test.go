package Me

import (
	"github.com/stretchr/testify/assert"
	"go_prometheus/config"
	"testing"
)

func TestNewMe4ServiceTagInfo(t *testing.T) {
	p := NewMe4ServiceTagInfo(me.baseUrl + "/service-tag-info")

	assert.Equal(t, "service-tag-info", p[0].ObjectName)
	assert.Equal(t, "/meta/service-tag-info", p[0].Meta)

}
