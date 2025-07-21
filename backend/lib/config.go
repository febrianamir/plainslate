package lib

type Config struct {
	RootPath string
}

func NewConfig() *Config {
	return &Config{}
}
