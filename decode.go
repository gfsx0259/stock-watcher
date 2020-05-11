package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Row struct {
	BuyPrice  float32 `xml:"BID,attr"`
	SalePrice float32 `xml:"OFFER,attr"`
	Ticker    string  `xml:"SECID,attr"`
}

func (row Row) String() string {
	return fmt.Sprintf("{buyPrice: %f, salePrice: %f, ticker: %s}\n", row.BuyPrice, row.SalePrice, row.Ticker)
}

type Data struct {
	Rows []Row  `xml:"rows>row"`
	Id   string `xml:"id,attr"`
}
type Document struct {
	Data []Data `xml:"data"`
}

const ApiEndpoint = "https://iss.moex.com/iss/engines/stock/markets/shares/boards/TQBR/securities.xml"

func main() {
	response, err := http.Get(ApiEndpoint)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	doc := Document{}
	err = xml.Unmarshal(body, &doc)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	var stocks []Row
	for _, data := range doc.Data {
		if data.Id == "marketdata" {
			stocks = data.Rows
		}
	}

	indexedStocks := make(map[string]Row)

	for _, row := range stocks {
		indexedStocks[row.Ticker] = row
	}
	fmt.Println(indexedStocks)
}
