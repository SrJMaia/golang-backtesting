package data

import (
  "encoding/csv"
  "fmt"
  "os"
  "strconv"
)

type dataStruct struct {
  open []float64
  high []float64
  low []float64
  close []float64
  pricetf []float64
  tpsl []float64
  sellFlag []bool
  buyFlag []bool
}

func ReadData() dataStruct {

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

  emp := dataStruct{
      open: openSlice,
      high: highSlice,
      low: lowSlice,
      close: closeSlice,
      pricetf: pricetfSlice,
      tpsl: tpslSlice,
      sellFlag: sellFlagSlice,
      buyFlag: buyFlagSlice,
  }

  return emp
}
// 
// func SaveData(data struct) {
//   var file, err = os.Create("C:/Users/johnk/Google Drive/Programming/Dados/results.csv")
//   if err != nil {
//     fmt.Println(err)
//     return
//   }
//   defer file.Close()
//
//   var writer = csv.NewWriter(file)
//   defer writer.Flush()
//
//   for _, value := range data {
//       var err = writer.Write(value)
//       if err != nil {
//         fmt.Println(err)
//         break
//       }
//   }
//   fmt.Println("Data Saved.")
// }
