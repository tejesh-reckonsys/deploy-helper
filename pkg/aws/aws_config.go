package aws

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/tejesh-reckonsys/deploy-helper/config"
)

type AWSConfig struct {
	Config aws.Config
}

func credentialRetriever(ctx context.Context) (aws.Credentials, error) {
	return aws.Credentials{
		AccessKeyID:     config.DefaultConfig.AWS_ACCESS_KEY_ID,
		SecretAccessKey: config.DefaultConfig.AWS_SECRET_ACCESS_KEY,
	}, nil
}

func GetDefaultConfig() AWSConfig {
	cfg, err := awsConfig.LoadDefaultConfig(
		context.TODO(),
		awsConfig.WithRegion(config.DefaultConfig.AWS_REGION),
		awsConfig.WithCredentialsProvider(aws.CredentialsProviderFunc(credentialRetriever)),
	)

	if err != nil {
		log.Fatalln("Error loading aws config:", err)
	}
	return AWSConfig{Config: cfg}
}
