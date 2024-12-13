package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func getConfigFieldError(fieldName string) error {
	return fmt.Errorf("please provide %s", fieldName)
}

type ConfigBuilder struct {
	AWS_ACCESS_KEY_ID           *string
	AWS_SECRET_ACCESS_KEY       *string
	AWS_REGION                  *string
	AWS_CLOUDFRONT_DISTRIBUTION *string
}

type yamlConfigStructure struct {
	AWS struct {
		AWS_ACCESS_KEY_ID           *string `yaml:"access_key"`
		AWS_SECRET_ACCESS_KEY       *string `yaml:"secret_key"`
		AWS_REGION                  *string `yaml:"region"`
		AWS_CLOUDFRONT_DISTRIBUTION *string `yaml:"cloudfront_dist"`
	} `yaml:"aws"`
}

func (cb *ConfigBuilder) UpdateWithEnv() {
	if value, found := os.LookupEnv("AWS_ACCESS_KEY_ID"); found {
		cb.AWS_ACCESS_KEY_ID = &value
	}
	if value, found := os.LookupEnv("AWS_SECRET_ACCESS_KEY"); found {
		cb.AWS_SECRET_ACCESS_KEY = &value
	}
	if value, found := os.LookupEnv("AWS_REGION"); found {
		cb.AWS_REGION = &value
	}
	if value, found := os.LookupEnv("AWS_CLOUDFRONT_DISTRIBUTION"); found {
		cb.AWS_CLOUDFRONT_DISTRIBUTION = &value
	}
}

func (cb *ConfigBuilder) UpdateWithYamlFile(filename string) {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return
	}

	var content yamlConfigStructure
	err = yaml.Unmarshal(fileContent, &content)
	if err != nil {
		log.Fatalln("Not able to process the config file:", err)
	}

	if content.AWS.AWS_ACCESS_KEY_ID != nil {
		cb.AWS_ACCESS_KEY_ID = content.AWS.AWS_ACCESS_KEY_ID
	}
	if content.AWS.AWS_SECRET_ACCESS_KEY != nil {
		cb.AWS_SECRET_ACCESS_KEY = content.AWS.AWS_SECRET_ACCESS_KEY
	}
	if content.AWS.AWS_REGION != nil {
		cb.AWS_REGION = content.AWS.AWS_REGION
	}
	if content.AWS.AWS_CLOUDFRONT_DISTRIBUTION != nil {
		cb.AWS_CLOUDFRONT_DISTRIBUTION = content.AWS.AWS_CLOUDFRONT_DISTRIBUTION
	}
}

func (cb *ConfigBuilder) Build() (*Config, error) {
	if cb.AWS_ACCESS_KEY_ID == nil {
		return nil, getConfigFieldError("aws access key id")
	}

	if cb.AWS_SECRET_ACCESS_KEY == nil {
		return nil, getConfigFieldError("aws secret access key")
	}
	if cb.AWS_REGION == nil {
		return nil, getConfigFieldError("aws region")
	}

	config := Config{}
	config.AWS_ACCESS_KEY_ID = *cb.AWS_ACCESS_KEY_ID
	config.AWS_SECRET_ACCESS_KEY = *cb.AWS_SECRET_ACCESS_KEY
	config.AWS_REGION = *cb.AWS_REGION
	if cb.AWS_CLOUDFRONT_DISTRIBUTION != nil {
		config.AWS_CLOUDFRONT_DISTRIBUTION = *cb.AWS_CLOUDFRONT_DISTRIBUTION
	}
	return &config, nil
}
