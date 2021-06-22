package data

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type DataStruct struct {
	Open     []float64
	High     []float64
	Low      []float64
	Close    []float64
	Pricetf  []float64
	Tpsl     []float64
	SellFlag []bool
	BuyFlag  []bool
}

func ReadData() DataStruct {

	var openSlice, highSlice, lowSlice, closeSlice = []float64{}, []float64{}, []float64{}, []float64{}
	var pricetfSlice, tpslSlice = []float64{}, []float64{}
	var sellFlagSlice, buyFlagSlice = []bool{}, []bool{}

	var openValue float64
	var highValue float64
	var lowValue float64
	var closeValue float64
	var pricetfValue float64
	var tpslValue float64
	var sellFlagValue bool
	var buyFlagValue bool

	csvFile, err := os.Open("C:/Users/johnk/Google Drive/Programming/Dados/testdata.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range csvLines {

		// Could I use fmt.Sprintf(".5f") to get the precision?

		openValue, err = strconv.ParseFloat(v[0], 64)
		if err != nil {
			fmt.Println(err)
			break
		}
		openSlice = append(openSlice, openValue)

		highValue, err = strconv.ParseFloat(v[1], 64)
		if err != nil {
			fmt.Println(err)
			break
		}
		highSlice = append(highSlice, highValue)

		lowValue, err = strconv.ParseFloat(v[2], 64)
		if err != nil {
			fmt.Println(err)
			break
		}
		lowSlice = append(lowSlice, lowValue)

		closeValue, err = strconv.ParseFloat(v[3], 64)
		if err != nil {
			fmt.Println(err)
			break
		}
		closeSlice = append(closeSlice, closeValue)

		pricetfValue, err = strconv.ParseFloat(v[4], 64)
		if err != nil {
			fmt.Println(err)
			break
		}
		pricetfSlice = append(pricetfSlice, pricetfValue)

		tpslValue, err = strconv.ParseFloat(v[5], 64)
		if err != nil {
			fmt.Println(err)
			break
		}
		tpslSlice = append(tpslSlice, tpslValue)

		sellFlagValue, err = strconv.ParseBool(v[6])
		if err != nil {
			fmt.Println(err)
			break
		}
		sellFlagSlice = append(sellFlagSlice, sellFlagValue)

		buyFlagValue, err = strconv.ParseBool(v[7])
		if err != nil {
			fmt.Println(err)
			break
		}
		buyFlagSlice = append(buyFlagSlice, buyFlagValue)

	}

	var emp = DataStruct{
		Open:     openSlice,
		High:     highSlice,
		Low:      lowSlice,
		Close:    closeSlice,
		Pricetf:  pricetfSlice,
		Tpsl:     tpslSlice,
		SellFlag: sellFlagSlice,
		BuyFlag:  buyFlagSlice,
	}

	return emp
}

func SaveData(data DataStruct) {

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
