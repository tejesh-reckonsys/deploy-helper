package cloudfront

import (
	"fmt"
	"os"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/tejesh-reckonsys/deploy-helper/config"
	"github.com/tejesh-reckonsys/deploy-helper/pkg/aws"
	"github.com/tejesh-reckonsys/deploy-helper/pkg/aws/cloudfront"
)

// invalidateCacheCmd represents the invalidateCache command
var invalidateCacheCmd = &cobra.Command{
	Use:     "invalidate-cache [flags] path-list",
	Aliases: []string{"invalidate"},
	Short:   "Invalidate the cache of cloudfront distribution",
	Long:    `Invalidate the cache of cloudfront distribution using the distribution id.`,
	Args:    cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		distId := cmd.Flag("dist-id").Value.String()
		if distId == "" {
			distId = config.DefaultConfig.AWS_CLOUDFRONT_DISTRIBUTION
		}

		paths := args
		if len(args) == 0 {
			paths = []string{"/*"}
		}

		// Start spinner
		spinner, _ := pterm.DefaultSpinner.Start("Starting invalidation process...")

		// Start invalidation in a goroutine
		progressCh := make(chan cloudfront.InvaldiationProgress)
		go cloudfront.InvalidateCache(aws.GetDefaultConfig(), distId, paths, progressCh)

		// Read progress updates
		for update := range progressCh {
			if update.Error != nil {
				spinner.Fail("❌ Error: " + update.Error.Error())
				os.Exit(1)
			}
			fmt.Print("\033[2K\r") // To clear the line
			spinner.UpdateText(fmt.Sprintf("Invalidation Progress: %s", update.Status))
		}

		spinner.Success("✅ Invalidation completed successfully")
	},
}

func init() {
	CloudfrontCmd.AddCommand(invalidateCacheCmd)

	invalidateCacheCmd.Flags().StringP("dist-id", "d", "", "Cloudfront distribution id")

}
