package cloudfront

import (
	"context"
	"log"

	awsCloudFront "github.com/aws/aws-sdk-go-v2/service/cloudfront"
)

type DistInfo struct {
	Id      string
	Aliases []string
}

func iterateOverDistributions(client *awsCloudFront.Client, handler func(DistInfo) bool) {
	var nextMarker *string
	for {
		distOutput, err := client.ListDistributions(context.TODO(), &awsCloudFront.ListDistributionsInput{Marker: nextMarker})
		if err != nil {
			log.Fatalln("Error listing cloudfront distributions:", err)
		}

		for _, distItem := range distOutput.DistributionList.Items {
			item := DistInfo{*distItem.Id, distItem.Aliases.Items}
			if handler(item) {
				return
			}
		}

		if distOutput.DistributionList.IsTruncated == nil || distOutput.DistributionList.NextMarker == nil {
			return
		}

		nextMarker = distOutput.DistributionList.NextMarker
	}
}
