package main

type CreateTransactionRequest struct {
	FullDir     *string
	Dir         string
	File        string
	Transaction Order
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
	PurchaseOrder Order
	*ListTransactionResponse
}

type ListTransactionResponse struct {
	Paging  Paging   `json:"paging"`
	Results []Result `json:"results"`
}

type Order struct {
	Header   Header     `json:"Header"`
	LineItem []LineItem `json:"LineItem"`
	Summary  Summary    `json:"Summary"`
}

type Header struct {
	Address            []Address            `json:"Address"`
	CarrierInformation []CarrierInformation `json:"CarrierInformation"`
	Dates              []Date               `json:"Dates"`
	OrderHeader        OrderHeader          `json:"OrderHeader"`
	PaymentTerms       []PaymentTerms       `json:"PaymentTerms"`
}

type Address struct {
	Address1              string `json:"Address1"`
	Address2              string `json:"Address2,omitempty"`
	AddressLocationNumber string `json:"AddressLocationNumber,omitempty"`
	AddressName           string `json:"AddressName"`
	AddressTypeCode       string `json:"AddressTypeCode"`
	City                  string `json:"City"`
	Country               string `json:"Country,omitempty"`
	LocationCodeQualifier string `json:"LocationCodeQualifier,omitempty"`
	PostalCode            string `json:"PostalCode"`
	State                 string `json:"State"`
}

type CarrierInformation struct {
	ServiceLevelCodes []ServiceLevelCode `json:"ServiceLevelCodes"`
}

type ServiceLevelCode struct {
	ServiceLevelCode string `json:"ServiceLevelCode"`
}

type Date struct {
	Date              string `json:"Date"`
	DateTimeQualifier string `json:"DateTimeQualifier"`
}

type OrderHeader struct {
	PrimaryPOTypeCode   string `json:"PrimaryPOTypeCode"`
	PurchaseOrderDate   string `json:"PurchaseOrderDate"`
	PurchaseOrderNumber string `json:"PurchaseOrderNumber"`
	TradingPartnerId    string `json:"TradingPartnerId"`
	TsetPurposeCode     string `json:"TsetPurposeCode"`
	Vendor              string `json:"Vendor"`
}

type PaymentTerms struct {
	TermsDescription string `json:"TermsDescription"`
}

type LineItem struct {
	OrderLine                OrderLine                  `json:"OrderLine"`
	PriceInformation         []PriceInformation         `json:"PriceInformation"`
	ProductOrItemDescription []ProductOrItemDescription `json:"ProductOrItemDescription"`
}

type OrderLine struct {
	BuyerPartNumber         string  `json:"BuyerPartNumber"`
	ConsumerPackageCode     string  `json:"ConsumerPackageCode"`
	LineSequenceNumber      string  `json:"LineSequenceNumber"`
	OrderQty                float64 `json:"OrderQty"`
	OrderQtyUOM             string  `json:"OrderQtyUOM"`
	ProductColorDescription string  `json:"ProductColorDescription"`
	PurchasePrice           float64 `json:"PurchasePrice"`
	VendorPartNumber        string  `json:"VendorPartNumber"`
}

type PriceInformation struct {
	PriceTypeIDCode string  `json:"PriceTypeIDCode"`
	UnitPrice       float64 `json:"UnitPrice"`
}

type ProductOrItemDescription struct {
	ProductCharacteristicCode string `json:"ProductCharacteristicCode"`
	ProductDescription        string `json:"ProductDescription"`
}

type Summary struct {
	TotalLineItemNumber int `json:"TotalLineItemNumber"`
}

type Result struct {
	ID        string      `json:"id"`
	Path      string      `json:"path"`
	FullURL   string      `json:"url"`
	Type      string      `json:"type"`
	StoreMode string      `json:"store_mode"`
	Direction string      `json:"direction"`
	StatusLog []StatusLog `json:"status_log"`
}

type StatusLog struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

type Paging struct {
	Offset     int `json:"offset"`
	Limit      int `json:"limit"`
	TotalCount int `json:"total_count"`
}
