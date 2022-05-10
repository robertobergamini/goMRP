package main

// order
type IOrderScheduer interface {
	ReadFromJSONFile(fileName string) map[string]ProductionOrder
}

// order
type OrderScheduer struct {
	a string
}

func (o OrderScheduer) ReadFromJSONFile(fileName string) {

}
