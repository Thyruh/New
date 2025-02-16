package main

func calculateFinalBill(costPerMessage float64, numMessages int) float64 {
	costPerMessage = calculateDiscountRate(numMessages)
	return costPerMessage * float64(numMessages)
}


func calculateDiscountRate(messagesSent int) float64 {
	if messagesSent > 5000 {
		return .8
	}
	if messagesSent > 1000 {
		return .9
	}
	return 1

}


func calculateBaseBill(costPerMessage float64, messagesSent int) float64 {
	return costPerMessage * float64(messagesSent)
}

