package certmon

import (
	"io/ioutil"
	"time"

	yaml "gopkg.in/yaml.v2"
)

//Config is the configuration for certmon, typically loaded from config file
type Config struct {
	Checks       []Check       `yaml:"checks"`
	LoopDuration time.Duration `yaml:"loop"`
}

//Check is an individual endpoint to test
type Check struct {
	Hostname  string   `yaml:"hostname"`  //The value in SNI and certificate will be validated against it.
	Endpoints []string `yaml:"endpoints"` //IP or hostname to connect to. If blank then the Hostname will be used
}

//LoadConfig loads the config from yaml file
func LoadConfig(fname string) (Config, error) {
	cfg := Config{}
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return cfg, err
	}
	err = yaml.Unmarshal(b, &cfg)
	return cfg, err
}
