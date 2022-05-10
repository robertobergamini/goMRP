package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"database/sql"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

var connectionstring = "server=192.168.173.147;database=RIGONIDIASIAGO;user id=sa;password=ordinifb"

func main() {
	initLog()
	InfoLogger.Print("log 1")
	InfoLogger.Print("log fatal")

	readJSONFile("D:\\Lavori\\BearProjects\\github\\robertobergamini\\gomrp\\mrp\\orders.json")
	fmt.Println("Hello, world.")
}

func initLog() {

	file, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.LUTC)
}

func readJSONFile(fileName string) (retval bool) {

	retval = false

	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// var orders []ProductionOrder
	var orders1 ProductionOrders
	err = json.Unmarshal(data, &orders1)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	fmt.Println(orders1)

	return true
}

func readDataFromDB() {
	sql.Open()
}
