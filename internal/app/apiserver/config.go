package apiserver

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log-level"`
}

func NewConfig() *Config {
	return &Config {
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}