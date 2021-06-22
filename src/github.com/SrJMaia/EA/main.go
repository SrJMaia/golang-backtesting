package main

/*
Ler o csv
Salvar o csv
Backtest
Analise dos resultados
*/
import (
	"fmt"

	"github.com/SrJMaia/EA/backtest"
	"github.com/SrJMaia/EA/data"
)

func main() {
	var data = data.ReadData()

	xii := data[0][:10]
	fmt.Println(xii)
	fmt.Println(len(data[0]))

	var buy, sell = backtest.SumArrayToBacktest(data)

	fmt.Println("Buy:", buy, "Sell:", sell)

	// data.SaveData(data)
}
