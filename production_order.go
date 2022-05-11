package main

import (
	"time"
)

// order
type ProductionOrder struct {
	Number                string    `json:"Number"`
	OrderDate             time.Time `json:"OrderDate"`
	DueDate               time.Time `json:"DueDate"`
	ItemCodeFiniteProduct string    `json:"ItemCodeFiniteProduct"`
	ItemCodeAnonymous     string    `json:"ItemCodeAnonymous"`
	ItemCodeBatch         string    `json:"ItemCodeBatch"`
	FormatoCategory       string    `json:"FormatoCategory"`
	ColoreCategory        string    `json:"ColoreCategory"`
	WorkcenterCode        string    `json:"WorkcenterCode"`
	Quantity              float32   `json:"Quantity"`
	QuantityPerTime       float32   `json:"QuantityPerTime"`
}

type ProductionOrders struct {
	Orders []ProductionOrder
}

type ColorCategory struct {
	ColorCode        string
	ColorDescription string
}
