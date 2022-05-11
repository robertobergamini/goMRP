package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"time"

	_ "github.com/microsoft/go-mssqldb"
)

type Solution struct {
	Sequence string
	Makespan float32
}

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger

	_formati []ProductionSequenceTime
	_colori  []ProductionSequenceTime

	_timechangecolors  map[string]map[string]float32
	_timechangeformati map[string]map[string]float32
	_solutionrank      []Solution

	_minmakespan     float32
	_bestsolution    string
	_solutionsnumber int32
)

var colours = []string{"ST", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11"}
var coloursToOptimize = []string{"01", "03", "05"}

var formati = []string{"ST", "25G", "250G", "300G", "330G", "350G", "500G", "630750G", "800G"}
var formatiToOptimize = []string{"25G", "250G", "300G", "330G", "350G", "500G", "630750G", "800G"}

var connectionstring string = "server=192.168.173.147;database=RIGONIDIASIAGO;user id=sa;password=ordinifb"
var db *sql.DB

type KeyType string
type ValueType float32
type Node struct {
	Nodes map[KeyType]ValueType
	Value ValueType
}

func main() {
	initLog()
	InfoLogger.Print("log 1")
	InfoLogger.Print("log fatal")

	// readJSONFile("D:\\Lavori\\BearProjects\\github\\robertobergamini\\gomrp\\mrp\\orders.json")

	_solutionsnumber = 0
	_minmakespan = math.MaxFloat32
	_solutionrank = make([]Solution, 0)
	// _allsolutionrank = make(map[string]float32)
	// readDataFromDB()
	_timechangecolors = readTimeSequenceColors()

	// _timechangeformati = readTimeSequenceFormati()

	fmt.Println("Start date and time is: ", time.Now().String())
	heapPermutation(coloursToOptimize, len(coloursToOptimize), len(coloursToOptimize))
	// heapPermutation(formatiToOptimize, len(formatiToOptimize), len(formatiToOptimize))
	fmt.Println("Stop date and time is: ", time.Now().String())

	fmt.Printf("Best makespan %v\n", _minmakespan)
	fmt.Printf("Best solution  %v\n", _bestsolution)
	fmt.Printf("Total solution number %v\n", _solutionsnumber)

	for i, v := range _solutionrank {
		fmt.Printf("%v) %v %v\n", i, v.Sequence, v.Makespan)
	}

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

	db, err := sql.Open("sqlserver", connectionstring)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	ctx := context.Background()
	db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")

	/*	Lettura formati */
	rows, err := db.Query(`
		select Id, TipoCategoria as CategoryType, ValorePrecedente as ValueFrom, ValoreSuccessivo as ValueTo, TempoTotale as ChangeTime 
		from
			TempiCicloSequenze 
		where 
			TipoCategorie = 'FORMATO' `)
	if err != nil {
		return
	}
	defer rows.Close()

	_formati := make([]ProductionSequenceTime, 0, 0)
	for rows.Next() {
		var itm ProductionSequenceTime
		// var itm1 := new(ProductionSequenceTime)
		err := rows.Scan(&itm.Id, &itm.CategoryType, &itm.ValueFrom, &itm.ValueTo, &itm.ChangeTime)
		if err != nil {
			log.Fatal(err)
		}

		_formati = append(_formati, itm)

	}

	rows = nil

	fmt.Printf("%v\n", _formati)
	for i, v := range _formati {
		fmt.Printf("%d %v\n", i, v)
	}

	/*	Lettura colori */
	rows, err = db.Query(`
		select Id, TipoCategoria as CategoryType, ValorePrecedente as ValueFrom, ValoreSuccessivo as ValueTo, TempoTotale as ChangeTime 
		from TempiCicloSequenze
		where 
			TipoCategorie = 'COLORE' `)
	if err != nil {
		return
	}
	defer rows.Close()

	_colori := make([]ProductionSequenceTime, 0, 0)
	for rows.Next() {
		var itm ProductionSequenceTime
		// var itm1 := new(ProductionSequenceTime)
		err := rows.Scan(&itm.Id, &itm.CategoryType, &itm.ValueFrom, &itm.ValueTo, &itm.ChangeTime)
		if err != nil {
			log.Fatal(err)
		}

		_colori = append(_colori, itm)

	}

	fmt.Printf("%v\n", _colori)
	for i, v := range _colori {
		fmt.Printf("%d %v\n", i, v)
	}

	db.Close()
	db = nil

}

func heapPermutation(arr []string, size int, n int) {

	// calculate permutaion
	//var permutations := [][]string
	var solution string
	var makespan float32
	var tmp string

	if size == 1 {
		// fmt.Printf("%v\n", arr)

		_solutionsnumber++
		makespan = calculateMakespan(arr, _timechangecolors)
		if makespan < _minmakespan {
			_minmakespan = makespan

			solution = ""
			for _, v := range arr {
				solution += v
			}
			_bestsolution = solution

			_solutionrank = append(_solutionrank, Solution{Sequence: solution, Makespan: makespan})

		}

	}

	for i := 0; i < size; i++ {

		heapPermutation(arr, size-1, n)

		if size%2 == 0 {

			tmp = arr[i]
			arr[i] = arr[size-1]
			arr[size-1] = tmp

		} else {

			tmp = arr[0]
			arr[0] = arr[size-1]
			arr[size-1] = tmp

		}

		_solutionsnumber++
		makespan = calculateMakespan(arr, _timechangecolors)
		// for _, v := range arr {
		// 	solution += v
		// }
		// _allsolutionrank[solution] = makespan

		if makespan < _minmakespan {
			_minmakespan = makespan

			solution = ""
			for _, v := range arr {
				solution += v
			}

			_bestsolution = solution
			_solutionrank = append(_solutionrank, Solution{Sequence: solution, Makespan: makespan})

		}

	}

}

func swap(values []string, i int, j int) {

	var tmp string = values[i]
	values[i] = values[j]
	values[j] = tmp
}

func readTimeSequenceColors() map[string]map[string]float32 {

	var data map[string]map[string]float32 = make(map[string]map[string]float32, len(colours))

	for _, v1 := range colours {
		data[v1] = make(map[string]float32, len(colours))

		for _, v2 := range colours {
			data[v1][v2] = 0
		}
	}

	db, err := sql.Open("sqlserver", connectionstring)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	/*	Lettura colori */
	rows, err := db.Query(`
		select Id, TipoCategoria as CategoryType, ValorePrecedente as ValueFrom, ValoreSuccessivo as ValueTo, TempoTotale as ChangeTime 
		from TempiCicloSequenze
		where 
		TipoCategoria = 'COLORE' `)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	_colori := make([]ProductionSequenceTime, 0, 0)
	for rows.Next() {

		var itm ProductionSequenceTime
		// var itm1 := new(ProductionSequenceTime)
		err := rows.Scan(&itm.Id, &itm.CategoryType, &itm.ValueFrom, &itm.ValueTo, &itm.ChangeTime)
		if err != nil {
			log.Fatal(err)
		}

		data[itm.ValueFrom][itm.ValueTo] = itm.ChangeTime

		_colori = append(_colori, itm)

	}

	fmt.Printf("%v\n", _colori)
	for i, v := range _colori {
		fmt.Printf("%d %v\n", i, v)
	}

	db.Close()
	db = nil

	return data
}

func readTimeSequenceFormati() map[string]map[string]float32 {

	var data map[string]map[string]float32 = make(map[string]map[string]float32, len(colours))

	for _, v1 := range formati {
		data[v1] = make(map[string]float32, len(colours))

		// for _, v2 := range formati {
		// 	data[v1][v2] = 0
		// }
	}

	db, err := sql.Open("sqlserver", connectionstring)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	/*	Lettura formati */
	rows, err := db.Query(`
		select Id, TipoCategoria as CategoryType, ValorePrecedente as ValueFrom, ValoreSuccessivo as ValueTo, TempoTotale as ChangeTime 
		from TempiCicloSequenze
		where 
		TipoCategoria = 'FORMATO' order by ValorePrecedente, ValoreSuccessivo `)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	_formati := make([]ProductionSequenceTime, 0, 0)
	for rows.Next() {

		var itm ProductionSequenceTime
		// var itm1 := new(ProductionSequenceTime)
		err := rows.Scan(&itm.Id, &itm.CategoryType, &itm.ValueFrom, &itm.ValueTo, &itm.ChangeTime)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v %v\n", itm.ValueFrom, itm.ValueTo)

		data[itm.ValueFrom][itm.ValueTo] = itm.ChangeTime

		_formati = append(_formati, itm)

	}

	fmt.Printf("%v\n", _formati)
	for i, v := range _formati {
		fmt.Printf("%d %v\n", i, v)
	}

	db.Close()
	db = nil

	return data
}

func calculateMakespan(values []string, timMatrix map[string]map[string]float32) float32 {

	var makespan float32
	var vprec string
	for i, v := range values {
		if i > 0 {

			makespan += timMatrix[vprec][v]
		}
		vprec = v
	}

	return makespan
}
