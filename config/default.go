package config

var DefaultConfig Config

func LoadDefault(configFile string) {
	config := New(configFile)
	DefaultConfig = config
}
