package config

var (
	logger *Logger
)

func GetLogger(l string) *Logger {
	logger = NewLogger(l)
	return logger
}
