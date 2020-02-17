package lib

import (
	"fmt"
	"os"
)

type Config struct {
	Username string
	Password string
	Host     string
}

var (
	c Config
)

func (config *Config) toDSN() string {
	return fmt.Sprintf("amqp://%s:%s@%s", config.Username, config.Password, config.Host)
}

func LoadConfig() {
	c.Username = os.Getenv("RMQ_USERNAME")
	c.Password = os.Getenv("RMQ_PASSWORD")
	c.Host = os.Getenv("RMQ_HOST")
}
