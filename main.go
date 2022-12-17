package main

import (
	"cnst"
	"fmt"
	"investmentHistory"
	"utils"
)

//This code will clearly tell you how much to invest where
func main() {
	//Ideal Investment
	idealInvestmentRatio := investmentHistory.GetInvestment(cnst.IdealInvestmentRatio)
	var categories = investmentHistory.GetCategories()
	totalRatio := 0.0
	for _, category := range categories {
		totalRatio += idealInvestmentRatio[category]
	}
	if totalRatio != 100.0 {
		fmt.Println("The Total Ideal Investment is not equal to 100 it is", totalRatio)
		idealInvestmentRatio = utils.ComputeNewIdealInvestmentRatioOutOf100(categories, idealInvestmentRatio)
	}
	utils.DisplayInvestmentCategories("Your Ideal Investment Ratio is:-", idealInvestmentRatio, categories)

	//Current Investment Information                 
	currentInvestmentValue := investmentHistory.GetInvestment("Month5")
	totalInvestment, currentInvestmentRatio := utils.GetDataOnInvestment(currentInvestmentValue, categories)
	fmt.Println("\nYour Current total Net Investment is rupees", totalInvestment)
	utils.DisplayInvestmentCategories("Current Investment Ratio", currentInvestmentRatio, categories)
	investmentAmount := currentInvestmentValue[cnst.TotalAmountToInvestNow]

	//Remove Max Weightage Categories To get the proportion of the low weightage categories to invest in for the given Investment Amount
	maxTotalAmountCategory, maxCategory := utils.GetMaxCategoryDetails(categories, currentInvestmentValue, idealInvestmentRatio)
	fmt.Println("The Category with the max weightage now is: ", maxCategory, " and total amount needed to balance accordingly is ", maxTotalAmountCategory)
	amountToBeInvestedforRatio, totalAmountNeededForRatio := utils.GetAmountToBeInvested(categories, maxTotalAmountCategory, idealInvestmentRatio, currentInvestmentValue)
	for totalAmountNeededForRatio > investmentAmount {
		categories = utils.Remove(categories, maxCategory)
		idealInvestmentRatio = utils.ComputeNewIdealInvestmentRatioOutOf100(categories, idealInvestmentRatio)
		maxTotalAmountCategory, maxCategory = utils.GetMaxCategoryDetails(categories, currentInvestmentValue, idealInvestmentRatio)
		amountToBeInvestedforRatio, totalAmountNeededForRatio = utils.GetAmountToBeInvested(categories, maxTotalAmountCategory, idealInvestmentRatio, currentInvestmentValue)
	}

	//Distrubute the remaining amount after the totalAmountNeededForRatio is removed from the investment amount
	fmt.Println("\nFinal Amount to be invested is ", investmentAmount)
	finalAmountToBeInvestedForCategory := utils.ComputeFinalAmountTobeInvested(investmentAmount-totalAmountNeededForRatio, categories, amountToBeInvestedforRatio, idealInvestmentRatio)
	utils.DisplayInvestmentCategories("Investment Amount to be invested in the following ratio", finalAmountToBeInvestedForCategory, categories)

	// cryptoInvestmentAmount := 20000.00
	// cryptoInvestments.CalculateCryptoInvestments(cryptoInvestmentAmount)
}
