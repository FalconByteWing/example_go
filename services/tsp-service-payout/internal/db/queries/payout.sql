-- name: CreatePayout :one
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
) RETURNING *;

-- name: GetPayoutByOrderID :one
SELECT * FROM payouts
WHERE order_id = $1 LIMIT 1;

-- name: UpdatePayoutStatus :one
UPDATE payouts
SET 
    status = $2,
    updated_at = NOW()
WHERE order_id = $1
RETURNING *;

-- name: GetMerchantBalance :one
SELECT COALESCE(balance, 0) as balance
FROM merchant_balances
WHERE merchant_id = $1; 