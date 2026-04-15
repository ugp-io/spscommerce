package spscommerce

import (
	"fmt"
	"strings"

	"github.com/ugp-io/spscommerce/models"
)

const baseURL = "https://api.spscommerce.com/transactions"

type TransactionServiceOp struct {
	client *Client
}

type TransactionService interface {
	GetTransaction(models.GetTransactionRequest) (*models.TransactionResponse, error)
	TransactionHistory(models.TransactionHistoryRequest) (*models.ListTransactionResponse, error)
	CreateTransaction(models.CreateTransactionRequest) (*models.CreateTransactionResponse, error)
	ListTransaction(models.GetTransactionRequest) (*models.ListTransactionResponse, error)
}

func (s *TransactionServiceOp) GetTransaction(req models.GetTransactionRequest) (*models.TransactionResponse, error) {

	var transaction interface{}
	if req.FullDir != nil && *req.FullDir != "" {
		if !strings.Contains(*req.FullDir, baseURL) {
			*req.FullDir = baseURL + "/v5/data/" + *req.FullDir
		}
		if err := s.client.Request("GET", *req.FullDir, nil, &transaction); err != nil {
			return nil, err
		}
	} else if req.File != "" && req.Dir != "" {
		url := baseURL + "/v5/data/" + req.Dir + "/" + req.File
		if err := s.client.Request("GET", url, nil, &transaction); err != nil {
			return nil, err
		}
	}

	return &models.TransactionResponse{Transaction: transaction}, nil
}

func (s *TransactionServiceOp) TransactionHistory(req models.TransactionHistoryRequest) (*models.ListTransactionResponse, error) {

	var parts []string
	if req.StartDate != nil {
		parts = append(parts, "after="+*req.StartDate)
	}
	if req.EndDate != nil {
		parts = append(parts, "until="+*req.EndDate)
	}

	var history models.ListTransactionResponse
	if err := s.client.Request("GET", fmt.Sprintf("%v/v5/history?%v", baseURL, strings.Join(parts, "&")), nil, &history); err != nil {
		return nil, err
	}

	return &history, nil
}

func (s *TransactionServiceOp) CreateTransaction(req models.CreateTransactionRequest) (*models.CreateTransactionResponse, error) {

	var resp models.CreateTransactionResponse
	if req.FullDir != nil && *req.FullDir != "" {
		if !strings.Contains(*req.FullDir, baseURL) {
			*req.FullDir = baseURL + "/v5/data/" + *req.FullDir
		}
		if err := s.client.Request("POST", *req.FullDir, req.Transaction, &resp); err != nil {
			return nil, err
		}
	} else if req.Dir != "" && req.File != "" {
		url := baseURL + "/v5/data/" + req.Dir + "/" + req.File
		if err := s.client.Request("POST", url, req.Transaction, &resp); err != nil {
			return nil, err
		}
	}

	return &resp, nil
}

func (s *TransactionServiceOp) ListTransaction(req models.GetTransactionRequest) (*models.ListTransactionResponse, error) {
	var resp models.ListTransactionResponse
	if req.FullDir != nil && *req.FullDir != "" {
		if !strings.Contains(*req.FullDir, baseURL) {
			*req.FullDir = baseURL + "/v5/data/" + *req.FullDir
		}
		if err := s.client.Request("GET", *req.FullDir, nil, &resp); err != nil {
			return nil, err
		}
	} else if req.Dir != "" {
		url := baseURL + "/v5/data/" + req.Dir
		if err := s.client.Request("GET", url, nil, &resp); err != nil {
			return nil, err
		}
	}
	return &resp, nil
}
