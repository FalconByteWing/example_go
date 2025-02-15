// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	CreatePayout(ctx context.Context, arg CreatePayoutParams) (Payout, error)
	GetMerchantBalance(ctx context.Context, merchantID uuid.UUID) (pgtype.Numeric, error)
	GetPayoutByOrderID(ctx context.Context, orderID string) (Payout, error)
	UpdatePayoutStatus(ctx context.Context, arg UpdatePayoutStatusParams) (Payout, error)
}

var _ Querier = (*Queries)(nil)
