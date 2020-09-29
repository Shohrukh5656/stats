package stats

import "github.com/Shohrukh5656/bank1/pkg/types"

func Avg(payments []types.Payment) types.Money {
	countPayments := types.Money(len(payments))
	sumPaymenys := types.Money(0)
	for _, payment := range payments {
		moneyPayments := payment.Amount
		sumPaymenys += moneyPayments
	}
	return sumPaymenys / countPayments
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sumPaymenys := types.Money(0)
	for _, payment := range payments {
		if payment.Category != category {
			continue
		}
		moneyPayments := payment.Amount
		sumPaymenys += moneyPayments
	}
	return sumPaymenys
}
