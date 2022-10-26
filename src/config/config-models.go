package config

type Config struct {
	ProjectName string `yaml:"projectName"`
	Collections map[string]Collection
}

type Collection struct {
	Fields Fields
	Routes Routes
}

type Fields map[string]Field

type Field struct {
	Type     string
	Required bool
}

type Routes struct {
	Get       bool
	GetSingle bool
	Post      bool
	Put       bool
	Delete    bool
}
