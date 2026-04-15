package models

type CreateTransactionRequest struct {
	FullDir     *string
	Dir         string
	File        string
	Transaction interface{}
}

type GetTransactionRequest struct {
	FullDir *string
	Dir     string
	File    string
}

type TransactionHistoryRequest struct {
	StartDate *string
	EndDate   *string
}

type CreateTransactionResponse struct {
	URL    string            `json:"url"`
	Path   string            `json:"path"`
	Status string            `json:"status"`
	Error  map[string]string `json:"error,omitempty"`
}

type TransactionResponse struct {
	Transaction interface{}
}

type ListTransactionResponse struct {
	Paging  *Paging   `json:"paging"`
	Results *[]Result `json:"results"`
}

type Result struct {
	ID        *string      `json:"id,omitempty"`
	Path      *string      `json:"path,omitempty"`
	FullURL   *string      `json:"url,omitempty"`
	Type      *string      `json:"type,omitempty"`
	StoreMode *string      `json:"store_mode,omitempty"`
	Direction *string      `json:"direction,omitempty"`
	StatusLog *[]StatusLog `json:"status_log,omitempty"`
}

type StatusLog struct {
	Status    *string `json:"status,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
}

type Paging struct {
	Offset     *int `json:"offset,omitempty"`
	Limit      *int `json:"limit,omitempty"`
	TotalCount *int `json:"total_count,omitempty"`
}
