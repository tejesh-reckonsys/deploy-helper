package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/tejesh-reckonsys/deploy-helper/cmd/cloudfront"
	"github.com/tejesh-reckonsys/deploy-helper/config"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "deploy-helper",
	Short: "Helper for various tasks related to deployment",
	Long:  "This program provides many helper commands for deployment.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.LoadDefault(cfgFile)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(cloudfront.CloudfrontCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "x", "config file (default is $HOME/.deploy-helper.yaml)")
}
