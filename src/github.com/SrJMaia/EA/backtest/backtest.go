package backtest

import (
	"log"
	"math"

	"github.com/SrJMaia/EA/conversion"
	"github.com/SrJMaia/EA/data"
)

type backtestVariables struct {
	buyPrice      float64
	sellPrice     float64
	tpSell        float64
	slSell        float64
	tpBuy         float64
	slBuy         float64
	buyResult     float64
	sellResult    float64
	capital       float64
	buyFlag       bool
	sellFlag      bool
	universalFlag bool
	updateBuy     bool
	updateSell    bool
	iBuy          uint32
	iSell         uint32
	iTotalTrades  uint32
	totalTrades   []float64
	buyTrades     []float64
	sellTrades    []float64
}

func SumArrayToBacktest(data data.LayoutData) (uint32, uint32) {
	// Sum the true and false and return the sum

	var buy uint32
	var sell uint32
	for i := range data.Open {
		if data.SellFlag[i] {
			buy++
		} else if data.BuyFlag[i] {
			sell++
		}
	}

	return buy * 2, sell * 2
}

func financeCalculation(balance float64, initialPrice float64, finalPrice float64, eurPrice float64, initialBalance float64, leverage bool) (float64, float64) {
	var lot float64
	var comission float64
	var result float64
	var total float64
	if leverage && balance > 3*initialBalance {
		lot = conversion.Round((balance - initialBalance), -3)
	} else {
		lot = 1000
	}
	comission = lot / 1000 * .07

	result = conversion.Round((lot*(initialPrice-finalPrice))/eurPrice-comission, 5)
	total = conversion.Round(result+balance, 2)

	if math.IsNaN(total) || total == 0 {
		log.Panic("Finance Calculation Error. NaN or 0")
	}
	return total, result

}

func tpslCalculationFix(price float64, multiplyTpF float64, multiplySlF float64, jpyRound bool, buy bool) (float64, float64) {
	var tp float64
	var sl float64
	var roundTPSL int
	if jpyRound {
		roundTPSL = 3
	} else {
		roundTPSL = 5
	}
	if buy {
		tp = conversion.Round(price+multiplyTpF, roundTPSL)
		sl = conversion.Round(price-multiplySlF, roundTPSL)
	} else {
		tp = conversion.Round(price-multiplyTpF, roundTPSL)
		sl = conversion.Round(price+multiplySlF, roundTPSL)
	}
	return tp, sl
}

func tpslCalculationNonFix(price float64, multiplyTpF float64, multiplySlF float64, tpSlValue float64, jpyRound bool, buy bool) (float64, float64) {
	// Example: ATR
	var tp float64
	var sl float64
	var roundTPSL int
	if jpyRound {
		roundTPSL = 3
	} else {
		roundTPSL = 5
	}
	if buy {
		tp = conversion.Round(price+(multiplyTpF*tpSlValue), roundTPSL)
		sl = conversion.Round(price-(multiplySlF*tpSlValue), roundTPSL)
	} else {
		tp = conversion.Round(price-(multiplyTpF*tpSlValue), roundTPSL)
		sl = conversion.Round(price+(multiplySlF*tpSlValue), roundTPSL)
	}
	return tp, sl
}

func checkTpslExit(open float64, high float64, low float64, close float64, tpsl float64, higher bool) bool {
	if higher {
		return open >= tpsl || high >= tpsl || low >= tpsl || close >= tpsl
	}
	return open <= tpsl || high <= tpsl || low <= tpsl || close <= tpsl
}

func backtestHeart(balance float64, sizeBuy uint32, sizeSell uint32) backtestVariables {
	var backtestMain = backtestVariables{
		buyPrice:      0.,
		sellPrice:     0.,
		tpSell:        0.,
		slSell:        0.,
		tpBuy:         0.,
		slBuy:         0.,
		buyResult:     0.,
		sellResult:    0.,
		capital:       balance,
		buyFlag:       true,
		sellFlag:      true,
		universalFlag: true,
		updateBuy:     false,
		updateSell:    false,
		iBuy:          1,
		iSell:         1,
		iTotalTrades:  1,
		totalTrades:   make([]float64, sizeBuy+sizeSell),
		buyTrades:     make([]float64, sizeBuy),
		sellTrades:    make([]float64, sizeSell),
	}

	backtestMain.totalTrades[0] = balance
	backtestMain.buyTrades[0] = balance
	backtestMain.sellTrades[0] = balance

	return backtestMain
}

func HedgingBacktest(dt *data.LayoutData, multiplyTpB float64, multiplySlB float64, balance float64, sizeBuy uint32, sizeSell uint32, jpy bool) ([]float64, []float64, []float64) {
	// Do headging backtest and return total trades, buy trades, sell trades

	var backtestMain backtestVariables = backtestHeart(balance, sizeBuy, sizeSell)

	for i := 1; i < len((*dt).Open); i++ {

		if (*dt).Pricetf[i] != (*dt).Pricetf[i-1] {
			if (*dt).BuyFlag[i] && backtestMain.buyFlag {
				backtestMain.buyPrice = (*dt).Pricetf[i]
				backtestMain.tpBuy, backtestMain.slBuy = tpslCalculationNonFix((*dt).Pricetf[i], multiplyTpB, multiplySlB, (*dt).Tpsl[i], jpy, true)
				backtestMain.buyFlag = false
			}
			if (*dt).SellFlag[i] && backtestMain.sellFlag {
				backtestMain.sellPrice = (*dt).Pricetf[i]
				backtestMain.tpSell, backtestMain.slSell = tpslCalculationNonFix((*dt).Pricetf[i], multiplyTpB, multiplySlB, (*dt).Tpsl[i], jpy, false)
				backtestMain.sellFlag = false
			}
		} else if (*dt).Pricetf[i] == (*dt).Pricetf[i-1] {
			if !backtestMain.buyFlag {
				if checkTpslExit((*dt).Open[i], (*dt).High[i], (*dt).Low[i], (*dt).Close[i], backtestMain.slBuy, false) {
					backtestMain.capital, backtestMain.buyResult = financeCalculation(backtestMain.capital, backtestMain.slBuy, backtestMain.buyPrice, backtestMain.slBuy, balance, false)
					backtestMain.updateBuy = true
				} else if checkTpslExit((*dt).Open[i], (*dt).High[i], (*dt).Low[i], (*dt).Close[i], backtestMain.tpBuy, true) {
					backtestMain.capital, backtestMain.buyResult = financeCalculation(backtestMain.capital, backtestMain.tpBuy, backtestMain.buyPrice, backtestMain.tpBuy, balance, false)
					backtestMain.updateBuy = true
				}
				if backtestMain.updateBuy {
					backtestMain.totalTrades[backtestMain.iTotalTrades] = backtestMain.capital
					backtestMain.buyTrades[backtestMain.iBuy] = conversion.Round(backtestMain.buyTrades[backtestMain.iBuy-1]+backtestMain.buyResult, 2)
					backtestMain.iTotalTrades++
					backtestMain.iBuy++
					backtestMain.buyFlag = true
					backtestMain.updateBuy = false
					backtestMain.tpBuy = 0.
					backtestMain.slBuy = 0.
					backtestMain.buyPrice = 0.
				}
			}
			if !backtestMain.sellFlag {
				if checkTpslExit((*dt).Open[i], (*dt).High[i], (*dt).Low[i], (*dt).Close[i], backtestMain.slSell, true) {
					backtestMain.capital, backtestMain.sellResult = financeCalculation(backtestMain.capital, backtestMain.sellPrice, backtestMain.slSell, backtestMain.sellPrice, balance, false)
					backtestMain.updateSell = true
				} else if checkTpslExit((*dt).Open[i], (*dt).High[i], (*dt).Low[i], (*dt).Close[i], backtestMain.tpSell, false) {
					backtestMain.capital, backtestMain.sellResult = financeCalculation(backtestMain.capital, backtestMain.sellPrice, backtestMain.tpSell, backtestMain.sellPrice, balance, false)
					backtestMain.updateSell = true
				}
				if backtestMain.updateSell {
					backtestMain.totalTrades[backtestMain.iTotalTrades] = backtestMain.capital
					backtestMain.sellTrades[backtestMain.iSell] = conversion.Round(backtestMain.sellTrades[backtestMain.iSell-1]+backtestMain.sellResult, 2)
					backtestMain.iTotalTrades++
					backtestMain.iSell++
					backtestMain.sellFlag = true
					backtestMain.updateSell = false
					backtestMain.tpSell = 0.
					backtestMain.slSell = 0.
					backtestMain.sellPrice = 0.
				}
			}
		}
	}

	var totalTrades []float64 = conversion.RemoveExcessZeros(backtestMain.totalTrades)
	var buyTrades []float64 = conversion.RemoveExcessZeros(backtestMain.buyTrades)
	var sellTrades []float64 = conversion.RemoveExcessZeros(backtestMain.sellTrades)

	return totalTrades, buyTrades, sellTrades

}

func NettingBacktest(dt *data.LayoutData, multiplyTpB float64, multiplySlB float64, balance float64, sizeBuy uint32, sizeSell uint32, jpy bool) ([]float64, []float64, []float64) {
	// Do netting backtest and return total trades, buy trades, sell trades

	var backtestMain backtestVariables = backtestHeart(balance, sizeBuy, sizeSell)

	for i := 1; i < len((*dt).Open); i++ {

		if (*dt).Pricetf[i] != (*dt).Pricetf[i-1] {
			if (*dt).BuyFlag[i] && backtestMain.universalFlag && backtestMain.buyFlag {
				backtestMain.buyPrice = (*dt).Pricetf[i]
				backtestMain.tpBuy, backtestMain.slBuy = tpslCalculationNonFix((*dt).Pricetf[i], multiplyTpB, multiplySlB, (*dt).Tpsl[i], jpy, true)
				backtestMain.buyFlag = false
				backtestMain.universalFlag = false
			}
			if (*dt).SellFlag[i] && backtestMain.universalFlag && backtestMain.sellFlag {
				backtestMain.sellPrice = (*dt).Pricetf[i]
				backtestMain.tpSell, backtestMain.slSell = tpslCalculationNonFix((*dt).Pricetf[i], multiplyTpB, multiplySlB, (*dt).Tpsl[i], jpy, false)
				backtestMain.sellFlag = false
				backtestMain.universalFlag = false
			}
		} else if (*dt).Pricetf[i] == (*dt).Pricetf[i-1] {
			if !backtestMain.universalFlag && !backtestMain.buyFlag {
				if checkTpslExit((*dt).Open[i], (*dt).High[i], (*dt).Low[i], (*dt).Close[i], backtestMain.slBuy, false) {
					backtestMain.capital, backtestMain.buyResult = financeCalculation(backtestMain.capital, backtestMain.slBuy, backtestMain.buyPrice, backtestMain.slBuy, balance, false)
					backtestMain.updateBuy = true
				} else if checkTpslExit((*dt).Open[i], (*dt).High[i], (*dt).Low[i], (*dt).Close[i], backtestMain.tpBuy, true) {
					backtestMain.capital, backtestMain.buyResult = financeCalculation(backtestMain.capital, backtestMain.tpBuy, backtestMain.buyPrice, backtestMain.tpBuy, balance, false)
					backtestMain.updateBuy = true
				}
				if backtestMain.updateBuy {
					backtestMain.totalTrades[backtestMain.iTotalTrades] = backtestMain.capital
					backtestMain.buyTrades[backtestMain.iBuy] = conversion.Round(backtestMain.buyTrades[backtestMain.iBuy-1]+backtestMain.buyResult, 2)
					backtestMain.iTotalTrades++
					backtestMain.iBuy++
					backtestMain.buyFlag = true
					backtestMain.universalFlag = true
					backtestMain.updateBuy = false
					backtestMain.tpBuy = 0.
					backtestMain.slBuy = 0.
					backtestMain.buyPrice = 0.
				}
			}
			if !backtestMain.universalFlag && !backtestMain.sellFlag {
				if checkTpslExit((*dt).Open[i], (*dt).High[i], (*dt).Low[i], (*dt).Close[i], backtestMain.slSell, true) {
					backtestMain.capital, backtestMain.sellResult = financeCalculation(backtestMain.capital, backtestMain.sellPrice, backtestMain.slSell, backtestMain.sellPrice, balance, false)
					backtestMain.updateSell = true
				} else if checkTpslExit((*dt).Open[i], (*dt).High[i], (*dt).Low[i], (*dt).Close[i], backtestMain.tpSell, false) {
					backtestMain.capital, backtestMain.sellResult = financeCalculation(backtestMain.capital, backtestMain.sellPrice, backtestMain.tpSell, backtestMain.sellPrice, balance, false)
					backtestMain.updateSell = true
				}
				if backtestMain.updateSell {
					backtestMain.totalTrades[backtestMain.iTotalTrades] = backtestMain.capital
					backtestMain.sellTrades[backtestMain.iSell] = conversion.Round(backtestMain.sellTrades[backtestMain.iSell-1]+backtestMain.sellResult, 2)
					backtestMain.iTotalTrades++
					backtestMain.iSell++
					backtestMain.sellFlag = true
					backtestMain.universalFlag = true
					backtestMain.updateSell = false
					backtestMain.tpSell = 0.
					backtestMain.slSell = 0.
					backtestMain.sellPrice = 0.
				}
			}
		}
	}

	var totalTrades []float64 = conversion.RemoveExcessZeros(backtestMain.totalTrades)
	var buyTrades []float64 = conversion.RemoveExcessZeros(backtestMain.buyTrades)
	var sellTrades []float64 = conversion.RemoveExcessZeros(backtestMain.sellTrades)

	return totalTrades, buyTrades, sellTrades

}
