package changenow

import (
	"fmt"
	"strings"
)

// Currencies ...
// GET /currencies
func (c *ChangeNow) Currencies() ([]*Currency, error) {
	path := "/currencies"
	currencies := make([]*Currency, 0)
	_, err := c.client.R().SetResult(&currencies).Get(path)
	if err != nil {
		return nil, err
	}
	return currencies, nil
}

// EstimatedExchangeAmount ...
// GET /exchange-amount/:send_amount/:from_to
func (c *ChangeNow) EstimatedExchangeAmount(amount float64, source, target string) (*EstimatedExchangeAmount, error) {
	path := fmt.Sprintf("/exchange-amount/%s/%s", floatToString(amount), fromTo(source, target))

	response := new(EstimatedExchangeAmount)
	_, err := c.client.R().SetResult(&response).Get(path)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// CurrenciesTo ...
// /currencies-to/:ticker
// List of available currencies for specific currency
// This API endpoint returns the list of available currencies. Some currencies get enabled or disabled from time to time, so make sure to refresh the list occasionally.
func (c *ChangeNow) CurrenciesTo(ticker string) ([]*Currency, error) {
	path := "/currencies-to/" + ticker

	currencies := make([]*Currency, 0)
	_, err := c.client.R().SetResult(&currencies).Get(path)
	if err != nil {
		return nil, err
	}
	return currencies, nil
}

// CurrencyInfo ...
// GET /currencies/:ticker
func (c *ChangeNow) CurrencyInfo(ticker string) (*CurrencyInfo, error) {
	path := "/currencies/" + strings.ToLower(ticker)

	result := new(CurrencyInfo)
	_, err := c.client.R().SetResult(&result).Get(path)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MinAmount ...
// GET /min-amount/:from_to
// This func is the one that mentions if the pair is inactive
func (c *ChangeNow) MinAmount(source, target string) (float64, error) {
	name := fromTo(source, target)
	path := "/min-amount/" + name

	result := new(MinAmount)

	// Make the request, setting the response and potential error from statuscodes > 399
	response, err := c.client.R().SetResult(&result).SetError(&RequestError{}).Get(path)
	if e := checkResponse(response, err); e != nil {
		if e.Error() == "pair_is_inactive" {
			getPair(name).IsActive = false
		}
		return 0, e
	}

	// Cache the amount
	getPair(name).MinAmount = result.Amount

	return result.Amount, nil
}

// AvailablePairs ...
// GET /api/v1/market-info/available-pairs/
func (c *ChangeNow) AvailablePairs() ([]string, error) {
	path := "/market-info/available-pairs"

	result := make([]string, 0)
	// Make the request, setting the response and potential error from statuscodes > 399
	response, err := c.client.R().SetResult(&result).SetError(&RequestError{}).Get(path)
	if e := checkResponse(response, err); e != nil {
		return nil, e
	}

	// cache. we cannot know here whether or not it is inactive
	// you must call MinAmount for that
	for _, name := range result {
		addPair(name)
	}

	fmt.Println(pairs)

	return result, nil
}

// CreateTransaction ...
// POST /api/v1/transactions/:api_key
// Errors encountered:
// - not_valid_params
// - cannot_create_transaction (when pair is not available?)
// - not_valid_address
// - out_of_range
func (c *ChangeNow) CreateTransaction(request *TransactionRequest) (*Transaction, error) {
	path := "/transactions/" + c.apikey

	// normalize
	request.From = strings.ToLower(request.From)
	request.To = strings.ToLower(request.To)

	transaction := new(Transaction)
	response, err := c.client.R().SetResult(transaction).SetError(&RequestError{}).SetBody(request).Post(path)
	if e := checkResponse(response, err); e != nil {
		return nil, e
	}
	return transaction, nil
}

// ListTransactions ...
// GET /api/v1/transactions/:api_key
func (c *ChangeNow) ListTransactions() ([]*TransactionDetails, error) {
	path := "/transactions/" + c.apikey

	transactions := make([]*TransactionDetails, 0)
	_, err := c.client.R().SetResult(&transactions).Get(path)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

// TransactionStatus ...
// GET /api/v1/transactions/:id/:api_key
func (c *ChangeNow) TransactionStatus(id string) (*TransactionDetails, error) {
	path := fmt.Sprintf("/transactions/%s/%s", id, c.apikey)
	transaction := new(TransactionDetails)
	response, err := c.client.R().SetResult(transaction).SetError(&RequestError{}).Get(path)
	if e := checkResponse(response, err); e != nil {
		return nil, e
	}
	return transaction, nil
}
