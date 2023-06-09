package main

import (
	"atom/http/database/migrations"
	"atom/http/database/seeders"
	"atom/http/modules/system"
	"log"

	"github.com/rogeecn/atom"
	"github.com/rogeecn/atom/providers/database/sqlite"
	"github.com/rogeecn/atom/providers/swagger"
	"github.com/rogeecn/atom/services"
	"github.com/spf13/cobra"
)

func main() {
	providers := atom.DefaultHTTP(
		sqlite.DefaultProvider(),
		swagger.DefaultProvider(),
	).With(system.Providers())
	// providers := atom.DefaultGRPC().With(greet.Providers())

	opts := []atom.Option{
		atom.RunE(func(cmd *cobra.Command, args []string) error {
			return services.ServeHttp()
			// return services.ServeGrpc()
		}),
		atom.Seeders(seeders.Seeders...),
		atom.Migrations(migrations.Migrations...),
	}

	if err := atom.Serve(providers, opts...); err != nil {
		log.Fatal(err)
	}
}
