package main

import (
	"os"

	"atom/http/database/migrations"
	"atom/http/database/seeders"
	"atom/http/modules/system"

	"github.com/rogeecn/atom"
	"github.com/rogeecn/atom/providers/database/sqlite"
	"github.com/rogeecn/atom/services"
	"github.com/spf13/cobra"
)

func main() {
	providers := atom.DefaultHTTP(sqlite.DefaultProvider()).
		With(system.Providers())

	opts := []atom.Option{
		atom.Name("http"),
		atom.RunE(func(cmd *cobra.Command, args []string) error {
			return services.ServeHttp()
		}),
		atom.Seeders(seeders.Seeders...),
		atom.Migrations(migrations.Migrations...),
	}

	if err := atom.Serve(providers, opts...); err != nil {
		os.Exit(1)
	}
}
