package config

import (
	"errors"
	"os"
)

const VERSION string = "0.0.1"

const (
	PATH_DIR   = "PATH_DIR"
	SCHEMA_DIR = "SCHEMA_DIR"
	MODEL_DIR  = "MODEL_DIR"
)

var config *Config

type Config struct {
	PathDir   string
	SchemaDir string
	ModelDir  string
}

func (c *Config) GetPathDir() string {
	return c.PathDir
}

func (c *Config) GetSchemaDir() string {
	return c.SchemaDir
}

func (c *Config) GetModelDir() string {
	return c.ModelDir
}

func NewConfig() (*Config, error) {
	pathSchema := os.Getenv(PATH_DIR)
	if pathSchema == "" {
		return nil, errors.New("PATH_DIR is not set")
	}

	schemaDir := os.Getenv(SCHEMA_DIR)
	if schemaDir == "" {
		return nil, errors.New("SCHEMA_DIR is not set")
	}

	modelDir := os.Getenv(MODEL_DIR)
	if modelDir == "" {
		return nil, errors.New("MODEL_DIR is not set")
	}

	return &Config{
		PathDir:   pathSchema,
		SchemaDir: schemaDir,
		ModelDir:  modelDir,
	}, nil
}

func GetConfig() *Config {
	return config
}

func SetConfig(c *Config) {
	config = c
}
