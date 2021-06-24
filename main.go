package main

/*
Ler o csv
Salvar o csv
Backtest
Analise dos resultados
*/
import (
	"fmt"

	"github.com/Go/data"
)

func main() {
	var emp = data.ReadData()

	xii := emp.Open[:10]
	fmt.Println(xii)
	fmt.Println(len(emp.Open))

	//data.SaveData(emp)
}
