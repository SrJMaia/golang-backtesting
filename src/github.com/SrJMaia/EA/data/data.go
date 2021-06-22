package data

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/SrJMaia/EA/check"
	"github.com/SrJMaia/EA/conversion"
)

const ArrayLen = 7747555

type DataBacktest [8][ArrayLen]float64

/*
[0] Open
[1] High
[2] Low
[3] Close
[4] Preços do TimeFrame
[5] TPSL
[6] Sell Flag
[7] Buy Flag
*/

func ReadData() DataBacktest {

	var myData DataBacktest

	var openValue float64
	var highValue float64
	var lowValue float64
	var closeValue float64
	var pricetfValue float64
	var tpslValue float64
	var sellFlagValue bool
	var buyFlagValue bool

	csvFile, err := os.Open("C:/Users/johnk/Google Drive/Programming/Dados/testdata.csv")
	check.CheckError(err)
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	check.CheckError(err)

	for i, v := range csvLines {

		// Could I use fmt.Sprintf(".5f") to get the precision?

		openValue, err = strconv.ParseFloat(v[0], 64)
		check.CheckError(err)
		myData[0][i] = openValue

		highValue, err = strconv.ParseFloat(v[1], 64)
		check.CheckError(err)
		myData[1][i] = highValue

		lowValue, err = strconv.ParseFloat(v[2], 64)
		check.CheckError(err)
		myData[2][i] = lowValue

		closeValue, err = strconv.ParseFloat(v[3], 64)
		check.CheckError(err)
		myData[3][i] = closeValue

		pricetfValue, err = strconv.ParseFloat(v[4], 64)
		check.CheckError(err)
		myData[4][i] = pricetfValue

		tpslValue, err = strconv.ParseFloat(v[5], 64)
		check.CheckError(err)
		myData[5][i] = tpslValue

		sellFlagValue, err = strconv.ParseBool(v[6])
		check.CheckError(err)
		myData[6][i] = conversion.BoolToFloat(sellFlagValue)

		buyFlagValue, err = strconv.ParseBool(v[7])
		check.CheckError(err)
		myData[7][i] = conversion.BoolToFloat(buyFlagValue)

	}

	return myData
}

func SaveData(data DataBacktest) {

	var openValue []string
	var highValue []string
	var lowValue []string
	var closeValue []string
	var pricetfValue []string
	var tpslValue []string
	var sellFlagValue []string
	var buyFlagValue []string

	var file, err = os.Create("C:/Users/johnk/Google Drive/Programming/Dados/results.csv")
	check.CheckError(err)
	defer file.Close()

	var writer = csv.NewWriter(file)
	defer writer.Flush()

	for i := range data[0] {
		openValue = append(openValue, strconv.FormatFloat(data[0][i], 'E', -1, 64))
		highValue = append(highValue, strconv.FormatFloat(data[1][i], 'E', -1, 64))
		lowValue = append(lowValue, strconv.FormatFloat(data[2][i], 'E', -1, 64))
		closeValue = append(closeValue, strconv.FormatFloat(data[3][i], 'E', -1, 64))
		pricetfValue = append(pricetfValue, strconv.FormatFloat(data[4][i], 'E', -1, 64))
		tpslValue = append(tpslValue, strconv.FormatFloat(data[5][i], 'E', -1, 64))
		sellFlagValue = append(sellFlagValue, strconv.FormatBool(conversion.FloatToBool(data[6][i])))
		buyFlagValue = append(buyFlagValue, strconv.FormatBool(conversion.FloatToBool(data[7][i])))
	}

	err = writer.Write(openValue)
	check.CheckError(err)

	err = writer.Write(highValue)
	check.CheckError(err)

	err = writer.Write(lowValue)
	check.CheckError(err)

	err = writer.Write(closeValue)
	check.CheckError(err)

	err = writer.Write(pricetfValue)
	check.CheckError(err)

	err = writer.Write(tpslValue)
	check.CheckError(err)

	err = writer.Write(sellFlagValue)
	check.CheckError(err)

	err = writer.Write(buyFlagValue)
	check.CheckError(err)

	fmt.Println("Data Saved.")
}
