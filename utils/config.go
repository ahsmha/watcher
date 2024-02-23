package utils

type Config struct {
	Token string
}

func NewConfig() *Config {
	return &Config{
		Token: "",
	}
}
