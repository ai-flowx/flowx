package config

var (
	Build   string
	Commit  string
	Version string
)

type Config struct {
	Cach Cache `yaml:"cache"`
	Rag  Rag   `yaml:"rag"`
}

type Cache struct {
	Url string `yaml:"url"`
}

type Rag struct {
	Url string `yaml:"url"`
}
