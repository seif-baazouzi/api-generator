package config

type Config struct {
	ProjectName string
	Collections map[string]Collection
}

type Collection map[string]Field

type Field struct {
	Type     string
	Required bool
}
