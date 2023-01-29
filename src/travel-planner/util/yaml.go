package util

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)
type ChatGPTConfig struct {
   Key string `yaml:"key"`
}
type MySQLInfo struct {
	Endpoint string `yaml:"endpoint"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type TokenInfo struct {
	Secret string `yaml:"secret"`
}

type ApplicationConfig struct {
	MySQLConfig *MySQLInfo `yaml:"mysql"`
	TokenConfig *TokenInfo `yaml:"token"`
	ChatGPTConfig *ChatGPTConfig `yaml:"ChatGPT"`
}


func LoadApplicationConfig(configDir, configFile string) (*ApplicationConfig, error) {
	content, err := ioutil.ReadFile(filepath.Join(configDir, configFile))
	if err != nil {
		return nil, err
	}

	var config ApplicationConfig
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
