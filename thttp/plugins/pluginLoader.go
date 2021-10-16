package plugins

import (
	"github.com/dop251/goja"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

type Plugin struct {
	ScriptPath string
}

func LoadPlugins(plugins []string) []Plugin {
	loadedPlugins := make([]Plugin, 0)
	for i := 0; i < len(plugins); i++ {
		loadedPlugins = append(loadedPlugins, Plugin{ScriptPath: plugins[i]})
	}
	return loadedPlugins
}

func (plugin *Plugin) Run() {
	vm := goja.New()
	file, err := os.Open(plugin.ScriptPath)
	if err != nil {
		log.Errorln(err)
		return
	}
	defer file.Close()
	scriptBytes, err := io.ReadAll(file)
	if err != nil {
		log.Errorln(err)
		return
	}
	scriptString := string(scriptBytes)
	_, err = vm.RunString(scriptString)
	if err != nil {
		log.Errorln(err)
		return
	}
	var fn func() map[string]string
	err = vm.ExportTo(vm.Get("thttp"), &fn)
	if err != nil {
		log.Errorln(err)
		return
	}
	pluginMetadata := fn()
	log.Println(runExec(vm, pluginMetadata["exec"]))

}

func runExec(vm *goja.Runtime, execFunction string) interface{} {
	var fn func() interface{}
	err := vm.ExportTo(vm.Get(execFunction), &fn)
	if err != nil {
		log.Errorln(err)
		return nil
	}
	r := fn()
	return r
}
