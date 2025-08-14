package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Maryam-nokohan/GoldCurrentPrice/config"
	"github.com/Maryam-nokohan/GoldCurrentPrice/massager"
	"github.com/Maryam-nokohan/GoldCurrentPrice/model"
)

const apiKey = "goldapi-10mc7psmd7stwiu-io"

type GoldPriceResponse struct {
	Price float64 `json:"price"`
}

func fetchGoldPrice() (float64, error) {
	url := "https://www.goldapi.io/api/XAU/USD"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}

	req.Header.Set("x-access-token", apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("error status: %s", resp.Status)
	}

	var result GoldPriceResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return 0, err
	}

	return result.Price, nil
}

func main() {
	price, err := fetchGoldPrice()
	if err != nil {
		fmt.Println("Error fetching gold price:", err)
		return
	}

	fmt.Printf("Current gold price in USD: $%.2f\n", price)

	// Mail sender logic
	sender := massager.NewMail(config.GetFromEnv("EMAIL_ADDRESS"), config.GetFromEnv("EMAIL_KEY"))
	sender.Send(config.GetFromEnv("EMAIL_RECEIVER") , &model.Message{
		Header: "Gold current price",
		Body: fmt.Sprintf("hello\nthe current price of gold in USD is %.2f\n" , price),
	})

}
