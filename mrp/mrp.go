package mrp

import "time"

type MRPParametri struct {
	IncludeProductionForecast bool
	IncludeSalesForecast      bool
	IncludeStock              bool
	IncludeReservedStock      bool
	IncludeSalesOrder         bool
	IncludePurchaseOrder      bool
	IncludeProductionOrder    bool

	GeneratePurchaseOrderProposal   bool
	GenerateProductionOrderProposal bool

	RecipeCalculationFormula string
	RecyperExplosionType     string

	ElaborationDetails string
}

const (
	MRPRECORDTYPE_STOCK              = 100
	MRPRECORDTYPE_RESERVEDSTOCK      = 110
	MRPRECORDTYPE_PRODUCTIONFORECAST = 120

	MRPRECORDTYPE_SALESORDER            = 200
	MRPRECORDTYPE_PURCHASEORDER         = 300
	MRPRECORDTYPE_PURCHASEORDERPROPOSAL = 310

	MRPRECORDTYPE_PRODUCTIONORDER                      = 400
	MRPRECORDTYPE_PRODUCTIONORDERPROPOSAL              = 400
	MRPRECORDTYPE_PRODUCTIONORDER_AVAILABILITY         = 410
	MRPRECORDTYPE_PRODUCTIONORDER_COMMITMENT           = 420
	MRPRECORDTYPE_PRODUCTIONORDERPROPOSAL_AVAILABILITY = 430
	MRPRECORDTYPE_PRODUCTIONORDERPROPOSAL_COMMITMENT   = 440
)

const (
	DOCUMENTYPE_PRODUCTIONORDER         = 100
	DOCUMENTYPE_SALESORDER              = 110
	DOCUMENTYPE_PURCHASEORDER           = 120
	DOCUMENTYPE_PRODUCTIONORDERPROPOSAL = 120
	DOCUMENTYPE_PURCHASEORDERPROPOSAL   = 120
)

type MrpItem struct {
	ItemCode         string
	ItemDescription  string
	VariantCode      string
	RecordType       int16
	RowNumberRelated int64
	OrderIndex       int16
	Level            int16

	Datetime       time.Time
	VariantCodce   string
	OperationCode  string
	WorkcenterCode string
	UM             string

	RecipeCode           string
	CycleCode            string
	MinimalStockQuantity float32
	StockQuantity        float32

	DocumentType      int16
	DocumentNumber    string
	DocumentRowNumber string
	DocumentQuantity  float32

	IsPhantomItem bool

	ProgressiveQuantity float32
}

type MrpItemDetails struct {
	IdMRP            int64
	RowNumber        int64
	RecordType       int16
	RowNumberRelated int64
	OrderIndex       int16
	Level            int16

	ItemCode        string
	ItemDescription string
	ItemType        string
	Datetime        time.Time
	VariantCodce    string
	OperationCode   string
	WorkcenterCode  string
	CustomerCode    string
	UM              string
	UMRecipe        string
	RecipeCode      string

	StockMinimumQuantity float32
	StockQuantity        float32

	DocumentType      int16
	DocumentNumber    string
	DocumentRowNumber string
	DocumentQuantity  float32

	ProgressiveQuantity float32
}

type IMRP interface {
}

func (IMRP mrp) GelItemList() {

}
