package initialize

import (
	"fmt"
	"io/ioutil"

	"github.com/thmlbdshoichi/thmproj_azure_ddns/azure-ddns-client/config"
	"gopkg.in/yaml.v3"
)

func ReadYamlFile(filename string) (*config.Config, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := &config.Config{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %w", filename, err)
	}
	return c, err
}
