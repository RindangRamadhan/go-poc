package cmd

import (
	"go-poc/common/server"
	"go-poc/internal/handlers"
	"go-poc/internal/pkg/env"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "start",
	Short: "Runs the server",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func run() {
	srv, err := server.NewGinHttpRouter(env.Get().ServicePort)
	if err != nil {
		panic(err)
	}

	// container := internal.NewContainer()

	router := handlers.NewRouter(srv.Router)
	router.RegisterRouter()

	srv.Start(env.Get().ServicePort)
}
