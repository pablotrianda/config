package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Conf struct {
	ConfigFilePath string
}

// Get the value on type string only
func (c *Conf) Get(key string) string {
	return "thing"
}

func Load(configFilePath string) map[interface{}]interface{} {

	file, _ := ioutil.ReadFile(configFilePath)

	m := make(map[interface{}]interface{})

	_ = yaml.Unmarshal([]byte(file), &m)

	return m
}
