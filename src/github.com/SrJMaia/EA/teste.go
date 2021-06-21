package main

import (
  "encoding/csv"
  "fmt"
  "os"
  "strconv"
)

type csvData struct {
  open []float64
  high []float64
  low []float64
  close []float64
  pricetf []float64
  tpsl []float64
  sellFlag []float64
  buyFlag []float64
}

func main() {
  csvFile, err := os.Open("C:\\Users\\johnk\\Google Drive\\Programming\\Dados\\EURUSD_20y_OHLC.csv")
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println("Successfully Opened CSV file")
  defer csvFile.Close()

  csvLines, err := csv.NewReader(csvFile).ReadAll()
  if err != nil {
      fmt.Println(err)
  }

  var o, h, l, c = []float64{}, []float64{}, []float64{}, []float64{}
  var oi float64
  var hi float64
  var li float64
  var ci float64
  for _, v := range csvLines {
    oi, err = strconv.ParseFloat(v[0], 64)
    if err != nil {
      fmt.Println(err)
      return
    }
    o = append(o, oi)
    hi, err = strconv.ParseFloat(v[1], 64)
    if err != nil {
      fmt.Println(err)
      return
    }
    h = append(h, hi)
    li, err = strconv.ParseFloat(v[2], 64)
    if err != nil {
      fmt.Println(err)
      return
    }
    l = append(l, li)
    ci, err = strconv.ParseFloat(v[3], 64)
    if err != nil {
      fmt.Println(err)
      return
    }
    c = append(c, ci)
  }

  emp := csvData{
      open: o,
      high: h,
      low: l,
      close: c,
  }

  xii := emp.open[:10]
  fmt.Println(xii)
  fmt.Println(len(emp.open))
}
