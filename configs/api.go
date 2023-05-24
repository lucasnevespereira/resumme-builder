package configs

type ApiConfig struct {
	AppName string
	Port    int
}

func LoadApiConfig() ApiConfig {
	conf := ApiConfig{}
	conf.AppName = "Resumme Builder"
	conf.Port = 9000
	return conf
}
