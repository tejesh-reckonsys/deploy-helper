package cloudfront

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tejesh-reckonsys/deploy-helper/pkg/aws"
	"github.com/tejesh-reckonsys/deploy-helper/pkg/aws/cloudfront"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the present cloudfront distributions available",
	Long:  `List the distribution id and cnames present in cloudfront.`,
	Run: func(cmd *cobra.Command, args []string) {
		alias := cmd.Flag("alias").Value.String()

		dists := cloudfront.ListDistributions(aws.GetDefaultConfig(), alias)
		if len(dists) == 0 {
			fmt.Println("No matching distributions found")
			os.Exit(1)
		}

		fmt.Println("ID\t\tCNames")
		for _, info := range dists {
			fmt.Printf("%s\t%s\n", info.Id, strings.Join(info.Aliases, ", "))
		}
	},
}

func init() {
	CloudfrontCmd.AddCommand(listCmd)

	listCmd.Flags().StringP("alias", "a", "", "cname/alias to filter with")
}
