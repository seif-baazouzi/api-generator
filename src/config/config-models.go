package config

type Config struct {
	ProjectName string `yaml:"projectName"`
	Collections map[string]Collection
}

type Collection map[string]Field

type Field struct {
	Type     string
	Required bool
}
