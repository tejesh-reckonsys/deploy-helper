package config

import (
	"log"
	"os"
)

type Config struct {
	AWS_ACCESS_KEY_ID           string
	AWS_SECRET_ACCESS_KEY       string
	AWS_REGION                  string
	AWS_CLOUDFRONT_DISTRIBUTION string
}

func GetFilesToCheckConfig() []string {
	files := make([]string, 0, 2)

	// If home is set, add HOME/.deploy-helper.yaml file
	home, ok := os.LookupEnv("HOME")
	if ok {
		files = append(files, home+"/.deploy-helper.yaml")
	}

	// Also look at current dir for .deploy-helper.yaml
	files = append(files, ".deploy-helper.yaml")
	return files
}

func New(customFile string) Config {
	builder := ConfigBuilder{}
	builder.UpdateWithEnv()

	for _, file := range GetFilesToCheckConfig() {
		builder.UpdateWithYamlFile(file)
	}

	if customFile != "" {
		builder.UpdateWithYamlFile(customFile)
	}

	config, err := builder.Build()
	if err != nil {
		log.Fatalln("Error loading config:", err)
	}

	return *config
}
