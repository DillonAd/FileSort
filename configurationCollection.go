package main

//ConfigurationCollection - Collection of configurations
type ConfigurationCollection struct {
	configurations []Config
}

//GetConfigurations - Returns array of all configurations

//GetMatchingConfigurations - Returns the all configurations to match the file name
func (cc *ConfigurationCollection) GetMatchingConfigurations(fileName string) ([]Config, error) {
	var configs []Config
	var err error
	var matched bool

	for _, cfg := range cc.configurations {
		matched, err = cfg.IsMatch(fileName)

		if matched {
			configs = append(configs, cfg)
		}
	}

	return configs, err
}
