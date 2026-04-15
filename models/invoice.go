package models

type Invoice struct {
	Header   *Header     `json:"Header,omitempty"`
	LineItem *[]LineItem `json:"LineItem,omitempty"`
	Summary  *Summary    `json:"Summary,omitempty"`
}

type InvoiceHeader struct {
	TradingPartnerId         *string                 `json:"TradingPartnerId,omitempty"`
	InvoiceNumber            *string                 `json:"InvoiceNumber,omitempty"`
	InvoiceDate              *string                 `json:"InvoiceDate,omitempty"`
	InvoiceTime              *string                 `json:"InvoiceTime,omitempty"`
	PurchaseOrderDate        *string                 `json:"PurchaseOrderDate,omitempty"`
	PurchaseOrderNumber      *string                 `json:"PurchaseOrderNumber,omitempty"`
	PrimaryPOTypeCode        *string                 `json:"PrimaryPOTypeCode,omitempty"`
	PrimaryPOTypeDescription *string                 `json:"PrimaryPOTypeDescription,omitempty"`
	ReleaseNumber            *string                 `json:"ReleaseNumber,omitempty"`
	InvoiceTypeCode          *string                 `json:"InvoiceTypeCode,omitempty"`
	TsetPurposeCode          *string                 `json:"TsetPurposeCode,omitempty"`
	ActionCode               *string                 `json:"ActionCode,omitempty"`
	BuyersCurrency           *string                 `json:"BuyersCurrency,omitempty"`
	SellersCurrency          *string                 `json:"SellersCurrency,omitempty"`
	ExchangeRate             *float64                `json:"ExchangeRate,omitempty"`
	InternalOrderNumber      *string                 `json:"InternalOrderNumber,omitempty"`
	InternalOrderDate        *string                 `json:"InternalOrderDate,omitempty"`
	Department               *string                 `json:"Department,omitempty"`
	DepartmentDescription    *string                 `json:"DepartmentDescription,omitempty"`
	Vendor                   *string                 `json:"Vendor,omitempty"`
	JobNumber                *string                 `json:"JobNumber,omitempty"`
	Division                 *string                 `json:"Division,omitempty"`
	CustomerAccountNumber    *string                 `json:"CustomerAccountNumber,omitempty"`
	CustomerOrderNumber      *string                 `json:"CustomerOrderNumber,omitempty"`
	CarrierProNumber         *string                 `json:"CarrierProNumber,omitempty"`
	BillOfLadingNumber       *string                 `json:"BillOfLadingNumber,omitempty"`
	ShipDate                 *string                 `json:"ShipDate,omitempty"`
	ShipTime                 *string                 `json:"ShipTime,omitempty"`
	ShipDeliveryDate         *string                 `json:"ShipDeliveryDate,omitempty"`
	ShipDeliveryTime         *string                 `json:"ShipDeliveryTime,omitempty"`
	LanguageCode             *string                 `json:"LanguageCode,omitempty"`
	DocumentVersion          *string                 `json:"DocumentVersion,omitempty"`
	DocumentRevision         *string                 `json:"DocumentRevision,omitempty"`
	AdditionalPOTypeCodes    *[]AdditionalPOTypeCode `json:"AdditionalPOTypeCodes,omitempty"`
}

type InvoiceLine struct {
	LineSequenceNumber              *string                  `json:"LineSequenceNumber,omitempty"`
	ApplicationId                   *string                  `json:"ApplicationId,omitempty"`
	BuyerPartNumber                 *string                  `json:"BuyerPartNumber,omitempty"`
	VendorPartNumber                *string                  `json:"VendorPartNumber,omitempty"`
	ConsumerPackageCode             *string                  `json:"ConsumerPackageCode,omitempty"`
	EAN                             *string                  `json:"EAN,omitempty"`
	GTIN                            *string                  `json:"GTIN,omitempty"`
	UPCCaseCode                     *string                  `json:"UPCCaseCode,omitempty"`
	NatlDrugCode                    *string                  `json:"NatlDrugCode,omitempty"`
	InternationalStandardBookNumber *string                  `json:"InternationalStandardBookNumber,omitempty"`
	InvoiceQty                      *float64                 `json:"InvoiceQty,omitempty"`
	InvoiceQtyUOM                   *string                  `json:"InvoiceQtyUOM,omitempty"`
	OrderQty                        *float64                 `json:"OrderQty,omitempty"`
	OrderQtyUOM                     *string                  `json:"OrderQtyUOM,omitempty"`
	PurchasePriceType               *string                  `json:"PurchasePriceType,omitempty"`
	PurchasePrice                   *float64                 `json:"PurchasePrice,omitempty"`
	PurchasePriceBasis              *string                  `json:"PurchasePriceBasis,omitempty"`
	ShipQty                         *float64                 `json:"ShipQty,omitempty"`
	ShipQtyUOM                      *string                  `json:"ShipQtyUOM,omitempty"`
	ShipDate                        *string                  `json:"ShipDate,omitempty"`
	QtyLeftToReceive                *float64                 `json:"QtyLeftToReceive,omitempty"`
	QtyLeftToReceiveStatus          *string                  `json:"QtyLeftToReceiveStatus,omitempty"`
	ExtendedItemTotal               *float64                 `json:"ExtendedItemTotal,omitempty"`
	ProductSizeCode                 *string                  `json:"ProductSizeCode,omitempty"`
	ProductSizeDescription          *string                  `json:"ProductSizeDescription,omitempty"`
	ProductColorCode                *string                  `json:"ProductColorCode,omitempty"`
	ProductColorDescription         *string                  `json:"ProductColorDescription,omitempty"`
	ProductMaterialCode             *string                  `json:"ProductMaterialCode,omitempty"`
	ProductMaterialDescription      *string                  `json:"ProductMaterialDescription,omitempty"`
	ProductProcessCode              *string                  `json:"ProductProcessCode,omitempty"`
	ProductProcessDescription       *string                  `json:"ProductProcessDescription,omitempty"`
	Department                      *string                  `json:"Department,omitempty"`
	DepartmentDescription           *string                  `json:"DepartmentDescription,omitempty"`
	Class                           *string                  `json:"Class,omitempty"`
	SellerDateCode                  *string                  `json:"SellerDateCode,omitempty"`
	CashDiscount                    *string                  `json:"CashDiscount,omitempty"`
	SaleOrService                   *string                  `json:"SaleOrService,omitempty"`
	ProductID                       *[]ProductID             `json:"ProductID,omitempty"`
	NRFStandardColorAndSize         *NRFStandardColorAndSize `json:"NRFStandardColorAndSize,omitempty"`
}
