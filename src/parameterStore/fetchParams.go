package parameterstore

import (
	"context"
	"fmt"
	"io"
	"log"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func cleanParamName(name, prefix string) string {
	without_prefix, _ := strings.CutPrefix(name, prefix+"/")
	return strings.ToUpper(without_prefix)
}

func FetchParams(prefix string) map[string]string {
	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		log.Fatalln("Error loading aws config:", err)
	}

	client := ssm.NewFromConfig(cfg)

	response := make(map[string]string)
	var nextToken *string
	for {

		params, err := client.GetParametersByPath(context.TODO(), &ssm.GetParametersByPathInput{
			Path:      &prefix,
			NextToken: nextToken,
		})
		if err != nil {
			log.Fatalln("Error fetching params from AWS:", err)
		}

		for _, param := range params.Parameters {
			name := cleanParamName(*param.Name, prefix)
			value := *param.Value
			response[name] = value
		}

		nextToken = params.NextToken
		if nextToken == nil {
			break
		}
	}

	return response
}

func getParamEnvText(params map[string]string) string {
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	envText := ""
	for _, k := range keys {
		envText += fmt.Sprintf("%s=%s\n", k, params[k])
	}
	return envText
}

func PrintParameters(params map[string]string) {
	fmt.Print(getParamEnvText(params))
}

func WriteToEnvFile(params map[string]string, file io.Writer) {
	envText := getParamEnvText(params)

	_, err := file.Write([]byte(envText))
	if err != nil {
		log.Fatalln("Not able to write to file:", err)
	}
}
