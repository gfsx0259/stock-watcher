package main

import(
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Row struct {
	BuyPrice float32 `xml:"BID,attr"`
	SalePrice float32 `xml:"OFFER,attr"`
	Ticker string `xml:"SECID,attr"`
}

type Rows struct {
	Row []Row `xml:"row"`
}

type Data struct {
	Rows Rows `xml:"rows"`
}

type Document struct {
	Data []Data `xml:"data"`
}

func main() {
	var err error

	data, err := ioutil.ReadFile("decode.xml")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	emptyDoc := Document{}

	err = xml.Unmarshal(data, &emptyDoc)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(emptyDoc.Data)
}
