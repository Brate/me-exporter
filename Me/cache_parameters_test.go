package Me

import (
	"github.com/stretchr/testify/assert"
	"go_prometheus/config"
	"testing"
)

var(
	client
)

func setupTest(t *testing.TB) func(tb testing.TB) {
	config.LoadConfig()

	return func(tb testing.TB) { }
}

func TestNewMe4CacheParameters(t *testing.T) {
	fimTeste := setupTest(t)
	defer fimTeste(t)

	p := NewMe4CacheParameters(me.baseUrl + "/cache-parameters")

	assert.Equal(t, "system-cache-parameters", p[0].ObjectName)
	assert.Equal(t, "/meta/cache-settings", p[0].Meta)
}
