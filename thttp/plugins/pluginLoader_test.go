package plugins

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadPlugins(t *testing.T) {
	plugins := []string{"test-requests/simple.js"}
	assert.Equal(t, 1, len(LoadPlugins(plugins)), "Not loading")
}

func TestPlugin_Run(t *testing.T) {
	pluginPaths := []string{"../../test-requests/simple.js"}
	plugins := LoadPlugins(pluginPaths)
	for i := 0; i < len(plugins); i++ {
		plugin := plugins[i]
		plugin.Run()
	}
}
