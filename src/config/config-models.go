package config

type Config struct {
	Collections map[string]Collection
}

type Collection map[string]Field

type Field struct {
	Type     string
	Required bool
}
