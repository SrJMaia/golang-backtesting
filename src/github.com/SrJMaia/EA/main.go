package main

/*
Ler o csv
Salvar o csv
Backtest
Analise dos resultados
*/
import (
	"fmt"
	"time"

	"github.com/SrJMaia/EA/backtest"
	"github.com/SrJMaia/EA/data"
)

const JPY bool = false

func main() {

	var tot []float64
	var buy []float64
	var sell []float64
	start := time.Now()
	var dt = data.ReadData(JPY)
	elapsed := time.Since(start)
	fmt.Println("Time to read:", elapsed)

	start = time.Now()
	var buySum, sellSum = backtest.SumArrayToBacktest(dt)
	elapsed = time.Since(start)
	fmt.Println("Time to sum:", elapsed)

	start = time.Now()
	//tot, buy, sell = backtest.NettingBacktest(dt, 1., 0.5, 1000., buySum, sellSum, JPY)
	tot, buy, sell = backtest.HedgingBacktest(dt, 1., 0.5, 1000., buySum, sellSum, JPY)
	elapsed = time.Since(start)
	fmt.Println("Time to backtest:", elapsed)

	// fmt.Println(dt.Open[:10])
	// fmt.Println(dt.High[:10])
	// fmt.Println(dt.Low[:10])
	// fmt.Println(dt.Close[:10])
	// fmt.Println(dt.BuyFlag[:10])
	// fmt.Println(dt.SellFlag[:10])
	// fmt.Println(dt.Pricetf[:10])
	// fmt.Println(dt.Tpsl[:10])

	fmt.Println(tot[len(tot)-5:])
	fmt.Println(buy[len(buy)-5:])
	fmt.Println(sell[len(sell)-5:])

	//data.SaveData(dt)
}
