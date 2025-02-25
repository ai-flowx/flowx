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
	Prompt Prompt `yaml:"prompt"`
	Store  Store  `yaml:"store"`
	Tool   []Tool `yaml:"tool"`
}

type Cache struct {
	Provider string `yaml:"provider"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Pass     string `yaml:"pass"`
}

type Flow struct {
	Channel string `yaml:"channel"`
}

type Gpt struct {
	Provider string `yaml:"provider"`
	Api      string `yaml:"api"`
	Key      string `yaml:"key"`
	Endpoint string `yaml:"endpoint"`
}

type Memory struct {
	Type string `yaml:"type"`
}

type Prompt struct {
	Provider string `yaml:"provider"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Pass     string `yaml:"pass"`
}

type Store struct {
	Provider string `yaml:"provider"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Path     string `yaml:"path"`
	User     string `yaml:"user"`
	Pass     string `yaml:"pass"`
}

type Tool struct {
	Name string
}
