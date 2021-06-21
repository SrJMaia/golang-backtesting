package main
/*
Ler o csv
Salvar o csv
Backtest
Analise dos resultados
*/
import (
  "fmt"
  "github.com/SrJMaia/EA/data"
)

func main() {
  var emp = data.ReadData()

  xii := emp.open[:10]
  fmt.Println(xii)
  fmt.Println(len(emp.open))

  //data.SaveData(emp)
}
