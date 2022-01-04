package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func flashGreen() {
	for _, n := range []int{0, 2, 5, 7, 9} {
		all[n].High()
	}

	time.Sleep(time.Second / 3)

	for _, n := range []int{0, 2, 5, 7, 9} {
		all[n].Low()
	}
}

func flashRed() {
	for _, n := range []int{1, 3, 4, 6, 8} {
		all[n].High()
	}

	time.Sleep(time.Second / 3)

	for _, n := range []int{1, 3, 4, 6, 8} {
		all[n].Low()
	}
}

func btcMinute(apiKey string) {
	// flash green if btc is up over the last reading in one minute
	// I don't really care about cryptocurrency but I knew there were easy APIs
	// and a lot of people seem to go off about it, so

	// now get our initial value
	lastPrice := btcValueNow(apiKey)
	fmt.Printf("start price: %f\n", lastPrice)

	// to show it started, flash back and forth
	for i := 0; i < 3; i++ {
		star.High()
		flashGreen()

		time.Sleep(time.Second / 3)

		star.Low()
		flashRed()

		time.Sleep(time.Second / 3)
	}

	for true {
		start := time.Now()
		waitFlag := true
		for waitFlag {
			delta := time.Now().Sub(start)
			fmt.Printf("Delta: %s\n", delta.String())
			if delta > (time.Minute + (5 * time.Second)) {
				waitFlag = false
			}

			// flash while we wait
			star.Low()
			flashGreen()
			time.Sleep(time.Second / 2)
			flashRed()
			star.High()
		}

		// we've waited a minute with a buffer to let the API update
		newPrice := btcValueNow(apiKey)

		if newPrice < lastPrice {
			fmt.Printf("price decreased: %f to %f\n", lastPrice, newPrice)
			for i := 0; i < 10; i++ {
				flashRed()
				time.Sleep(time.Second / 2)
			}
		} else {
			fmt.Printf("price increased: %f to %f\n", lastPrice, newPrice)
			for i := 0; i < 10; i++ {
				flashGreen()
				time.Sleep(time.Second / 2)
			}
		}

		lastPrice = newPrice
	}
}

func btcValueNow(apiKey string) float64 {
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest", nil)
	if err != nil {
		panic(err)
	}

	q := url.Values{}
	q.Add("symbol", "BTC")
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", apiKey)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var decodeMap map[string]interface{}
	err = json.Unmarshal(bodyBytes, &decodeMap)
	if err != nil {
		panic(err)
	}

	dataObj := decodeMap["data"].(map[string]interface{})
	btcObj := dataObj["BTC"].(map[string]interface{})
	quoteObj := btcObj["quote"].(map[string]interface{})
	usdObj := quoteObj["USD"].(map[string]interface{})

	return usdObj["price"].(float64)
}
