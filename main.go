package main

import (
	"atom/modules"
	"os"

	"github.com/rogeecn/atom"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/providers/http"
	"github.com/rogeecn/atom/providers/http/gin"
	"github.com/rogeecn/atom/providers/log"
	"github.com/rogeecn/atom/services"
	"github.com/rogeecn/atom/utils/opt"
	"github.com/spf13/cobra"
)

var providers = container.Providers{
	{
		Provider: log.Provide,
		Options: []opt.Option{
			opt.Prefix(log.DefaultPrefix),
		},
	},
	{
		Provider: gin.Provide,
		Options: []opt.Option{
			opt.Prefix(http.DefaultPrefix),
		},
	},
	{
		Provider: modules.Provide,
	},
}

func main() {
	opts := []atom.Option{
		atom.Name("http"),
		atom.RunE(func(cmd *cobra.Command, args []string) error {
			return services.ServeHttp()
		}),
	}

	if err := atom.Serve(providers, opts...); err != nil {
		os.Exit(1)
	}
}
