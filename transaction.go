package spscommerce

import (
	"fmt"
	"net/url"
)

const baseURL = "https://api.spscommerce.com/transactions"

type TransactionServiceOp struct {
	client *Client
}

type TransactionService interface {
	ListTransaction(ListTransactionRequest) (*ListTransactionResponse, error)
	GetTransaction(GetTransactionRequest) (*GetTransactionResponse, error)
	TransactionHistory(TransactionHistoryRequest) (*ListTransactionResponse, error)
}

func (s *TransactionServiceOp) GetTransaction(req GetTransactionRequest) (*GetTransactionResponse, error) {

	var po Order
	url := baseURL + "/v5/data/"
	if req.FullDir != nil {
		url += *req.FullDir
		if err := s.client.Request("GET", url, nil, &po); err != nil {
			return nil, err
		}
	} else if req.File != nil && req.Dir != nil {
		url += *req.Dir + "/" + *req.File
		if err := s.client.Request("GET", url, nil, &po); err != nil {
			return nil, err
		}
	}

	return &GetTransactionResponse{PurchaseOrder: po}, nil
}

func (s *TransactionServiceOp) TransactionHistory(req TransactionHistoryRequest) (*ListTransactionResponse, error) {

	params := url.Values{}
	if req.StartDate != nil {
		params.Set("after", *req.StartDate)
	}
	if req.EndDate != nil {
		params.Set("until", *req.EndDate)
	}

	var history ListTransactionResponse
	if err := s.client.Request("GET", fmt.Sprintf("%v/v5/history?%v", baseURL, params.Encode()), nil, &history); err != nil {
		return nil, err
	}

	return &history, nil
}

func (s *TransactionServiceOp) ListTransaction(req ListTransactionRequest) (*ListTransactionResponse, error) {
	return nil, nil
}
