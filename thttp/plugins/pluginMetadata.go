package plugins

type GetPluginMetadata func() interface{}
type ExecFunction func() interface{}

type PluginMetadata struct {
	Name string
	Type string
	Exec ExecFunction
}
