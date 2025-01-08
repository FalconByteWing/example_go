#!/bin/bash

# Generate Swagger files
make generate-tsp-payout-swagger

# Generate SQLC files
make generate-tsp-payout-sqlc

# Run database migrations
make migrate-tsp-payout-up

# Run the service
make run-tsp-payout 