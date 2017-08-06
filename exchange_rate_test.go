package gomoney

import (
	"io/ioutil"
	"testing"

	httpmock "gopkg.in/jarcoal/httpmock.v1"

	"github.com/stretchr/testify/assert"
)

func TestNewEurExchangeRate(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	t.Log("When the endpoint is working")
	data, err := ioutil.ReadFile("./fixtures/exchange_rates/1.xml")
	if err != nil {
		panic(err)
	}
	httpmock.RegisterResponder("GET", "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml",
		httpmock.NewStringResponder(200, string(data)))

	er, err := newEurExchangeRate()
	assert.Nil(t, err, "Expected no error to be returned")
	assert.Equal(t, expectedEurExchangeRate1(), er, "Eur exchange rate does not match")
}
