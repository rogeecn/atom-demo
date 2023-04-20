package cmd

import (
	"os"

	_ "atom/modules"

	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/services"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "atom",
	Short:   "atom",
	Long:    `the app long description`,
	Version: "Version",
	RunE: func(cmd *cobra.Command, args []string) error {
		return container.Container.Invoke(services.ServeHttp)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
