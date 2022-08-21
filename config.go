package config

import (
	"errors"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Conf struct {
	ConfigFilePath string
	Info           map[interface{}]interface{}
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

	file, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return Conf{}, errors.New("Cannot read the file")
	}

	m := make(map[interface{}]interface{})
	_ = yaml.Unmarshal([]byte(file), &m)

	conf.Info = m

	return conf, nil
}
