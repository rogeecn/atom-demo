package main

import (
	"os"

	"github.com/rogeecn/atom/cmd"
	"github.com/rogeecn/atom/services"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:     "atom",
		Short:   "atom",
		Long:    `the app long description`,
		Version: "Version",
		PreRun: func(cmd *cobra.Command, args []string) {
			bootstrap()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return services.ServeHttp()
		},
	}

	cmd.WithMigration(rootCmd)
	cmd.WithSeeder(rootCmd)
	cmd.WithModel(rootCmd)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
