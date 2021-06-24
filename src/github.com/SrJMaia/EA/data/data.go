package data

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/SrJMaia/EA/check"
	"github.com/SrJMaia/EA/conversion"
)

type LayoutData struct {
	Open     []float64
	High     []float64
	Low      []float64
	Close    []float64
	Pricetf  []float64
	Tpsl     []float64
	SellFlag []bool
	BuyFlag  []bool
}

func ReadData(isJpy bool) LayoutData {

	var roundPlace int
	if isJpy {
		roundPlace = 3
	} else {
		roundPlace = 5
	}

	var openValue float64
	var highValue float64
	var lowValue float64
	var closeValue float64
	var pricetfValue float64
	var tpslValue float64
	var sellFlagValue bool
	var buyFlagValue bool

	csvFile, err := os.Open("C:/Users/johnk/Google Drive/Programming/Dados/testdata.csv")
	check.MyCheckingError(err)
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	check.MyCheckingError(err)

	var openSlice, highSlice = make([]float64, len(csvLines)), make([]float64, len(csvLines))
	var lowSlice, closeSlice = make([]float64, len(csvLines)), make([]float64, len(csvLines))
	var pricetfSlice, tpslSlice = make([]float64, len(csvLines)), make([]float64, len(csvLines))
	var sellFlagSlice, buyFlagSlice = make([]bool, len(csvLines)), make([]bool, len(csvLines))

	for i, v := range csvLines {

		// Could I use fmt.Sprintf(".5f") to get the precision?
		// Change everything from append to attribution

		openValue, err = strconv.ParseFloat(v[0], 64)
		check.MyCheckingError(err)
		openSlice[i] = openValue

		highValue, err = strconv.ParseFloat(v[1], 64)
		check.MyCheckingError(err)
		highSlice[i] = highValue

		lowValue, err = strconv.ParseFloat(v[2], 64)
		check.MyCheckingError(err)
		lowSlice[i] = lowValue

		closeValue, err = strconv.ParseFloat(v[3], 64)
		check.MyCheckingError(err)
		closeSlice[i] = closeValue

		pricetfValue, err = strconv.ParseFloat(v[4], 64)
		check.MyCheckingError(err)
		pricetfSlice[i] = pricetfValue

		tpslValue, err = strconv.ParseFloat(v[5], 64)
		check.MyCheckingError(err)
		tpslSlice[i] = conversion.Round(tpslValue, roundPlace)

		sellFlagValue, err = strconv.ParseBool(v[6])
		check.MyCheckingError(err)
		sellFlagSlice[i] = sellFlagValue

		buyFlagValue, err = strconv.ParseBool(v[7])
		check.MyCheckingError(err)
		buyFlagSlice[i] = buyFlagValue

	}

	var myData = LayoutData{
		Open:     openSlice,
		High:     highSlice,
		Low:      lowSlice,
		Close:    closeSlice,
		Pricetf:  pricetfSlice,
		Tpsl:     tpslSlice,
		SellFlag: sellFlagSlice,
		BuyFlag:  buyFlagSlice,
	}

	fmt.Println("Successfully Opened CSV file")

	return myData
}

func SaveData(data LayoutData) {
	/*
		There is a stackoverflow error
	*/

	var openValue []string
	var highValue []string
	var lowValue []string
	var closeValue []string
	var pricetfValue []string
	var tpslValue []string
	var sellFlagValue []string
	var buyFlagValue []string

	var file, err = os.Create("C:/Users/johnk/Google Drive/Programming/Dados/results.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var writer = csv.NewWriter(file)
	defer writer.Flush()

	for i := range data.Open {
		openValue = append(openValue, strconv.FormatFloat(data.Open[i], 'E', -1, 64))
		highValue = append(highValue, strconv.FormatFloat(data.High[i], 'E', -1, 64))
		lowValue = append(lowValue, strconv.FormatFloat(data.Low[i], 'E', -1, 64))
		closeValue = append(closeValue, strconv.FormatFloat(data.Close[i], 'E', -1, 64))
		pricetfValue = append(pricetfValue, strconv.FormatFloat(data.Pricetf[i], 'E', -1, 64))
		tpslValue = append(tpslValue, strconv.FormatFloat(data.Tpsl[i], 'E', -1, 64))
		sellFlagValue = append(sellFlagValue, strconv.FormatBool(data.SellFlag[i]))
		buyFlagValue = append(buyFlagValue, strconv.FormatBool(data.BuyFlag[i]))
	}

	err = writer.Write(openValue)
	if err != nil {
		fmt.Println(err)
	}

	err = writer.Write(highValue)
	if err != nil {
		fmt.Println(err)
	}

	err = writer.Write(lowValue)
	if err != nil {
		fmt.Println(err)
	}

	err = writer.Write(closeValue)
	if err != nil {
		fmt.Println(err)
	}

	err = writer.Write(pricetfValue)
	if err != nil {
		fmt.Println(err)
	}

	err = writer.Write(tpslValue)
	if err != nil {
		fmt.Println(err)
	}

	err = writer.Write(sellFlagValue)
	if err != nil {
		fmt.Println(err)
	}

	err = writer.Write(buyFlagValue)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Data Saved.")
}

func SaveBacktest(tot []float64, buy []float64, sell []float64, columns float64, first bool) {

	var totValue = make([]string, len(tot)+1)
	var buyValue = make([]string, len(buy)+1)
	var sellValue = make([]string, len(sell)+1)
	var index = make([]string, len(tot)*2)

	var file, err = os.OpenFile("C:/Users/johnk/Google Drive/Programming/Dados/BacktestResults.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check.MyCheckingError(err)
	defer file.Close()

	var writer = csv.NewWriter(file)
	defer writer.Flush()

	var s string = fmt.Sprintf("%f", columns)

	totValue[0] = s
	buyValue[0] = s
	sellValue[0] = s
	index[0] = "0"

	for i := 1; i < len(tot); i++ {
		totValue[i] = strconv.FormatFloat(tot[i], 'E', -1, 64)
	}
	for i := 1; i < len(buy); i++ {
		buyValue[i] = strconv.FormatFloat(buy[i], 'E', -1, 64)
	}
	for i := 1; i < len(sell); i++ {
		sellValue[i] = strconv.FormatFloat(sell[i], 'E', -1, 64)
	}

	if first {
		for i := 1; i < len(index); i++ {
			s = fmt.Sprintf("%d", i)
			index[i] = s
		}

		err = writer.Write(index)
		check.MyCheckingError(err)
	}

	err = writer.Write(totValue)
	check.MyCheckingError(err)

	err = writer.Write(buyValue)
	check.MyCheckingError(err)

	err = writer.Write(sellValue)
	check.MyCheckingError(err)

	fmt.Println("Backtest Saved.")
}
