package config

var (
	Build   string
	Commit  string
	Version string
)

type Config struct {
	Cachex Cachex `yaml:"cachex"`
	Ragx   Ragx   `yaml:"ragx"`
}

type Cachex struct {
	Url string `yaml:"url"`
}

type Ragx struct {
	Url string `yaml:"url"`
}
