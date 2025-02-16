package config

// 应用信息
type app struct {
	Desc       string `yaml:"desc"`
	Addr       string `yaml:"addr"`
	ConfigFile string `yaml:"configFile"`
	Version    string `yaml:"version"`
	Env        string `yaml:"env"`
}

// ServerConfig 配置信息
type ServerConfig struct {
	App     app   `yaml:"app"`
	Mysql   mysql `yaml:"mysql"`
	Redis   redis `yaml:"redis"`
	Log     Log   `yaml:"log"`
	UserLog Log   `yaml:"userLog"`
}
