// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: payout.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createPayout = `-- name: CreatePayout :one
INSERT INTO payouts (
    order_id,
    bank_name,
    user_id,
    amount,
    email,
    phone,
    ifsc_code,
    account_no,
    name,
    notify_url,
    status,
    merchant_id
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
) RETURNING id, order_id, bank_name, user_id, amount, email, phone, ifsc_code, account_no, name, notify_url, status, merchant_id, created_at, updated_at
`

type CreatePayoutParams struct {
	OrderID    string         `json:"order_id"`
	BankName   string         `json:"bank_name"`
	UserID     string         `json:"user_id"`
	Amount     pgtype.Numeric `json:"amount"`
	Email      string         `json:"email"`
	Phone      string         `json:"phone"`
	IfscCode   string         `json:"ifsc_code"`
	AccountNo  string         `json:"account_no"`
	Name       string         `json:"name"`
	NotifyUrl  string         `json:"notify_url"`
	Status     PayoutStatus   `json:"status"`
	MerchantID uuid.UUID      `json:"merchant_id"`
}

func (q *Queries) CreatePayout(ctx context.Context, arg CreatePayoutParams) (Payout, error) {
	row := q.db.QueryRow(ctx, createPayout,
		arg.OrderID,
		arg.BankName,
		arg.UserID,
		arg.Amount,
		arg.Email,
		arg.Phone,
		arg.IfscCode,
		arg.AccountNo,
		arg.Name,
		arg.NotifyUrl,
		arg.Status,
		arg.MerchantID,
	)
	var i Payout
	err := row.Scan(
		&i.ID,
		&i.OrderID,
		&i.BankName,
		&i.UserID,
		&i.Amount,
		&i.Email,
		&i.Phone,
		&i.IfscCode,
		&i.AccountNo,
		&i.Name,
		&i.NotifyUrl,
		&i.Status,
		&i.MerchantID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getMerchantBalance = `-- name: GetMerchantBalance :one
SELECT COALESCE(balance, 0) as balance
FROM merchant_balances
WHERE merchant_id = $1
`

func (q *Queries) GetMerchantBalance(ctx context.Context, merchantID uuid.UUID) (pgtype.Numeric, error) {
	row := q.db.QueryRow(ctx, getMerchantBalance, merchantID)
	var balance pgtype.Numeric
	err := row.Scan(&balance)
	return balance, err
}

const getPayoutByOrderID = `-- name: GetPayoutByOrderID :one
SELECT id, order_id, bank_name, user_id, amount, email, phone, ifsc_code, account_no, name, notify_url, status, merchant_id, created_at, updated_at FROM payouts
WHERE order_id = $1 LIMIT 1
`

func (q *Queries) GetPayoutByOrderID(ctx context.Context, orderID string) (Payout, error) {
	row := q.db.QueryRow(ctx, getPayoutByOrderID, orderID)
	var i Payout
	err := row.Scan(
		&i.ID,
		&i.OrderID,
		&i.BankName,
		&i.UserID,
		&i.Amount,
		&i.Email,
		&i.Phone,
		&i.IfscCode,
		&i.AccountNo,
		&i.Name,
		&i.NotifyUrl,
		&i.Status,
		&i.MerchantID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updatePayoutStatus = `-- name: UpdatePayoutStatus :one
UPDATE payouts
SET 
    status = $2,
    updated_at = NOW()
WHERE order_id = $1
RETURNING id, order_id, bank_name, user_id, amount, email, phone, ifsc_code, account_no, name, notify_url, status, merchant_id, created_at, updated_at
`

type UpdatePayoutStatusParams struct {
	OrderID string       `json:"order_id"`
	Status  PayoutStatus `json:"status"`
}

func (q *Queries) UpdatePayoutStatus(ctx context.Context, arg UpdatePayoutStatusParams) (Payout, error) {
	row := q.db.QueryRow(ctx, updatePayoutStatus, arg.OrderID, arg.Status)
	var i Payout
	err := row.Scan(
		&i.ID,
		&i.OrderID,
		&i.BankName,
		&i.UserID,
		&i.Amount,
		&i.Email,
		&i.Phone,
		&i.IfscCode,
		&i.AccountNo,
		&i.Name,
		&i.NotifyUrl,
		&i.Status,
		&i.MerchantID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
