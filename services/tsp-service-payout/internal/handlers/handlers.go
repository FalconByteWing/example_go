package handlers

import (
	"context"
	"fmt"

	"example_go/services/tsp-service-payout/generated/models"
	"example_go/services/tsp-service-payout/generated/restapi/operations/payout"
	"example_go/services/tsp-service-payout/internal/db"

	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	queries *db.Queries
}

func NewHandler(dbPool *pgxpool.Pool) *Handler {
	return &Handler{
		queries: db.New(dbPool),
	}
}

func (h *Handler) CreatePayout(params payout.CreatePayoutParams) middleware.Responder {
	// Validate input parameters
	if params.Request == nil || params.Request.Amount == nil {
		return payout.NewCreatePayoutUnprocessableEntity().WithPayload(&models.Error{
			Code:    422,
			Message: "Amount is required",
		})
	}

	// Validate amount value
	if *params.Request.Amount <= 0 {
		return payout.NewCreatePayoutUnprocessableEntity().WithPayload(&models.Error{
			Code:    422,
			Message: "Amount must be greater than 0",
		})
	}

	// Convert amount to pgtype.Numeric
	var amount pgtype.Numeric
	amount.Valid = true
	if err := amount.Scan(fmt.Sprintf("%.2f", *params.Request.Amount)); err != nil {
		return payout.NewCreatePayoutUnprocessableEntity().WithPayload(&models.Error{
			Code:    422,
			Message: fmt.Sprintf("Invalid amount format: %v", err),
		})
	}

	// Create payout record
	payoutRecord, err := h.queries.CreatePayout(context.Background(), db.CreatePayoutParams{
		OrderID:    *params.Request.OrderID,
		BankName:   *params.Request.BankName,
		UserID:     *params.Request.UserID,
		Amount:     amount,
		Email:      string(*params.Request.Email),
		Phone:      *params.Request.Phone,
		IfscCode:   *params.Request.IfscCode,
		AccountNo:  *params.Request.AccountNo,
		Name:       *params.Request.Name,
		NotifyUrl:  params.Request.NotifyURL.String(),
		Status:     db.PayoutStatusPending,
		MerchantID: uuid.New(),
	})

	if err != nil {
		return payout.NewCreatePayoutUnprocessableEntity().WithPayload(&models.Error{
			Code:    422,
			Message: fmt.Sprintf("Failed to create payout: %v", err),
		})
	}

	return payout.NewCreatePayoutOK().WithPayload(&models.PayoutResponse{
		Status:  200,
		Message: "success",
		Data: &models.PayoutResponseData{
			Status:  string(payoutRecord.Status),
			OrderID: payoutRecord.OrderID,
			Message: "Payment initiated",
		},
	})
}

func (h *Handler) CheckOrderStatus(params payout.CheckOrderStatusParams) middleware.Responder {
	payoutRecord, err := h.queries.GetPayoutByOrderID(context.Background(), *params.Request.OrderID)
	if err != nil {
		return payout.NewCheckOrderStatusUnprocessableEntity().WithPayload(&models.Error{
			Code:    422,
			Message: "Order not found",
		})
	}

	amount, err := payoutRecord.Amount.Float64Value()
	if err != nil {
		return payout.NewCheckOrderStatusUnprocessableEntity().WithPayload(&models.Error{
			Code:    422,
			Message: "Invalid amount format",
		})
	}

	return payout.NewCheckOrderStatusOK().WithPayload(&models.PayoutStatusResponse{
		Status:  200,
		Message: "success",
		Data: &models.PayoutStatusResponseData{
			OrderID:       payoutRecord.ID,
			TransactionID: payoutRecord.ID,
			Amount:        amount.Float64,
			Status:        string(payoutRecord.Status),
			Message:       "Payment status retrieved",
		},
	})
}

func (h *Handler) CheckBalance(params payout.CheckBalanceParams) middleware.Responder {
	merchantID, err := uuid.Parse(params.MerchantID.String())
	if err != nil {
		return payout.NewCheckBalanceUnprocessableEntity().WithPayload(&models.Error{
			Code:    422,
			Message: "Invalid merchant ID",
		})
	}

	balance, err := h.queries.GetMerchantBalance(context.Background(), merchantID)
	if err != nil {
		return payout.NewCheckBalanceUnprocessableEntity().WithPayload(&models.Error{
			Code:    422,
			Message: err.Error(),
		})
	}

	balanceValue, err := balance.Float64Value()
	if err != nil {
		return payout.NewCheckBalanceUnprocessableEntity().WithPayload(&models.Error{
			Code:    422,
			Message: "Invalid balance format",
		})
	}

	balanceData := &models.BalanceResponseDataItems0{
		Status:  "success",
		Message: "",
		Balance: balanceValue.Float64,
	}

	return payout.NewCheckBalanceOK().WithPayload(&models.BalanceResponse{
		Status:  200,
		Message: "get user balance",
		Data:    []*models.BalanceResponseDataItems0{balanceData},
	})
}
