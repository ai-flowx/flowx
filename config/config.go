package config

var (
	Build   string
	Commit  string
	Version string
)

type Config struct {
	Cache []Cache `yaml:"cache"`
	Gpt   []Gpt   `yaml:"gpt"`
	Store []Store `yaml:"store"`
}

type Cache struct {
	Provider string `yaml:"provider"`
	Url      string `yaml:"url"`
}

type Gpt struct {
	Provider string `yaml:"provider"`
	Url      string `yaml:"url"`
}

type Store struct {
	Provider string `yaml:"provider"`
	Url      string `yaml:"url"`
}
