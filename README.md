<img src="http://crypdex.io/img/full-logo.svg" width=300 style="margin-bottom:20px;"/>

# ChangeNOW API in Go
A Golang wrapper for the ChangeNOW instant coin exchange


```bash
go get github.com/crypdex/go-changenow
```

This package is currently used in production, but is not yet locked down by version and is subject to change.

## Implemention Status (v1)

- Standard flow
    - [x] [Estimated exchange amount](https://changenow.io/api/docs/standard-flow/method-exchange-amount)
    - [x] [Create transaction](https://changenow.io/api/docs/standard-flow/method-create-transaction)
    - [x] [List of available currencies](https://changenow.io/api/docs/standard-flow/method-currencies)
    - [x] [List of available currencies for specific currency](https://changenow.io/api/docs/standard-flow/method-currencies-to)
    - [x] [Specific currency info](https://changenow.io/api/docs/standard-flow/method-currency-ticker)
    - [x] [List of all available pairs](https://changenow.io/api/docs/standard-flow/method-available-pairs)
    - [x] [Minimal exchange amount](https://changenow.io/api/docs/standard-flow/method-min-amount)
    - [x] [List of transactions](https://changenow.io/api/docs/standard-flow/method-transactions)
    - [x] [Transaction status](https://changenow.io/api/docs/standard-flow/method-transaction-status)
    - [ ] [Live transactions updates](https://changenow.io/api/docs/standard-flow/method-socket)
    - [ ] [Scam identification](https://changenow.io/api/docs/standard-flow/method-scam-check)
    
- Fixed rate flow
    - [ ] List of available fixed-rate markets
    - [ ] Estimated fixed-rate exchange amount
    - [ ] Create fixed-rate transaction
    - [ ] List of transactions
    - [ ] Live transactions updates`

## Contributing

* Fork the repo
* Add desired functionality
* Issue a Pull Request

Better testing with mocks needs to be written.

## 🐞 Bugs

If you find a problem with the lib, please submit an issue.