package main

import (
	"errors"
)

type TaxBracket struct {
	MinimumPrice  float64
	MaximumPrice  float64
	TaxMultiplier float64
}

var TaxBrackets []TaxBracket = []TaxBracket{
	TaxBracket{0.0, 145000, 0.0},
	TaxBracket{145000, 250000, 0.02},
	TaxBracket{250000, 325000, 0.05},
	TaxBracket{325000, 750000, 0.1},
	TaxBracket{750000, 99999999, 0.12},
}

func CalculateLBTT(price float64) (float64, error) {
	var tax float64

	if price < 0 {
		return 0, errors.New("negative number")
	}

	for i := 0; i < len(TaxBrackets); i++ {
		bracket := TaxBrackets[i]
		if price > bracket.MinimumPrice {
			if price <= bracket.MaximumPrice {
				TaxForBand := (price - bracket.MinimumPrice) * bracket.TaxMultiplier
				tax += TaxForBand
			} else {
				TaxForBand := (bracket.MaximumPrice - bracket.MinimumPrice) * bracket.TaxMultiplier
				tax += TaxForBand
			}
		}

	}

	return tax, nil
}
