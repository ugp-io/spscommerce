package spscommerce

import (
	"fmt"
	"strings"
)

const baseURL = "https://api.spscommerce.com/transactions"

type TransactionServiceOp struct {
	client *Client
}

type TransactionService interface {
	GetTransaction(GetTransactionRequest) (*TransactionResponse, error)
	TransactionHistory(TransactionHistoryRequest) (*ListTransactionResponse, error)
	CreateTransaction(CreateTransactionRequest) (*CreateTransactionResponse, error)
	ListTransaction(GetTransactionRequest) (*ListTransactionResponse, error)
}

func (s *TransactionServiceOp) GetTransaction(req GetTransactionRequest) (*TransactionResponse, error) {

	var po Order
	if req.FullDir != nil && *req.FullDir != "" {
		if !strings.Contains(*req.FullDir, baseURL) {
			*req.FullDir = baseURL + "/v5/data/" + *req.FullDir
		}
		if err := s.client.Request("GET", *req.FullDir, nil, &po); err != nil {
			return nil, err
		}
	} else if req.File != "" && req.Dir != "" {
		url := baseURL + "/v5/data/" + req.Dir + "/" + req.File
		if err := s.client.Request("GET", url, nil, &po); err != nil {
			return nil, err
		}
	}

	return &TransactionResponse{PurchaseOrder: po}, nil
}

func (s *TransactionServiceOp) TransactionHistory(req TransactionHistoryRequest) (*ListTransactionResponse, error) {

	var parts []string
	if req.StartDate != nil {
		parts = append(parts, "after="+*req.StartDate)
	}
	if req.EndDate != nil {
		parts = append(parts, "until="+*req.EndDate)
	}

	var history ListTransactionResponse
	if err := s.client.Request("GET", fmt.Sprintf("%v/v5/history?%v", baseURL, strings.Join(parts, "&")), nil, &history); err != nil {
		return nil, err
	}

	return &history, nil
}

func (s *TransactionServiceOp) CreateTransaction(req CreateTransactionRequest) (*CreateTransactionResponse, error) {

	var resp CreateTransactionResponse
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

func (s *TransactionServiceOp) ListTransaction(req GetTransactionRequest) (*ListTransactionResponse, error) {
	var resp ListTransactionResponse
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
