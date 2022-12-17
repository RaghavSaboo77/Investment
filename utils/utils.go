package utils


import (
	"fmt"
)


func DisplayInvestmentCategories(name string,CategoryToAmountMap  map[string]float64,categories []string){
	fmt.Println(name)
	for _, category := range categories {
		if CategoryToAmountMap[category] == 0.0{
			continue
		}
		fmt.Printf("Ratio : %.2f \t Category: %s\n", CategoryToAmountMap[category], category)
	}
}

func ComputeFinalAmountTobeInvested(amountLeftForInvestment float64, categories []string, amountToBeInvestedforRatio map[string]float64, idealInvestmentRatio map[string]float64) map[string]float64 {
	finalAmountToBeInvestedForCategory := amountToBeInvestedforRatio
	for _, category := range categories {
		finalAmountToBeInvestedForCategory[category] += amountLeftForInvestment * (idealInvestmentRatio[category] / 100)
	}
	return finalAmountToBeInvestedForCategory
}

func GetAmountToBeInvested(categories []string, maxTotalAmountCategory float64, idealInvestmentRatio map[string]float64, currentInvestmentValue map[string]float64) (map[string]float64, float64) {
	amountToBeInvestedforRatio := make(map[string]float64)
	totalAmountNeeded := 0.0
	for _, category := range categories {
		requiredValueCategory := (maxTotalAmountCategory * idealInvestmentRatio[category]) / 100
		differenceInCategory := requiredValueCategory - currentInvestmentValue[category]
		amountToBeInvestedforRatio[category] = differenceInCategory
		totalAmountNeeded += differenceInCategory
	}
	return amountToBeInvestedforRatio, totalAmountNeeded
}

func ComputeNewIdealInvestmentRatioOutOf100(categories []string, idealInvesmentRatio map[string]float64) map[string]float64 {
	totalRatio := 0.0
	for _, category := range categories {
		totalRatio += idealInvesmentRatio[category]
	}

	for _, category := range categories {
		idealInvesmentRatio[category] = (idealInvesmentRatio[category] * 100.0) / totalRatio
	}
	return idealInvesmentRatio
}

func GetMaxCategoryDetails(categories []string, currentInvestmentValue map[string]float64, idealInvestmentRatio map[string]float64) (float64, string) {
	maxTotalAmountCategory := 0.0
	maxCategory := ""
	for _, category := range categories {
		totalAmountCategory := (currentInvestmentValue[category] * 100) / idealInvestmentRatio[category]
		if totalAmountCategory > maxTotalAmountCategory {
			maxTotalAmountCategory = totalAmountCategory
			maxCategory = category
		}
	}
	return maxTotalAmountCategory, maxCategory
}

func Remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func GetDataOnInvestment(currentInvestmentValue map[string]float64, categories []string) (float64, map[string]float64) {
	totalInvestment := 0.0
	for _, investmentCategory := range categories {
		totalInvestment += currentInvestmentValue[investmentCategory]
	}

	currentInvestmentRatio := make(map[string]float64)
	for _, investmentCategory := range categories {
		currentInvestmentRatio[investmentCategory] = (currentInvestmentValue[investmentCategory] / totalInvestment) * 100.0
	}
	return totalInvestment, currentInvestmentRatio
}



func CalculateCryptoInvestments() string{
	fmt.Print("Hello")
	return "hello"
	// client := &http.Client{}
	// req, err := http.NewRequest("GET","https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
	// if err != nil {
	//   log.Print(err)
	//   os.Exit(1)
	// }
  
	// q := url.Values{}
	// q.Add("start", "1")
	// q.Add("limit", "3")
	// q.Add("convert", "INR")
  
	// req.Header.Set("Accepts", "application/json")
	// req.Header.Add("X-CMC_PRO_API_KEY", "0e57b309-0bcb-44f3-b04b-2fc7f0ab149d")
	// req.URL.RawQuery = q.Encode()
  
  
	// resp, err := client.Do(req);
	// if err != nil {
	//   fmt.Println("Error sending request to server")
	//   os.Exit(1)
	// }
	// fmt.Println(resp.Status);
	// respBody, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(respBody));
  
  }