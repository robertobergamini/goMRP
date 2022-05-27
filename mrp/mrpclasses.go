package mrp

type ProductionStrategy struct {
	Code        string
	Descriprion string
	Produce     bool
	Buy         bool
}

type ProductType struct {
	Code           string
	Descriprion    string
	IsRawMaterial  bool
	IsSemiFinisced bool
	IsEndProduct   bool
}

type Item struct {
	Code                   string
	Descriprion            string
	UM                     string
	MinimumStock           float32
	ProductType            string
	RecipeCode             string
	CycleCode              string
	PreferredSuipplierCode string
	StockQuatity           float32
}
