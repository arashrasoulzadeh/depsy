package structs

type ExecArgs struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}
type ExecStruct struct {
	Name        string     `yaml:"name"`
	Type        string     `yaml:"type"`
	Command     string     `yaml:"command"`
	Path        string     `yaml:"path"`
	Args        []ExecArgs `yaml:"args"`
	PassOnError bool       `yaml:"pass_on_error"`
}
type Config []struct {
	Name   string       `yaml:"name"`
	Become bool         `yaml:"become"`
	Exec   []ExecStruct `yaml:"exec"`
	Hook   string       `yaml:"hook"`
}
