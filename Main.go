package main

func CalculateLBTT(price float64) float64 {
	var tax float64

	if price <= 145000 {
		tax = 0
	} else if price <= 250000 {
		tax = (price - 145000) * 0.02
	} else if price <= 325000 {
		tax = (price-250000)*0.05 + 2100
	} else if price <= 750000 {
		tax = (price-325000)*0.10 + 5850
	} else {
		tax = (price-750000)*0.12 + 48350
	}
	//Numbers being added are the amount taxed by default when the amount is above a threshold

	return tax
}
