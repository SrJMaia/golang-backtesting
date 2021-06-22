package backtest

import "github.com/SrJMaia/EA/data"

func SumArrayToBacktest(data data.DataBacktest) (uint32, uint32) {
	var buy uint32
	var sell uint32
	for i := range data[0] {
		if data[7][i] == 1. {
			buy += 1
		} else if data[6][i] == 1. {
			sell += 1
		}
	}

	return buy, sell
}

func SingleBacktest(data data.DataBacktest, multiplyTp float32, multiplySl float32, sizeBuy uint32, sizeSell uint32, balance uint32) {

}
