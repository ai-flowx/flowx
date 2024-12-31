package config

var (
	Build   string
	Commit  string
	Version string
)

type Config struct {
	Cache  Cache  `yaml:"cache"`
	Flow   Flow   `yaml:"flow"`
	Gpt    Gpt    `yaml:"gpt"`
	Memory Memory `yaml:"memory"`
	Store  Store  `yaml:"store"`
	Tool   []Tool `yaml:"tool"`
}

type Cache struct {
	Provider string `yaml:"provider"`
	Url      string `yaml:"url"`
}

type Flow struct {
	Channel string `yaml:"channel"`
}

type Gpt struct {
	Provider string `yaml:"provider"`
	Api      string `yaml:"api"`
	Token    string `yaml:"token"`
}

type Memory struct {
	Type string `yaml:"type"`
}

type Store struct {
	Provider string `yaml:"provider"`
	Url      string `yaml:"url"`
	Path     string `yaml:"path"`
}

type Tool struct {
	Name string
}
