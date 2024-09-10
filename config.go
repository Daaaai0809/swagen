package swagen

import (
	"errors"
	"os"
)

const VERSION string = "0.0.1"

const (
	PATH_DIR    = "PATH_DIR"
	MESSAGE_DIR = "MESSAGE_DIR"
	MODEL_DIR   = "MODEL_DIR"
)

var config *Config

type Config struct {
	PathDir    string
	MessageDir string
	ModelDir   string
}

func (c *Config) GetPathDir() string {
	return c.PathDir
}

func (c *Config) GetMessageDir() string {
	return c.MessageDir
}

func (c *Config) GetModelDir() string {
	return c.ModelDir
}

func NewConfig() (*Config, error) {
	pathMessage := os.Getenv(PATH_DIR)
	if pathMessage == "" {
		return nil, errors.New("PATH_DIR is not set")
	}

	messageDir := os.Getenv(MESSAGE_DIR)
	if messageDir == "" {
		return nil, errors.New("MESSAGE_DIR is not set")
	}

	modelDir := os.Getenv(MODEL_DIR)
	if modelDir == "" {
		return nil, errors.New("MODEL_DIR is not set")
	}

	return &Config{
		PathDir:    pathMessage,
		MessageDir: messageDir,
		ModelDir:   modelDir,
	}, nil
}

func GetConfig() *Config {
	return config
}

func SetConfig(c *Config) {
	config = c
}
