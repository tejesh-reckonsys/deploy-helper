package cloudfront

import (
	"strings"

	awsCloudFront "github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/tejesh-reckonsys/deploy-helper/pkg/aws"
)

func ListDistributions(cfg aws.AWSConfig, search string) []DistInfo {
	client := awsCloudFront.NewFromConfig(cfg.Config)

	dists := make([]DistInfo, 0)

	iterateOverDistributions(client, func(ds DistInfo) bool {
		if search == "" {
			dists = append(dists, ds)
			return false
		}

		for _, alias := range ds.Aliases {
			if strings.Contains(alias, search) {
				dists = append(dists, ds)
			}
		}
		return false
	})

	return dists

}
