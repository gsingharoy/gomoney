# Gomoney

A library implemented in golang to convert money into other currencies

[![Build Status](https://travis-ci.org/gsingharoy/gomoney.svg?branch=master)](https://travis-ci.org/gsingharoy/gomoney)

## Description
This package contains a simple library which converts a currency amount to another supported currency. The exchange rates used here are published by the official [European Central Bank](http://www.ecb.europa.eu/stats/exchange/eurofxref/html/index.en.html).

## Usage

```go
import "github.com/gsingharoy/gomoney"

// generate a new money object
// Note: Check the section below for the supported currencies
m, err := gomoney.NewMoney(100, "USD")

// now convert the money to another currency
// convertedAmount will contain the conversion of 100 USD to CHF
convertedAmount, err := m.Convert("CHF")
```

The following currencies are supported :

* `EUR` : Euro
* `USD` : US dollar
* `JPY` : Japanese yen
* `BGN` : Bulgarian lev
* `CZK` : Czech koruna
* `DKK` : Danish krone
* `GBP` : Pound sterling
* `HUF` : Hungarian forint
* `PLN` : Polish zloty
* `RON` : Romanian leu
* `SEK` : Swedish krona
* `CHF` : Swiss franc
* `NOK` : Norwegian krone
* `HRK` : Croatian kuna
* `RUB` : Russian rouble
* `TRY` : Turkish lira
* `AUD` : Australian dollar
* `BRL` : Brazilian real
* `CAD` : Canadian dollar
* `CNY` : Chinese yuan renminbi
* `HKD` : Hong Kong dollar
* `IDR` : Indonesian rupiah
* `ILS` : Israeli shekel
* `INR` : Indian rupee
* `KRW` : South Korean won
* `MXN` : Mexican peso
* `MYR` : Malaysian ringgit
* `NZD` : New Zealand dollar
* `PHP` : Philippine peso
* `SGD` : Singapore dollar
* `THB` : Thai baht
* `ZAR` : South African rand
