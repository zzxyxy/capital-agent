package connection

type Session_reply_account_info struct {
	Balance    int `json:"balance"`
	Deposit    int `json:"deposit"`
	Profitloss int `json:"profitLoss"`
	Available  int `json:"available"`
}

type Session_reply_account struct {
	Accountid   string `json:"accounId"`
	Accountname string `json:"accountName"`
	Preferred   bool   `json:"preferred"`
	AccountType string `json:"accountType"`
}

type Session_reply struct {
	Accounttype           string                     `json:"accountType"`
	Accountinfo           Session_reply_account_info `json:"accountInfo"`
	Currencyisocode       string                     `json:"currencyIsoCode"`
	Currencysymbol        string                     `json:"currencySymbol"`
	Currentaccountid      string                     `json:"currentAccountId"`
	Streaminghost         string                     `json:"streamingHost"`
	Accounts              []Session_reply_account    `json:"accounts"`
	Clientid              string                     `json:"clientId"`
	TimezoneOffset        int                        `json:"timezoneOffset"`
	HasActivedemoaccounts bool                       `json:"hasActiveDemoAccounts"`
	HasActiveliveaccounts bool                       `json:"hasActiveLiveAccounts"`
	Trailingstopenabled   bool                       `json:"trailingStopsEnabled"`
}
