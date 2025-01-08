package main

import (
	"context"
	"log"

	"example_go/services/tsp-service-payout/generated/restapi"
	"example_go/services/tsp-service-payout/generated/restapi/operations"
	"example_go/services/tsp-service-payout/generated/restapi/operations/payout"
	"example_go/services/tsp-service-payout/internal/handlers"

	"github.com/go-openapi/loads"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/cors"
)

func main() {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	// Create API with the correct name as generated
	api := operations.NewTspPayoutServiceAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	// Configure CORS
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		Debug:          false,
	})

	// Initialize database connection

	dbpool, err := pgxpool.New(context.Background(), "postgres://postgres:postgres@localhost:5432/example_db?sslmode=disable")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer dbpool.Close()

	// Initialize handlers
	h := handlers.NewHandler(dbpool)

	// Register handlers
	api.PayoutCreatePayoutHandler = payout.CreatePayoutHandlerFunc(h.CreatePayout)
	api.PayoutCheckOrderStatusHandler = payout.CheckOrderStatusHandlerFunc(h.CheckOrderStatus)
	api.PayoutCheckBalanceHandler = payout.CheckBalanceHandlerFunc(h.CheckBalance)

	// Apply CORS middleware
	server.SetHandler(corsMiddleware.Handler(api.Serve(nil)))

	// Configure server
	server.Port = 8083
	server.Host = "0.0.0.0"
	server.TLSPort = 0
	server.EnabledListeners = []string{"http"}

	// Start server
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
