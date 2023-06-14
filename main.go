//go:generate atomctl gen routes
//go:generate swag fmt
//go:generate swag init -ot json
package main

import (
	"log"

	"atom/http/database/migrations"
	"atom/http/database/query"
	"atom/http/database/seeders"
	"atom/http/modules/boot"
	"atom/http/modules/system"

	"github.com/rogeecn/atom"
	"github.com/rogeecn/atom-addons/providers/database/sqlite"
	"github.com/rogeecn/atom-addons/providers/faker"
	"github.com/rogeecn/atom-addons/providers/swagger"
	"github.com/rogeecn/atom-addons/services/http"
	"github.com/spf13/cobra"
)

func main() {
	providers := http.Default(
		sqlite.DefaultProvider(),
		swagger.DefaultProvider(),
		faker.DefaultProvider(),
		query.DefaultProvider(),
	).With(
		boot.Providers(),
		system.Providers(),
	)
	// providers := atom.DefaultGRPC().With(greet.Providers())

	opts := []atom.Option{
		atom.RunE(func(cmd *cobra.Command, args []string) error {
			return http.Serve()
			// return services.ServeGrpc()
		}),
		atom.Seeders(seeders.Seeders...),
		atom.Migrations(migrations.Migrations...),
	}

	if err := atom.Serve(providers, opts...); err != nil {
		log.Fatal(err)
	}
}
