package main

import (
	"time"
)

// order
type ProductionOrder struct {
	Number         string    `json:"Number"`
	OrderDate      time.Time `json:"OrderDate"`
	DueDate        time.Time `json:"DueDate"`
	ItemCode       string    `json:"ItemCode"`
	WorkcenterCode string    `json:"WorkcenterCode"`
	Quantity       float32   `json:"Quantity"`
}

type ProductionOrders struct {
	Orders []ProductionOrder
}
