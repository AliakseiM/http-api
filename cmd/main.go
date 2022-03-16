package main

import (
	"context"
	"log"

	"http-api/internal/handler"
	"http-api/internal/restapi"
	"http-api/internal/restapi/operations"

	"github.com/go-openapi/loads"
	"golang.org/x/sync/errgroup"
)

func main() {
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

		server.Host = "0.0.0.0"
		server.Port = 8080

		return server.Serve()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
