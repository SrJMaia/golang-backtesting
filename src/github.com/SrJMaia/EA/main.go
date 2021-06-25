package main

import (
	"fmt"
	"time"

	"github.com/SrJMaia/EA/backtest"
	"github.com/SrJMaia/EA/data"
)

const jpy bool = false

func main() {

	var tot []float64
	var buy []float64
	var sell []float64
	start := time.Now()
	var dt = data.ReadData(jpy)
	elapsed := time.Since(start)
	fmt.Println("Time to read:", elapsed)

	start = time.Now()
	var buySum, sellSum = backtest.SumArrayToBacktest(dt)
	elapsed = time.Since(start)
	fmt.Println("Time to sum:", elapsed)

	start = time.Now()
	tot, buy, sell = backtest.HedgingBacktest(&dt, 1., 0.5, 1000., buySum, sellSum, jpy)
	//tot, buy, sell = backtest.HedgingBacktest(dt, 1., 0.5, 1000., buySum, sellSum, jpy)
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
	fmt.Println("Total Trades: ", len(tot))

	fmt.Println()

	fmt.Println(len(tot))
	fmt.Println(len(buy))
	fmt.Println(len(sell))

	var first bool = true
	for i := 0.5; i < 1.5; i += 0.5 {
		tot, buy, sell = backtest.HedgingBacktest(&dt, i, 0.5, 1000., buySum, sellSum, jpy)
		data.SaveBacktest(tot, buy, sell, i, first)
		first = false
	}

	//data.SaveData(dt)
}
