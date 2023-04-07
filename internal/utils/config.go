package utils

type Config struct {
	AppName string
	Port    int
}

func LoadConfig() Config {
	conf := Config{}
	conf.AppName = "Resumme Builder"
	conf.Port = 9000
	return conf
}
