package cmd

import (
	"fmt"
	"log"

	"go-poc/internal/pkg/env"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "service_poc",
	Short: "Welcome to the service poc.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the service poc.")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	env.Load()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
