package changenow

import (
	"encoding/json"
	"fmt"
	"testing"
)

var apikey = ""

func TestCurrencies(t *testing.T) {
	c := New(apikey)
	currencies, err := c.Currencies()
	if err != nil {
		t.Fatal(err)
	}

	print(currencies)
}

func TestEstimatedExchangeAmount(t *testing.T) {
	c := New(apikey)
	estimate, err := c.EstimatedExchangeAmount(0.1, "XZC", "USDC")
	if err != nil {
		t.Fatal(err)
	}

	print(estimate)
}

func TestCurrenciesForTicker(t *testing.T) {
	c := New(apikey)
	resp, err := c.CurrenciesTo("XZC")
	if err != nil {
		t.Fatal(err)
	}
	print(resp)
}

func TestCurrencyInfo(t *testing.T) {
	c := New(apikey)
	resp, err := c.CurrencyInfo("XZC")
	if err != nil {
		t.Fatal(err)
	}
	print(resp)
}

func TestMinAmount(t *testing.T) {
	c := New(apikey)
	resp, err := c.MinAmount("XZC", "ETH")
	if err != nil {
		t.Fatal(err)
	}
	print(resp)
}

func TestAvailablePairs(t *testing.T) {
	c := New(apikey)
	resp, err := c.AvailablePairs()
	if err != nil {
		t.Fatal(err)
	}
	print(resp)
}

func TestCreateTransactions(t *testing.T) {
	c := New(apikey)
	min, err := c.MinAmount("xzc", "eth")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(min)

	resp, err := c.CreateTransaction(&TransactionRequest{
		From:    "XZC",
		To:      "ETH",
		Amount:  4,
		Address: "0x23",
	})
	if err != nil {
		t.Fatal(err)
	}
	print(resp)
}

func TestListTransactions(t *testing.T) {
	c := New(apikey)
	resp, err := c.ListTransactions()
	if err != nil {
		t.Fatal(err)
	}
	print(resp)
}

func TestTransactionStatus(t *testing.T) {
	id := ""
	c := New(apikey)
	resp, err := c.TransactionStatus(id)
	if err != nil {
		t.Fatal(err)
	}
	print(resp)
}

func print(o interface{}) {
	j, _ := json.MarshalIndent(o, "", "  ")
	fmt.Println(string(j))
}
