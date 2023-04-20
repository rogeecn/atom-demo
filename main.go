package main

import (
	"os"

	"atom/modules"

	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/providers/http"
	"github.com/rogeecn/atom/providers/http/gin"
	"github.com/rogeecn/atom/providers/log"
	"github.com/rogeecn/atom/services"
	"github.com/spf13/cobra"
)

func main() {
	if err := log.Provide(&log.Config{}); err != nil {
		log.Fatal("provide http service failed, err: ", err)
	}

	if err := gin.Provide(&http.Config{Port: 8077}); err != nil {
		log.Fatal("provide http service failed, err: ", err)
	}

	if err := modules.Provide(); err != nil {
		log.Fatal("provide modules failed, err: ", err)
	}

	var rootCmd = &cobra.Command{
		Use:     "atom",
		Short:   "atom",
		Long:    `the app long description`,
		Version: "Version",
		RunE: func(cmd *cobra.Command, args []string) error {
			return container.Container.Invoke(services.ServeHttp)
		},
	}

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
