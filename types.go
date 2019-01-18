package changenow

// CurrencyInfo response
type CurrencyInfo struct {
	Name           string `json:"name"`
	Image          string `json:"image"`
	WarningMessage string `json:"warningMessage"`
	HasExternalID  bool   `json:"hasExternalId"`
	IsFiat         bool   `json:"isFiat"`
	IsAnonymous    bool   `json:"isAnonymous"`
	Wallets        struct {
		Primary []struct {
			Name     string `json:"name"`
			URL      string `json:"url"`
			ImageURL string `json:"imageUrl"`
		} `json:"primary"`
	} `json:"wallets"`
}

// MinAmount response
type MinAmount struct {
	Amount float64 `json:"minAmount"`
}

// EstimatedExchangeAmount response
type EstimatedExchangeAmount struct {
	EstimatedAmount          float64 `json:"estimatedAmount"`
	TransactionSpeedForecast string  `json:"transactionSpeedForecast"`
	WarningMessage           string  `json:"warningMessage"`
}

// Currency response
type Currency struct {
	Ticker            string `json:"ticker"`
	Image             string `json:"image"`
	HasExternalID     bool   `json:"hasExternalId"`
	IsFiat            bool   `json:"isFiat"`
	SupportsFixedRate bool   `json:"supportsFixedRate"`
	Featured          bool   `json:"featured"`
}

// TransactionRequest request
type TransactionRequest struct {
	From          string  `json:"from"`                    // from (Required): Ticker of a currency you want to send
	To            string  `json:"to"`                      // to (Required): Ticker of a currency you want to receive
	Address       string  `json:"address"`                 // address (Required): Address to receive a currency
	Amount        float64 `json:"amount"`                  // amount (Required): Amount you want to exchange
	ExtraID       string  `json:"extraId,omitempty"`       // extraId (Optional): Extra Id for currencies that require it
	RefundAddress string  `json:"refundAddress,omitempty"` // refundAddress (Optional): Refund address
}

// Transaction ...
type Transaction struct {
	// Base fields (create transaction)
	ID            string `json:"id"`
	PayinAddress  string `json:"payinAddress"`
	PayoutAddress string `json:"payoutAddress"`
	PayinExtraID  string `json:"payinExtraId"`
	FromCurrency  string `json:"fromCurrency"`
	ToCurrency    string `json:"toCurrency"`
	RefundAddress string `json:"refundAddress"`
}

// TransactionDetails ...
type TransactionDetails struct {
	ID            string  `json:"id"`
	PayinAddress  string  `json:"payinAddress"`
	PayoutAddress string  `json:"payoutAddress"`
	PayinExtraID  string  `json:"payinExtraId"`
	FromCurrency  string  `json:"fromCurrency"`
	ToCurrency    string  `json:"toCurrency"`
	Status        string  `json:"status"`
	Hash          string  `json:"hash"`
	PayinHash     string  `json:"payinHash"`
	PayoutHash    string  `json:"payoutHash"`
	AmountSend    float64 `json:"amountSend"`
	AmountReceive float64 `json:"amountReceive"`
	NetworkFee    float64 `json:"networkFee"`
	UpdatedAt     string  `json:"updatedAt"`
}

// // Status enum (transactions)
// type Status string

// Statuses
const (
	Waiting    string = "waiting"
	Confirming        = "confirming"
	Sending           = "sending"
	Finished          = "finished"
	Expired           = "expired"
)
