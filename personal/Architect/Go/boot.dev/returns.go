package main

func getProductMessage(tier string) string {
	quantityMsg, priceMsg := getProductInfo(tier)
	return "You get " + quantityMsg + " for " + priceMsg + "."
}

// don't touch below this line

func getProductInfo(tier string) (string, string) {
	if tier == "basic" {
		return "1,000 texts per month", "$30 per month"
	} else if tier == "premium" {
		return "50,000 texts per month", "$60 per month"
	} else if tier == "enterprise" {
		return "unlimited texts per month", "$100 per month"
	} else {
		return "", ""
	}
}
