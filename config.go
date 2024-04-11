package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Conf struct {
	ConfigFilePath string
	Info           map[any]any
}

// Get the value on type string only
func (c *Conf) Get(key string) (string, error) {
	val, ok := c.Info[key]
	if !ok {
		return "", errors.New("Key non exist")
	}
	return fmt.Sprintf("%v", val), nil
}

func Load(configFilePath string) (Conf, error) {
	conf := Conf{
		ConfigFilePath: configFilePath,
	}

	file, err := os.ReadFile(configFilePath)
	if err != nil {
		return Conf{}, errors.New("Cannot read the file")
	}

	m := make(map[any]any)
	_ = yaml.Unmarshal([]byte(file), &m)

	conf.Info = m

	return conf, nil
}
