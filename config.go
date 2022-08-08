package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Conf struct {
	ConfigFilePath string
	Info           map[interface{}]interface{}
}

// Get the value on type string only
func (c *Conf) Get(key string) string {
	return fmt.Sprintf("%v", c.Info[key])
}

func Load(configFilePath string) Conf {
	conf := Conf{
		ConfigFilePath: configFilePath,
	}

	file, _ := ioutil.ReadFile(configFilePath)

	m := make(map[interface{}]interface{})
	_ = yaml.Unmarshal([]byte(file), &m)

	conf.Info = m

	return conf
}
