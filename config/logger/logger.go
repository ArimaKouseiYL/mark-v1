package logger

type LogConfig struct {
	Level      string `json:"level"`
	Mode       string `json:"mode"`
	MaxSize    int    `json:"max_size"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
}
