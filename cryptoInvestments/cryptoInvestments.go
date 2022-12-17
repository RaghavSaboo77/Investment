package cryptoInvestments

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type CryptoApiResponse struct {
    data []*CryptoData
}

type CryptoData struct{
  id int
  name string
  symbol string
}

func CalculateCryptoInvestments(investmentAmount float64) {
  client := &http.Client{}
  req, err := http.NewRequest("GET","https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
  if err != nil {
    log.Print(err)
    os.Exit(1)
  }

  q := url.Values{}
  q.Add("start", "1")
  q.Add("limit", "3")
  q.Add("convert", "INR")

  req.Header.Set("Accepts", "application/json")
  req.Header.Add("X-CMC_PRO_API_KEY", "0e57b309-0bcb-44f3-b04b-2fc7f0ab149d")
  req.URL.RawQuery = q.Encode()

  resp, err := client.Do(req);
  if err != nil {
    fmt.Println("Error sending request to server")
    os.Exit(1)
  }
  fmt.Println(resp.Status);
  respBody, _ := ioutil.ReadAll(resp.Body)
  fmt.Println(string(respBody));
}