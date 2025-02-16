package config

// 默认日志信息
type Log struct {
	Path       string     `yaml:"path"`
	Level      string     `yaml:"level"`
	FilePrefix string     `yaml:"filePrefix"`
	FileFormat string     `yaml:"fileFormat"`
	OutFormat  string     `yaml:"outFormat"`
	LumberJack LumberJack `yaml:"lumberJack"`
}

// 日志切割信息
type LumberJack struct {
	MaxSize    int  `yaml:"maxSize"`    //单文件最大容量(单位MB)
	MaxBackups int  `yaml:"maxBackups"` // 保留旧文件的最大数量
	MaxAge     int  `yaml:"maxAge"`     // 旧文件最多保存几天
	Compress   bool `yaml:"compress"`   // 是否压缩/归档旧文件
}
