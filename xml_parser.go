package gomoney

import (
	"encoding/xml"
	"io"
	"time"
)

type xmlEurExchangeRate struct {
	OuterCube xmlOuterCube `xml:"Cube"`
}

type xmlOuterCube struct {
	InnerCube xmlInnerCube `xml:"Cube"`
}

type xmlInnerCube struct {
	Time      string        `xml:"time,attr"`
	CurrCubes []xmlCurrCube `xml:"Cube"`
}

type xmlCurrCube struct {
	Currency     string  `xml:"currency,attr"`
	ExchangeRate float64 `xml:"rate,attr"`
}

// Parses the response into an exchangeRate struct
func parseSOAPResponse(data io.Reader) (*eurExchangeRate, error) {
	xmlTarget := &xmlEurExchangeRate{}
	if err := xml.NewDecoder(data).Decode(xmlTarget); err != nil {
		return nil, err
	}
	exgDate, _ := time.Parse("2006-01-02", xmlTarget.OuterCube.InnerCube.Time)
	result := &eurExchangeRate{
		Date:  exgDate,
		Rates: map[string]float64{"EUR": 1},
	}

	for _, currCube := range xmlTarget.OuterCube.InnerCube.CurrCubes {
		result.Rates[currCube.Currency] = currCube.ExchangeRate
	}
	return result, nil
}
