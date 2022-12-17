package investmentHistory

import "cnst"

func GetCategories() []string {
	return []string{
		"US_Total_Stock",
		"India_Large_Cap_Nifty_50",
		"India_Large_Cap_Nifty_Next_50",
		"India_Mid_Cap",
		"India_Small_Cap",
		"India_Bank_Index",
		"Crypto_Index",
		"Soverign_Gold_Bond",
		"Employee_Provident_Fund",
		"Public_Provident_Fund",
		"National_Pension_System",
		"Equity_Linked_Saving_Scheme",
		"Cash",
		"Fixed_Deposits",
	}
}

//Mission Statement :- I beleive and am confident that in the next 40 years the world will grow
//Strategy:- keep investing as much as you can
//Give no fucks if the market is up or down-> doesn't matter -> Even if it goes down you will get the right things for the discounted price
//Volatility is the price of the market
//After recession is done 5% to soverign gold bonds and 15% to crypto

func GetInvestment(description string) map[string]float64 {
	switch description {
	case cnst.IdealInvestmentRatio:
		return map[string]float64{
			//Aim get the average return of the Indian and The US Stock market consistently without selling anything
			//Ideal Ratio India:40, America:40 and Crypto 20 all in index funds
			//Extra Indian Percentage to be added to NiftyBankIndex
			cnst.US_Total_Stock:                45,
			cnst.India_Large_Cap_Nifty_50:      (66.8) * 0.45,
			cnst.India_Large_Cap_Nifty_Next_50: 10 * 0.45,
			cnst.India_Mid_Cap:                 (12.9) * 0.45,
			cnst.India_Small_Cap:               (6.4) * 0.45,
			cnst.India_Bank:                    (100 - 66.8 - 10 - 12.9 - 6.4) * 0.45,
			cnst.Crypto_Index:                  10,
		}
	//Rule :- Only work on investment procedures for max of 2 hrs on the closest Saturday of 13 th (Other Than this don't even think about investments)
	case "Month1": //9th July 2022
		return map[string]float64{
			cnst.US_Total_Stock:                1447,
			cnst.India_Large_Cap_Nifty_50:      476,
			cnst.India_Large_Cap_Nifty_Next_50: 481,
			cnst.India_Mid_Cap:                 1421,
			cnst.India_Small_Cap:               2329,
			cnst.Crypto_Index:                  0,
			cnst.TotalAmountToInvestNow:        53800,
		}
	case "Month2": //11th August 2022
	return map[string]float64{
		cnst.US_Total_Stock:                31987,
		cnst.India_Large_Cap_Nifty_50:      21681,
		cnst.India_Large_Cap_Nifty_Next_50: 3270,
		cnst.India_Mid_Cap:                 4268,
		cnst.India_Small_Cap:               2547,
		cnst.India_Bank:                    1276,
		cnst.Crypto_Index:                  0,
		cnst.TotalAmountToInvestNow:        30000,
	}
	case "Mumma1": //15th August 2022
	return map[string]float64{
		cnst.US_Total_Stock:                0,
		cnst.India_Large_Cap_Nifty_50:      0,
		cnst.India_Large_Cap_Nifty_Next_50: 0,
		cnst.India_Mid_Cap:                 0,
		cnst.India_Small_Cap:               0,
		cnst.India_Bank:                    0,
		cnst.TotalAmountToInvestNow:        15000,
	}
	case "Mumma2": //26th August 2022
	return map[string]float64{
		cnst.US_Total_Stock:                7626,
		cnst.India_Large_Cap_Nifty_50:      5018,
		cnst.India_Large_Cap_Nifty_Next_50: 758,
		cnst.India_Mid_Cap:                 986,
		cnst.India_Small_Cap:               491,
		cnst.India_Bank:                    508,
		cnst.TotalAmountToInvestNow:        18000,
	}
	case "Month3": //8th September 2022
	return map[string]float64{
		cnst.US_Total_Stock:                45846,
		cnst.India_Large_Cap_Nifty_50:      28691,
		cnst.India_Large_Cap_Nifty_Next_50: 4506,
		cnst.India_Mid_Cap:                 5789,
		cnst.India_Small_Cap:               2937,
		cnst.India_Bank:                    1720,
		cnst.Crypto_Index:                  8995,
		cnst.TotalAmountToInvestNow:        25000,
	}
	case "Mumma3": //8th September 2022
	return map[string]float64{
		cnst.US_Total_Stock:                16161,
		cnst.India_Large_Cap_Nifty_50:      11357,
		cnst.India_Large_Cap_Nifty_Next_50: 1716,
		cnst.India_Mid_Cap:                 2258,
		cnst.India_Small_Cap:               1114,
		cnst.India_Bank:                    669,
		cnst.Crypto_Index:                  0,
		cnst.TotalAmountToInvestNow:        10000,
	}
	case "Month4": //1st October 2022
	return map[string]float64{
		cnst.US_Total_Stock:                51242,
		cnst.India_Large_Cap_Nifty_50:      34889,
		cnst.India_Large_Cap_Nifty_Next_50: 5244,
		cnst.India_Mid_Cap:                 6763,
		cnst.India_Small_Cap:               3342,
		cnst.India_Bank:                    2042,
		cnst.Crypto_Index:                  16000,
		cnst.TotalAmountToInvestNow:        25000,
	}
	case "Month5": //3rd Nov 2022
	return map[string]float64{
		cnst.US_Total_Stock:                71130,
		cnst.India_Large_Cap_Nifty_50:      46171,
		cnst.India_Large_Cap_Nifty_Next_50: 6596,
		cnst.India_Mid_Cap:                 8625,
		cnst.India_Small_Cap:               4249,
		cnst.India_Bank:                    2725,
		cnst.Crypto_Index:                  17517,
		cnst.TotalAmountToInvestNow:        20000,
	}
	}
	return nil
}
