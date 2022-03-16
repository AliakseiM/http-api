package main

import (
	"context"
	"log"

	config "http-api/internal/config"
	"http-api/internal/handler"
	"http-api/internal/restapi"
	"http-api/internal/restapi/operations"

	"github.com/go-openapi/loads"
	"golang.org/x/sync/errgroup"
)

func main() {
	cfg := config.NewConfig()

	g, _ := errgroup.WithContext(context.Background())

	g.Go(func() error {
		swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
		if err != nil {
			return err
		}

		api := operations.NewSkeletonAPI(swaggerSpec)

		api.DebugHealthHandler = handler.NewHealth()

		server := restapi.NewServer(api)
		defer server.Shutdown()

		server.Host = cfg.ApiHost
		server.Port = cfg.ApiPort

		server.ConfigureAPI()

		return server.Serve()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
