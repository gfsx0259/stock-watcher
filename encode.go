package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type MarketData struct {
	SalePrice float32 `xml:"salePrice"`
	BuyPrice float32 `xml:"buyPrice"`
}

func main() {
	marketData := MarketData{
		SalePrice: 10,
		BuyPrice:  8,
	}

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "    ")
	err := enc.Encode(marketData)

	if err != nil {
		fmt.Println(err)
	}
}