/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/spf13/cobra"
	parameterstore "github.com/tejesh-reckonsys/deploy-helper/src/parameterStore"
)

// fetchEnvCmd represents the fetchEnv command
var fetchEnvCmd = &cobra.Command{
	Use:   "fetch-env [flags] ssm-path",
	Short: "Fetch environment variables from AWS SSM",
	Long: `Fetch the parameters from the AWS Parameter store starting with a prefix.
Can either output to stdin, or provide output flag to output to a file.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		outputPath := cmd.Flag("output").Value.String()
		if outputPath != "" {
			fmt.Println("Fetching parameters...")
		}

		params := parameterstore.FetchParams(args[0])
		if outputPath != "" {
			fmt.Println("Parameters fetched successfully.")
			fmt.Println("Writing to", outputPath)
		}

		var file *os.File
		var err error
		if outputPath != "" {
			file, err = os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, fs.FileMode(0664))
			if err != nil {
				log.Fatalln("Not able to open file for writing:", err)
			}
		} else {
			file = os.Stdout
		}
		parameterstore.WriteToEnvFile(params, file)
	},
}

func init() {
	rootCmd.AddCommand(fetchEnvCmd)

	fetchEnvCmd.PersistentFlags().StringP("output", "o", "", "File to output to. Ex: -o .env")
}
