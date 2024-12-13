package cloudfront

import (
	"github.com/spf13/cobra"
)

// cloudfrontCmd represents the cloudfront command
var CloudfrontCmd = &cobra.Command{
	Use:     "cloudfront",
	Aliases: []string{"cf"},
	Short:   "Cloudfront related helper commands",
	Long:    `Cloudfront related helper commands. For things like listing distributions based on cname, or invalidating cache`,
}
