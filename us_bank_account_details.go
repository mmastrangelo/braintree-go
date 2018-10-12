package braintree

type USBankAccountDetails struct {
	RoutingNumber     string      `xml:"routing-number,omitempty"`
	Last4             string      `xml:"last-4,omitempty"`
	AccountType       string      `xml:"account-type,omitempty"`
	AccountHolderName string      `xml:"account-holder-name,omitempty"`
	Token             string      `xml:"token,omitempty"`
	ImageURL          string      `xml:"image-url,omitempty"`
	BankName          string      `xml:"bank-name,omitempty"`
	ACHMandate        *ACHMandate `xml:"ach-mandate,omitempty"`
}
