package round

import (
	"github.com/shopspring/decimal"
)

func round() string {
	quantity := decimal.NewFromInt(3)

	return quantity.StringFixed(2)
}
