package cloudfront

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"time"

	awsCloudFront "github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/tejesh-reckonsys/deploy-helper/pkg/aws"
)

type InvaldiationProgress struct {
	Status string
	Error  error
}

func InvalidateCache(cfg aws.AWSConfig, distId string, paths []string, progressCh chan InvaldiationProgress) {
	defer close(progressCh)

	client := awsCloudFront.NewFromConfig(cfg.Config)

	callerRef := getCallerRef(distId, paths)
	quantity := int32(len(paths))
	invalidation, err := client.CreateInvalidation(context.Background(), &awsCloudFront.CreateInvalidationInput{
		DistributionId: &distId,
		InvalidationBatch: &types.InvalidationBatch{
			Paths:           &types.Paths{Quantity: &quantity, Items: paths},
			CallerReference: &callerRef,
		},
	})

	if err != nil {
		progressCh <- InvaldiationProgress{Error: err}
		return
	}

	invalidationId := *invalidation.Invalidation.Id
	invalidationStatus := *invalidation.Invalidation.Status

	progressCh <- InvaldiationProgress{Status: "Invalidation request created"}

	for invalidationStatus == "InProgress" {
		time.Sleep(time.Second)
		response, err := client.GetInvalidation(context.Background(), &awsCloudFront.GetInvalidationInput{
			DistributionId: &distId,
			Id:             &invalidationId,
		})
		if err != nil {
			progressCh <- InvaldiationProgress{Error: errors.New("failed to fetch invalidation status")}
			return
		}

		invalidationStatus = *response.Invalidation.Status
		progressCh <- InvaldiationProgress{Status: invalidationStatus}
	}

	progressCh <- InvaldiationProgress{Status: "Invaldiation completed"}
}

func getCallerRef(distId string, paths []string) string {
	data := make([]byte, 32)
	rand.Read(data)
	return fmt.Sprintf("invalidate-%s-%s-%x", distId, paths[0], data)
}
