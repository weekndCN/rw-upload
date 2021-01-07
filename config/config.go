package config

import (
	"log"
	"sync"

	yaml "gopkg.in/yaml.v3"
)

// Config upload server config
type Config struct {
	config map[string]UploadConfig
	lock   sync.RWMutex
}

// UploadConfig .
type UploadConfig struct {
	Name      string   `yaml:"name"`
	Directory string   `yaml:"directory"`
	BodyLimit string   `yaml:"bodyLimit"`
	FileLimit string   `yaml:"fileLimit"`
	Types     []string `yaml:"types"`
}

// SetConfig set the config from the yaml
func (c *Config) SetConfig(data []byte) error {
	var rawConfig map[string]UploadConfig

	if err := yaml.Unmarshal(data, &rawConfig); err != nil {
		log.Println(err)
		return err
	}

	// get lock
	c.lock.Lock()
	defer c.lock.Unlock()

	c.config = rawConfig
	return nil
}

// GetAll get all config from database yaml file
func (c *Config) GetAll() interface{} {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.config
}
