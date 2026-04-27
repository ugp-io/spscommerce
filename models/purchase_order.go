package models

type Order struct {
	Header    *Header     `json:"Header,omitempty"`
	LineItems *[]LineItem `json:"LineItem,omitempty"`
	Summary   *Summary    `json:"Summary,omitempty"`
}

type Header struct {
	OrderHeader                   *OrderHeader                    `json:"OrderHeader,omitempty"`
	InvoiceHeader                 *InvoiceHeader                  `json:"InvoiceHeader,omitempty"`
	Addresses                     *[]Address                      `json:"Address,omitempty"`
	CarrierInformation            *[]CarrierInformation           `json:"CarrierInformation,omitempty"`
	Dates                         *[]Date                         `json:"Dates,omitempty"`
	PaymentTerms                  *[]PaymentTerm                  `json:"PaymentTerms,omitempty"`
	References                    *[]Reference                    `json:"References,omitempty"`
	ChargesAllowances             *[]ChargesAllowance             `json:"ChargesAllowances,omitempty"`
	Contacts                      *[]Contact                      `json:"Contacts,omitempty"`
	FOBRelatedInstructions        *[]FOBRelatedInstruction        `json:"FOBRelatedInstruction,omitempty"`
	Notes                         *[]Note                         `json:"Notes,omitempty"`
	QuantityTotals                *[]QuantityTotal                `json:"QuantityTotals,omitempty"`
	Commodity                     *[]Commodity                    `json:"Commodity,omitempty"`
	Taxes                         *[]Tax                          `json:"Taxes,omitempty"`
	MonetaryAmounts               *[]MonetaryAmount               `json:"MonetaryAmounts,omitempty"`
	QuantityAndWeight             *[]QuantityAndWeight            `json:"QuantityAndWeight,omitempty"`
	RegulatoryCompliances         *[]RegulatoryCompliance         `json:"RegulatoryCompliances,omitempty"`
	Measurements                  *[]Measurement                  `json:"Measurements,omitempty"`
	Paperworks                    *[]Paperwork                    `json:"Paperwork,omitempty"`
	Packaging                     *[]Packaging                    `json:"Packaging,omitempty"`
	CarrierSpecialHandlingDetails *[]CarrierSpecialHandlingDetail `json:"CarrierSpecialHandlingDetail,omitempty"`
	MarksAndNumbersCollections    *[]MarksAndNumbersCollection    `json:"MarksAndNumbersCollection,omitempty"`
	RestrictionsOrConditions      *[]RestrictionOrCondition       `json:"RestrictionsOrConditions,omitempty"`
	LeadTimes                     *[]LeadTime                     `json:"LeadTime,omitempty"`
}

type LineItem struct {
	OrderLine                    *OrderLine                      `json:"OrderLine,omitempty"`
	InvoiceLine                  *InvoiceLine                    `json:"InvoiceLine,omitempty"`
	Contacts                     *[]Contact                      `json:"Contacts,omitempty"`
	Dates                        *[]Date                         `json:"Dates,omitempty"`
	Measurements                 *[]Measurement                  `json:"Measurements,omitempty"`
	PriceInformation             *[]PriceInformation             `json:"PriceInformation,omitempty"`
	ProductOrItemDescription     *[]ProductOrItemDescription     `json:"ProductOrItemDescription,omitempty"`
	Paperwork                    *[]Paperwork                    `json:"Paperwork,omitempty"`
	PhysicalDetails              *[]PhysicalDetail               `json:"PhysicalDetails,omitempty"`
	PalletInformations           *[]PalletInformation            `json:"PalletInformation,omitempty"`
	References                   *[]Reference                    `json:"References,omitempty"`
	Notes                        *[]Note                         `json:"Notes,omitempty"`
	FloorReady                   *[]FloorReady                   `json:"FloorReady,omitempty"`
	Addresses                    *[]Address                      `json:"Address,omitempty"`
	Subline                      *[]Subline                      `json:"Subline,omitempty"`
	QuantitiesSchedulesLocations *[]QuantitiesSchedulesLocations `json:"QuantitiesSchedulesLocations,omitempty"`
	Taxes                        *[]Tax                          `json:"Taxes,omitempty"`
	ChargesAllowances            *[]ChargesAllowance             `json:"ChargesAllowances,omitempty"`
	PaymentTerms                 *[]PaymentTerm                  `json:"PaymentTerms,omitempty"`
	ConditionOfSale              *[]ConditionOfSale              `json:"ConditionOfSale,omitempty"`
	FOBRelatedInstruction        *[]FOBRelatedInstruction        `json:"FOBRelatedInstruction,omitempty"`
	Commodity                    *[]Commodity                    `json:"Commodity,omitempty"`
	CarrierInformation           *[]CarrierInformation           `json:"CarrierInformation,omitempty"`
	CarrierSpecialHandlingDetail *[]CarrierSpecialHandlingDetail `json:"CarrierSpecialHandlingDetail,omitempty"`
	MarksAndNumbersCollection    *[]MarksAndNumbersCollection    `json:"MarksAndNumbersCollection,omitempty"`
	RestrictionsOrConditions     *[]RestrictionOrCondition       `json:"RestrictionsOrConditions,omitempty"`
	Packaging                    *[]Packaging                    `json:"Packaging,omitempty"`
	MonetaryAmounts              *[]MonetaryAmount               `json:"MonetaryAmounts,omitempty"`
	RegulatoryCompliances        *[]RegulatoryCompliance         `json:"RegulatoryCompliances,omitempty"`
	LineItemAcknowledgements     *[]LineItemAcknowledgement      `json:"LineItemAcknowledgement,omitempty"`
	MasterItemAttribute          *struct {
		ItemAttributes []ItemAttribute `json:"ItemAttributes,omitempty"`
	} `json:"MasterItemAttribute,omitempty"`
}

type Summary struct {
	TotalAmount         *float64 `json:"TotalAmount,omitempty"`
	TotalLineItemNumber *int     `json:"TotalLineItemNumber,omitempty"`
	Description         *string  `json:"Description,omitempty"`
}

type OrderHeader struct {
	Department               *string  `json:"Department,omitempty"`
	Division                 *string  `json:"Division,omitempty"`
	PrimaryPOTypeCode        *string  `json:"PrimaryPOTypeCode,omitempty"`
	ShipCompleteCode         *string  `json:"ShipCompleteCode,omitempty"`
	BuyersCurrency           *string  `json:"BuyersCurrency,omitempty"`
	PurchaseOrderDate        *string  `json:"PurchaseOrderDate,omitempty"`
	PurchaseOrderNumber      *string  `json:"PurchaseOrderNumber,omitempty"`
	TradingPartnerId         *string  `json:"TradingPartnerId,omitempty"`
	TsetPurposeCode          *string  `json:"TsetPurposeCode,omitempty"`
	Vendor                   *string  `json:"Vendor,omitempty"`
	PrimaryPOTypeDescription *string  `json:"PrimaryPOTypeDescription,omitempty"`
	PurchaseOrderTime        *string  `json:"PurchaseOrderTime,omitempty"`
	ReleaseNumber            *string  `json:"ReleaseNumber,omitempty"`
	AcknowledgementNumber    *string  `json:"AcknowledgementNumber,omitempty"`
	InternalOrderNumber      *string  `json:"InternalOrderNumber,omitempty"`
	InternalOrderDate        *string  `json:"InternalOrderDate,omitempty"`
	AcknowledgementTime      *string  `json:"AcknowledgementTime,omitempty"`
	SellersCurrency          *string  `json:"SellersCurrency,omitempty"`
	ExchangeRate             *float64 `json:"ExchangeRate,omitempty"`
	DepartmentDescription    *string  `json:"DepartmentDescription,omitempty"`
	JobNumber                *string  `json:"JobNumber,omitempty"`
	CustomerAccountNumber    *string  `json:"CustomerAccountNumber,omitempty"`
	CustomerOrderNumber      *string  `json:"CustomerOrderNumber,omitempty"`
	DocumentVersion          *string  `json:"DocumentVersion,omitempty"`
	DocumentRevision         *string  `json:"DocumentRevision,omitempty"`

	AcknowledgementDate   *string                 `json:"AcknowledgementDate,omitempty"`
	AcknowledgementType   *string                 `json:"AcknowledgementType,omitempty"`
	AdditionalPOTypeCodes *[]AdditionalPOTypeCode `json:"AdditionalPOTypeCodes,omitempty"`
}

type AdditionalPOTypeCode struct {
	POTypeCode        *string `json:"POTypeCode,omitempty"`
	POTypeDescription *string `json:"POTypeDescription,omitempty"`
}

type Address struct {
	AddressTypeCode        *string      `json:"AddressTypeCode,omitempty"`
	LocationCodeQualifier  *string      `json:"LocationCodeQualifier,omitempty"`
	AddressLocationNumber  *string      `json:"AddressLocationNumber,omitempty"`
	AddressName            *string      `json:"AddressName,omitempty"`
	AddressAlternateName   *string      `json:"AddressAlternateName,omitempty"`
	AddressAlternateName2  *string      `json:"AddressAlternateName2,omitempty"`
	Address1               *string      `json:"Address1,omitempty"`
	Address2               *string      `json:"Address2,omitempty"`
	Address3               *string      `json:"Address3,omitempty"`
	Address4               *string      `json:"Address4,omitempty"`
	City                   *string      `json:"City,omitempty"`
	State                  *string      `json:"State,omitempty"`
	PostalCode             *string      `json:"PostalCode,omitempty"`
	Country                *string      `json:"Country,omitempty"`
	LocationID             *string      `json:"LocationID,omitempty"`
	CountrySubDivision     *string      `json:"CountrySubDivision,omitempty"`
	AddressTaxIdNumber     *string      `json:"AddressTaxIdNumber,omitempty"`
	AddressTaxExemptNumber *string      `json:"AddressTaxExemptNumber,omitempty"`
	References             *[]Reference `json:"References,omitempty"`
	Contacts               *[]Contact   `json:"Contacts,omitempty"`
	Dates                  *[]Date      `json:"Dates,omitempty"`
}

type CarrierInformation struct {
	StatusCode               *string             `json:"StatusCode,omitempty"`
	CarrierTransMethodCode   *string             `json:"CarrierTransMethodCode,omitempty"`
	CarrierAlphaCode         *string             `json:"CarrierAlphaCode,omitempty"`
	CarrierRouting           *string             `json:"CarrierRouting,omitempty"`
	EquipmentDescriptionCode *string             `json:"EquipmentDescriptionCode,omitempty"`
	CarrierEquipmentInitial  *string             `json:"CarrierEquipmentInitial,omitempty"`
	CarrierEquipmentNumber   *string             `json:"CarrierEquipmentNumber,omitempty"`
	EquipmentType            *string             `json:"EquipmentType,omitempty"`
	OwnershipCode            *string             `json:"OwnershipCode,omitempty"`
	RoutingSequenceCode      *string             `json:"RoutingSequenceCode,omitempty"`
	TransitDirectionCode     *string             `json:"TransitDirectionCode,omitempty"`
	TransitTimeQual          *string             `json:"TransitTimeQual,omitempty"`
	TransitTime              *float64            `json:"TransitTime,omitempty"`
	ServiceLevelCodes        *[]ServiceLevelCode `json:"ServiceLevelCodes,omitempty"`
	Addresses                *[]Address          `json:"Address,omitempty"`
	SealNumbers              *[]SealNumber       `json:"SealNumbers,omitempty"`
}

type Date struct {
	Date              *string `json:"Date,omitempty"`
	Time              *string `json:"Time,omitempty"`
	DateTimeQualifier *string `json:"DateTimeQualifier,omitempty"`
	DateTimePeriod    *string `json:"DateTimePeriod,omitempty"`
}

type PaymentTerm struct {
	TermsType                *string  `json:"TermsType,omitempty"`
	TermsBasisDateCode       *string  `json:"TermsBasisDateCode,omitempty"`
	TermsTimeRelationCode    *string  `json:"TermsTimeRelationCode,omitempty"`
	TermsDiscountPercentage  *float64 `json:"TermsDiscountPercentage,omitempty"`
	TermsDiscountDate        *string  `json:"TermsDiscountDate,omitempty"`
	TermsDiscountDueDays     *int     `json:"TermsDiscountDueDays,omitempty"`
	TermsNetDueDate          *string  `json:"TermsNetDueDate,omitempty"`
	TermsNetDueDays          *int     `json:"TermsNetDueDays,omitempty"`
	TermsDiscountAmount      *float64 `json:"TermsDiscountAmount,omitempty"`
	TermsDeferredDueDate     *string  `json:"TermsDeferredDueDate,omitempty"`
	TermsDeferredAmountDue   *float64 `json:"TermsDeferredAmountDue,omitempty"`
	PercentOfInvoicePayable  *float64 `json:"PercentOfInvoicePayable,omitempty"`
	TermsDescription         *string  `json:"TermsDescription,omitempty"`
	TermsDueDay              *int     `json:"TermsDueDay,omitempty"`
	PaymentMethodCode        *string  `json:"PaymentMethodCode,omitempty"`
	PaymentMethodID          *string  `json:"PaymentMethodID,omitempty"`
	LatePaymentChargePercent *float64 `json:"LatePaymentChargePercent,omitempty"`
	TermsStartDate           *string  `json:"TermsStartDate,omitempty"`
	TermsDueDateQual         *string  `json:"TermsDueDateQual,omitempty"`
	AmountSubjectToDiscount  *float64 `json:"AmountSubjectToDiscount,omitempty"`
	DiscountAmountDue        *float64 `json:"DiscountAmountDue,omitempty"`
}

type Reference struct {
	ReferenceQual *string      `json:"ReferenceQual,omitempty"`
	ReferenceID   *string      `json:"ReferenceID,omitempty"`
	Description   *string      `json:"Description,omitempty"`
	Date          *string      `json:"Date,omitempty"`
	Time          *string      `json:"Time,omitempty"`
	ReferenceIDs  *[]Reference `json:"ReferenceIDs,omitempty"`
}

type ChargesAllowance struct {
	AllowChrgIndicator           *string  `json:"AllowChrgIndicator,omitempty"`
	AllowChrgCode                *string  `json:"AllowChrgCode,omitempty"`
	AllowChrgAgencyCode          *string  `json:"AllowChrgAgencyCode,omitempty"`
	AllowChrgAgency              *string  `json:"AllowChrgAgency,omitempty"`
	AllowChrgAmt                 *float64 `json:"AllowChrgAmt,omitempty"`
	AllowChrgPercentQual         *string  `json:"AllowChrgPercentQual,omitempty"`
	AllowChrgPercent             *float64 `json:"AllowChrgPercent,omitempty"`
	PercentDollarBasis           *float64 `json:"PercentDollarBasis,omitempty"`
	AllowChrgRate                *float64 `json:"AllowChrgRate,omitempty"`
	AllowChrgQtyUOM              *string  `json:"AllowChrgQtyUOM,omitempty"`
	AllowChrgQty                 *float64 `json:"AllowChrgQty,omitempty"`
	AllowChrgHandlingCode        *string  `json:"AllowChrgHandlingCode,omitempty"`
	ReferenceIdentification      *string  `json:"ReferenceIdentification,omitempty"`
	AllowChrgHandlingDescription *string  `json:"AllowChrgHandlingDescription,omitempty"`
	OptionNumber                 *string  `json:"OptionNumber,omitempty"`
	ExceptionNumber              *string  `json:"ExceptionNumber,omitempty"`
	AllowChrgQty2                *float64 `json:"AllowChrgQty2,omitempty"`
	LanguageCode                 *string  `json:"LanguageCode,omitempty"`
	CalculationSequence          *int     `json:"CalculationSequence,omitempty"`
	Taxes                        *[]Tax   `json:"Taxes,omitempty"`
}

type Contact struct {
	ContactName     *string `json:"ContactName,omitempty"`
	ContactTypeCode *string `json:"ContactTypeCode,omitempty"`
	PrimaryEmail    *string `json:"PrimaryEmail,omitempty"`
	PrimaryFax      *string `json:"PrimaryFax,omitempty"`
	PrimaryPhone    *string `json:"PrimaryPhone,omitempty"`
}

type FOBRelatedInstruction struct {
	FOBPayCode              *string `json:"FOBPayCode,omitempty"`
	FOBLocationQualifier    *string `json:"FOBLocationQualifier,omitempty"`
	FOBLocationDescription  *string `json:"FOBLocationDescription,omitempty"`
	FOBTitlePassageCode     *string `json:"FOBTitlePassageCode,omitempty"`
	FOBTitlePassageLocation *string `json:"FOBTitlePassageLocation,omitempty"`
	TransportationTermsType *string `json:"TransportationTermsType,omitempty"`
	TransportationTerms     *string `json:"TransportationTerms,omitempty"`
	RiskOfLossCode          *string `json:"RiskOfLossCode,omitempty"`
	Description             *string `json:"Description,omitempty"`
}

type Note struct {
	Note         *string `json:"Note,omitempty"`
	NoteCode     *string `json:"NoteCode,omitempty"`
	LanguageCode *string `json:"LanguageCode,omitempty"`
}

type QuantityTotal struct {
	QuantityTotalsQualifier *string  `json:"QuantityTotalsQualifier,omitempty"`
	Quantity                *float64 `json:"Quantity,omitempty"`
	QuantityUOM             *string  `json:"QuantityUOM,omitempty"`
	WeightQualifier         *string  `json:"WeightQualifier,omitempty"`
	Weight                  *float64 `json:"Weight,omitempty"`
	WeightUOM               *string  `json:"WeightUOM,omitempty"`
	Volume                  *float64 `json:"Volume,omitempty"`
	VolumeUOM               *string  `json:"VolumeUOM,omitempty"`
	Description             *string  `json:"Description,omitempty"`
}

type Commodity struct {
	CommodityCodeQualifier *string `json:"CommodityCodeQualifier,omitempty"`
	CommodityCode          *string `json:"CommodityCode,omitempty"`
}

type Tax struct {
	TaxTypeCode        *string  `json:"TaxTypeCode,omitempty"`
	TaxAmount          *float64 `json:"TaxAmount,omitempty"`
	TaxPercentQual     *string  `json:"TaxPercentQual,omitempty"`
	TaxPercent         *float64 `json:"TaxPercent,omitempty"`
	JurisdictionQual   *string  `json:"JurisdictionQual,omitempty"`
	JurisdictionCode   *string  `json:"JurisdictionCode,omitempty"`
	TaxExemptCode      *string  `json:"TaxExemptCode,omitempty"`
	RelationshipCode   *string  `json:"RelationshipCode,omitempty"`
	PercentDollarBasis *float64 `json:"PercentDollarBasis,omitempty"`
	TaxHandlingCode    *string  `json:"TaxHandlingCode,omitempty"`
	TaxID              *string  `json:"TaxID,omitempty"`
	AssignedID         *string  `json:"AssignedID,omitempty"`
	Description        *string  `json:"Description,omitempty"`
}

type MonetaryAmount struct {
	MonetaryAmountCode *string  `json:"MonetaryAmountCode,omitempty"`
	MonetaryAmount     *float64 `json:"MonetaryAmount,omitempty"`
	CreditDebitFlag    *string  `json:"CreditDebitFlag,omitempty"`
}

type QuantityAndWeight struct {
	PackingMedium      *string  `json:"PackingMedium,omitempty"`
	PackingMaterial    *string  `json:"PackingMaterial,omitempty"`
	LadingQuantity     *int     `json:"LadingQuantity,omitempty"`
	LadingDescription  *string  `json:"LadingDescription,omitempty"`
	WeightQualifier    *string  `json:"WeightQualifier,omitempty"`
	Weight             *float64 `json:"Weight,omitempty"`
	WeightUOM          *string  `json:"WeightUOM,omitempty"`
	Volume             *float64 `json:"Volume,omitempty"`
	VolumeUOM          *string  `json:"VolumeUOM,omitempty"`
	PalletExchangeCode *string  `json:"PalletExchangeCode,omitempty"`
}

type RegulatoryCompliance struct {
	RegulatoryComplianceQual *string `json:"RegulatoryComplianceQual,omitempty"`
	YesOrNoResponse          *string `json:"YesOrNoResponse,omitempty"`
	RegulatoryComplianceID   *string `json:"RegulatoryComplianceID,omitempty"`
	RegulatoryAgency         *string `json:"RegulatoryAgency,omitempty"`
	Description              *string `json:"Description,omitempty"`
}

type Measurement struct {
	MeasurementRefIDCode        *string  `json:"MeasurementRefIDCode,omitempty"`
	MeasurementQualifier        *string  `json:"MeasurementQualifier,omitempty"`
	MeasurementValue            *float64 `json:"MeasurementValue,omitempty"`
	CompositeUOM                *string  `json:"CompositeUOM,omitempty"`
	RangeMinimum                *float64 `json:"RangeMinimum,omitempty"`
	RangeMaximum                *float64 `json:"RangeMaximum,omitempty"`
	MeasurementSignificanceCode *string  `json:"MeasurementSignificanceCode,omitempty"`
	MeasurementAttributeCode    *string  `json:"MeasurementAttributeCode,omitempty"`
	SurfaceLayerPositionCode    *string  `json:"SurfaceLayerPositionCode,omitempty"`
	IndustryCodeQualifier       *string  `json:"IndustryCodeQualifier,omitempty"`
	IndustryCode                *string  `json:"IndustryCode,omitempty"`
}

type Paperwork struct {
	ReportTypeCode         *string `json:"ReportTypeCode,omitempty"`
	ReportTransmissionCode *string `json:"ReportTransmissionCode,omitempty"`
	ReportCopiesNeeded     *int    `json:"ReportCopiesNeeded,omitempty"`
	AddressTypeCode        *string `json:"AddressTypeCode,omitempty"`
	LocationCodeQualifier  *string `json:"LocationCodeQualifier,omitempty"`
	LocationNumber         *string `json:"LocationNumber,omitempty"`
	Description            *string `json:"Description,omitempty"`
	ActionsIndicated       *string `json:"ActionsIndicated,omitempty"`
	RequestCategoryCode    *string `json:"RequestCategoryCode,omitempty"`
}

type Packaging struct {
	PackagingCharacteristicCode *string        `json:"PackagingCharacteristicCode,omitempty"`
	AgencyQualifierCode         *string        `json:"AgencyQualifierCode,omitempty"`
	PackagingDescriptionCode    *string        `json:"PackagingDescriptionCode,omitempty"`
	PackagingDescription        *string        `json:"PackagingDescription,omitempty"`
	UnitLoadOptionCode          *string        `json:"UnitLoadOptionCode,omitempty"`
	Measurements                *[]Measurement `json:"Measurements,omitempty"`
}

type CarrierSpecialHandlingDetail struct {
	SpecialHandlingCode    *string `json:"SpecialHandlingCode,omitempty"`
	HazardousMaterialCode  *string `json:"HazardousMaterialCode,omitempty"`
	HazardousMaterialClass *string `json:"HazardousMaterialClass,omitempty"`
	Description            *string `json:"Description,omitempty"`
	YesOrNoResponse        *string `json:"YesOrNoResponse,omitempty"`
}

type MarksAndNumbersCollection struct {
	MarksAndNumbersQualifier1 *string `json:"MarksAndNumbersQualifier1,omitempty"`
	MarksAndNumbers1          *string `json:"MarksAndNumbers1,omitempty"`
}

type RestrictionOrCondition struct {
	RestrictionsConditionsQualifier *string  `json:"RestrictionsConditionsQualifier,omitempty"`
	Description                     *string  `json:"Description,omitempty"`
	QuantityQualifier               *string  `json:"QuantityQualifier,omitempty"`
	Quantity1                       *float64 `json:"Quantity1,omitempty"`
	AmountQualifier                 *string  `json:"AmountQualifier,omitempty"`
	Amount                          *float64 `json:"Amount,omitempty"`
}

type LeadTime struct {
	LeadTimeCode           *string      `json:"LeadTimeCode,omitempty"`
	LeadTimeQuantity       *float64     `json:"LeadTimeQuantity,omitempty"`
	LeadTimePeriodInterval *string      `json:"LeadTimePeriodInterval,omitempty"`
	LeadTimeDate           *string      `json:"LeadTimeDate,omitempty"`
	References             *[]Reference `json:"References,omitempty"`
	Notes                  *[]Note      `json:"Notes,omitempty"`
}

type SealNumber struct {
	SealStatusCode *string `json:"SealStatusCode,omitempty"`
	SealNumber     *string `json:"SealNumber,omitempty"`
}

type ServiceLevelCode struct {
	ServiceLevelCode *string `json:"ServiceLevelCode,omitempty"`
}

type OrderLine struct {
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
	OrderQty                        *float64                 `json:"OrderQty,omitempty"`
	OrderQtyUOM                     *string                  `json:"OrderQtyUOM,omitempty"`
	PurchasePriceType               *string                  `json:"PurchasePriceType,omitempty"`
	PurchasePrice                   *float64                 `json:"PurchasePrice,omitempty"`
	PurchasePriceBasis              *string                  `json:"PurchasePriceBasis,omitempty"`
	BuyersCurrency                  *string                  `json:"BuyersCurrency,omitempty"`
	SellersCurrency                 *string                  `json:"SellersCurrency,omitempty"`
	ExchangeRate                    *float64                 `json:"ExchangeRate,omitempty"`
	ShipDate                        *string                  `json:"ShipDate,omitempty"`
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
	Gender                          *string                  `json:"Gender,omitempty"`
	SellerDateCode                  *string                  `json:"SellerDateCode,omitempty"`
	ProductIDs                      *[]ProductID             `json:"ProductID,omitempty"`
	NRFStandardColorAndSize         *NRFStandardColorAndSize `json:"NRFStandardColorAndSize,omitempty"`
}

type ProductID struct {
	PartNumber     *string `json:"PartNumber,omitempty"`
	PartNumberQual *string `json:"PartNumberQual,omitempty"`
}

type NRFStandardColorAndSize struct {
	NRFColorCode             *string `json:"NRFColorCode,omitempty"`
	ColorCategoryName        *string `json:"ColorCategoryName,omitempty"`
	ColorPrimaryDescription  *string `json:"ColorPrimaryDescription,omitempty"`
	NRFSizeCode              *string `json:"NRFSizeCode,omitempty"`
	SizeCategoryName         *string `json:"SizeCategoryName,omitempty"`
	SizePrimaryDescription   *string `json:"SizePrimaryDescription,omitempty"`
	SizeSecondaryDescription *string `json:"SizeSecondaryDescription,omitempty"`
	SizeTableName            *string `json:"SizeTableName,omitempty"`
	SizeHeading1             *string `json:"SizeHeading1,omitempty"`
	SizeHeading2             *string `json:"SizeHeading2,omitempty"`
	SizeHeading3             *string `json:"SizeHeading3,omitempty"`
	SizeHeading4             *string `json:"SizeHeading4,omitempty"`
}

type LineItemAcknowledgement struct {
	LineSequenceNumber              *string      `json:"LineSequenceNumber,omitempty"`
	ItemScheduleDate                *string      `json:"ItemScheduleDate,omitempty"`
	ItemScheduleQty                 *float64     `json:"ItemScheduleQty,omitempty"`
	ItemScheduleQualifier           *string      `json:"ItemScheduleQualifier,omitempty"`
	ItemScheduleUOM                 *string      `json:"ItemScheduleUOM,omitempty"`
	ItemStatusCode                  *string      `json:"ItemStatusCode,omitempty"`
	BuyerPartNumber                 *string      `json:"BuyerPartNumber,omitempty"`
	VendorPartNumber                *string      `json:"VendorPartNumber,omitempty"`
	ConsumerPackageCode             *string      `json:"ConsumerPackageCode,omitempty"`
	EAN                             *string      `json:"EAN,omitempty"`
	GTIN                            *string      `json:"GTIN,omitempty"`
	UPCCaseCode                     *string      `json:"UPCCaseCode,omitempty"`
	NatlDrugCode                    *string      `json:"NatlDrugCode,omitempty"`
	InternationalStandardBookNumber *string      `json:"InternationalStandardBookNumber,omitempty"`
	PurchasePriceType               *string      `json:"PurchasePriceType,omitempty"`
	PurchasePrice                   *float64     `json:"PurchasePrice,omitempty"`
	PurchasePriceBasis              *string      `json:"PurchasePriceBasis,omitempty"`
	BuyersCurrency                  *string      `json:"BuyersCurrency,omitempty"`
	SellersCurrency                 *string      `json:"SellersCurrency,omitempty"`
	ExchangeRate                    *float64     `json:"ExchangeRate,omitempty"`
	ExtendedItemTotal               *float64     `json:"ExtendedItemTotal,omitempty"`
	ProductSizeCode                 *string      `json:"ProductSizeCode,omitempty"`
	ProductSizeDescription          *string      `json:"ProductSizeDescription,omitempty"`
	ProductColorCode                *string      `json:"ProductColorCode,omitempty"`
	ProductColorDescription         *string      `json:"ProductColorDescription,omitempty"`
	ProductMaterialCode             *string      `json:"ProductMaterialCode,omitempty"`
	ProductMaterialDescription      *string      `json:"ProductMaterialDescription,omitempty"`
	ProductProcessCode              *string      `json:"ProductProcessCode,omitempty"`
	ProductProcessDescription       *string      `json:"ProductProcessDescription,omitempty"`
	Department                      *string      `json:"Department,omitempty"`
	DepartmentDescription           *string      `json:"DepartmentDescription,omitempty"`
	Class                           *string      `json:"Class,omitempty"`
	Gender                          *string      `json:"Gender,omitempty"`
	IndustryCode                    *string      `json:"IndustryCode,omitempty"`
	ProductIDs                      *[]ProductID `json:"ProductID,omitempty"`
}

type PriceInformation struct {
	ChangeReasonCode         *string  `json:"ChangeReasonCode,omitempty"`
	EffectiveDate            *string  `json:"EffectiveDate,omitempty"`
	PriceTypeIDCode          *string  `json:"PriceTypeIDCode,omitempty"`
	UnitPrice                *float64 `json:"UnitPrice,omitempty"`
	UnitPriceBasis           *string  `json:"UnitPriceBasis,omitempty"`
	UnitPriceBasisMultiplier *float64 `json:"UnitPriceBasisMultiplier,omitempty"`
	Currency                 *string  `json:"Currency,omitempty"`
	PriceMultiplierQual      *string  `json:"PriceMultiplierQual,omitempty"`
	PriceMultiplier          *float64 `json:"PriceMultiplier,omitempty"`
	RebateAmount             *float64 `json:"RebateAmount,omitempty"`
	Quantity                 *float64 `json:"Quantity,omitempty"`
	QuantityUOM              *string  `json:"QuantityUOM,omitempty"`
	MultiplePriceQuantity    *float64 `json:"MultiplePriceQuantity,omitempty"`
	ClassOfTradeCode         *string  `json:"ClassOfTradeCode,omitempty"`
	ConditionValue           *string  `json:"ConditionValue,omitempty"`
	Description              *string  `json:"Description,omitempty"`
}

type ProductOrItemDescription struct {
	ProductCharacteristicCode *string `json:"ProductCharacteristicCode,omitempty"`
	AgencyQualifierCode       *string `json:"AgencyQualifierCode,omitempty"`
	ProductDescriptionCode    *string `json:"ProductDescriptionCode,omitempty"`
	ProductDescription        *string `json:"ProductDescription,omitempty"`
	SurfaceLayerPositionCode  *string `json:"SurfaceLayerPositionCode,omitempty"`
	SourceSubqualifier        *string `json:"SourceSubqualifier,omitempty"`
	YesOrNoResponse           *string `json:"YesOrNoResponse,omitempty"`
	LanguageCode              *string `json:"LanguageCode,omitempty"`
}

type ItemAttribute struct {
	ItemAttributeQualifier *string        `json:"ItemAttributeQualifier,omitempty"`
	Value                  *float64       `json:"Value,omitempty"`
	ValueUOM               *string        `json:"ValueUOM,omitempty"`
	Description            *string        `json:"Description,omitempty"`
	YesOrNoResponse        *string        `json:"YesOrNoResponse,omitempty"`
	Measurements           *[]Measurement `json:"Measurements,omitempty"`
}

type PhysicalDetail struct {
	PackQualifier            *string  `json:"PackQualifier,omitempty"`
	PackSize                 *float64 `json:"PackSize,omitempty"`
	PackUOM                  *string  `json:"PackUOM,omitempty"`
	PackValue                *int     `json:"PackValue,omitempty"`
	PackingMedium            *string  `json:"PackingMedium,omitempty"`
	PackingMaterial          *string  `json:"PackingMaterial,omitempty"`
	WeightQualifier          *string  `json:"WeightQualifier,omitempty"`
	PackWeight               *float64 `json:"PackWeight,omitempty"`
	PackWeightUOM            *string  `json:"PackWeightUOM,omitempty"`
	PackVolume               *float64 `json:"PackVolume,omitempty"`
	PackVolumeUOM            *string  `json:"PackVolumeUOM,omitempty"`
	PackLength               *float64 `json:"PackLength,omitempty"`
	PackWidth                *float64 `json:"PackWidth,omitempty"`
	PackHeight               *float64 `json:"PackHeight,omitempty"`
	DimensionUOM             *string  `json:"DimensionUOM,omitempty"`
	Description              *string  `json:"Description,omitempty"`
	SurfaceLayerPositionCode *string  `json:"SurfaceLayerPositionCode,omitempty"`
	AssignedID               *string  `json:"AssignedID,omitempty"`
}

type PalletInformation struct {
	PalletQualifier     *string  `json:"PalletQualifier,omitempty"`
	PalletValue         *int     `json:"PalletValue,omitempty"`
	PalletTypeCode      *string  `json:"PalletTypeCode,omitempty"`
	PalletTiers         *int     `json:"PalletTiers,omitempty"`
	PalletBlocks        *int     `json:"PalletBlocks,omitempty"`
	UnitWeight          *float64 `json:"UnitWeight,omitempty"`
	UnitWeightUOM       *string  `json:"UnitWeightUOM,omitempty"`
	Length              *float64 `json:"Length,omitempty"`
	Width               *float64 `json:"Width,omitempty"`
	Height              *float64 `json:"Height,omitempty"`
	DimensionUOM        *string  `json:"DimensionUOM,omitempty"`
	WeightQualifier     *string  `json:"WeightQualifier,omitempty"`
	PalletWeight        *float64 `json:"PalletWeight,omitempty"`
	PalletWeightUOM     *string  `json:"PalletWeightUOM,omitempty"`
	PalletVolume        *float64 `json:"PalletVolume,omitempty"`
	PalletVolumeUOM     *string  `json:"PalletVolumeUOM,omitempty"`
	UnloadedHeight      *float64 `json:"UnloadedHeight,omitempty"`
	UnloadedHeightUOM   *string  `json:"UnloadedHeightUOM,omitempty"`
	PalletExchangeCode  *string  `json:"PalletExchangeCode,omitempty"`
	PalletStructureCode *string  `json:"PalletStructureCode,omitempty"`
}

type FloorReady struct {
	FloorReadyRequired    *string `json:"FloorReadyRequired,omitempty"`
	FloorReadyTypeCode    *string `json:"FloorReadyTypeCode,omitempty"`
	FloorReadyDescription *string `json:"FloorReadyDescription,omitempty"`
	FloorReadyID          *string `json:"FloorReadyID,omitempty"`
}

type Subline struct {
	SublineItemDetail        *SublineItemDetail          `json:"SublineItemDetail,omitempty"`
	Dates                    *[]Date                     `json:"Dates,omitempty"`
	PriceInformation         *[]PriceInformation         `json:"PriceInformation,omitempty"`
	ProductOrItemDescription *[]ProductOrItemDescription `json:"ProductOrItemDescription,omitempty"`
	PhysicalDetails          *[]PhysicalDetail           `json:"PhysicalDetails,omitempty"`
	References               *[]Reference                `json:"References,omitempty"`
	Notes                    *[]Note                     `json:"Notes,omitempty"`
	FloorReady               *[]FloorReady               `json:"FloorReady,omitempty"`
	Taxes                    *[]Tax                      `json:"Taxes,omitempty"`
	ChargesAllowances        *[]ChargesAllowance         `json:"ChargesAllowances,omitempty"`
	Addresses                *[]Address                  `json:"Address,omitempty"`
	Commodity                *[]Commodity                `json:"Commodity,omitempty"`
	RegulatoryCompliances    *[]RegulatoryCompliance     `json:"RegulatoryCompliances,omitempty"`
}

type SublineItemDetail struct {
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
	ProductSizeCode                 *string                  `json:"ProductSizeCode,omitempty"`
	ProductSizeDescription          *string                  `json:"ProductSizeDescription,omitempty"`
	ProductColorCode                *string                  `json:"ProductColorCode,omitempty"`
	ProductColorDescription         *string                  `json:"ProductColorDescription,omitempty"`
	ProductMaterialCode             *string                  `json:"ProductMaterialCode,omitempty"`
	ProductMaterialDescription      *string                  `json:"ProductMaterialDescription,omitempty"`
	ProductProcessCode              *string                  `json:"ProductProcessCode,omitempty"`
	ProductProcessDescription       *string                  `json:"ProductProcessDescription,omitempty"`
	QtyPer                          *float64                 `json:"QtyPer,omitempty"`
	QtyPerUOM                       *string                  `json:"QtyPerUOM,omitempty"`
	PurchasePriceType               *string                  `json:"PurchasePriceType,omitempty"`
	PurchasePrice                   *float64                 `json:"PurchasePrice,omitempty"`
	PurchasePriceBasis              *string                  `json:"PurchasePriceBasis,omitempty"`
	Gender                          *string                  `json:"Gender,omitempty"`
	NRFStandardColorAndSize         *NRFStandardColorAndSize `json:"NRFStandardColorAndSize,omitempty"`
	ProductID                       *[]ProductID             `json:"ProductID,omitempty"`
}

type QuantitiesSchedulesLocations struct {
	QuantityQualifier      *string             `json:"QuantityQualifier,omitempty"`
	TotalQty               *float64            `json:"TotalQty,omitempty"`
	TotalQtyUOM            *string             `json:"TotalQtyUOM,omitempty"`
	QuantityDescription    *string             `json:"QuantityDescription,omitempty"`
	LocationCodeQualifier  *string             `json:"LocationCodeQualifier,omitempty"`
	LocationDescription    *string             `json:"LocationDescription,omitempty"`
	AssignedID             *string             `json:"AssignedID,omitempty"`
	LeadTimeCode           *string             `json:"LeadTimeCode,omitempty"`
	LeadTimeQuantity       *float64            `json:"LeadTimeQuantity,omitempty"`
	LeadTimePeriodInterval *string             `json:"LeadTimePeriodInterval,omitempty"`
	LeadTimeDate           *string             `json:"LeadTimeDate,omitempty"`
	LocationQuantity       *[]LocationQuantity `json:"LocationQuantity,omitempty"`
	Dates                  *[]Date             `json:"Dates,omitempty"`
}

type ConditionOfSale struct {
	LineSequenceNumber              *string      `json:"LineSequenceNumber,omitempty"`
	SalesRequirementCode            *string      `json:"SalesRequirementCode,omitempty"`
	ActionCode                      *string      `json:"ActionCode,omitempty"`
	Amount                          *float64     `json:"Amount,omitempty"`
	AccountNumber                   *string      `json:"AccountNumber,omitempty"`
	Date                            *string      `json:"Date,omitempty"`
	AgencyQualifierCode             *string      `json:"AgencyQualifierCode,omitempty"`
	SubstitutionCode                *string      `json:"SubstitutionCode,omitempty"`
	ApplicationId                   *string      `json:"ApplicationId,omitempty"`
	BuyerPartNumber                 *string      `json:"BuyerPartNumber,omitempty"`
	VendorPartNumber                *string      `json:"VendorPartNumber,omitempty"`
	ConsumerPackageCode             *string      `json:"ConsumerPackageCode,omitempty"`
	EAN                             *string      `json:"EAN,omitempty"`
	GTIN                            *string      `json:"GTIN,omitempty"`
	UPCCaseCode                     *string      `json:"UPCCaseCode,omitempty"`
	NatlDrugCode                    *string      `json:"NatlDrugCode,omitempty"`
	InternationalStandardBookNumber *string      `json:"InternationalStandardBookNumber,omitempty"`
	ProductID                       *[]ProductID `json:"ProductID,omitempty"`
}

type LocationQuantity struct {
	Location *string  `json:"Location,omitempty"`
	Qty      *float64 `json:"Qty,omitempty"`
}
