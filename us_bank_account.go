package braintree

import (
	"encoding/xml"
)

type USBankAccount struct {
	XMLName                    xml.Name                    `xml:"us-bank-account"`
	RoutingNumber              string                      `xml:"routing-number"`
	Last4                      string                      `xml:"last-4"`
	AccountType                string                      `xml:"account-type"`
	AccountHolderName          string                      `xml:"account-holder-name"`
	Token                      string                      `xml:"token"`
	ImageURL                   string                      `xml:"image-url"`
	BankName                   string                      `xml:"bank-name"`
	Subscriptions              *Subscriptions              `xml:"subscriptions"`
	CustomerId                 string                      `xml:"customer-id"`
	Default                    bool                        `xml:"default"`
	ACHMandate                 *ACHMandate                 `xml:"ach-mandate"`
	USBankAccountVerifications *USBankAccountVerifications `xml:"verifications"`

	Verified bool `xml:"verified"`
}

type USBankAccounts struct {
	USBankAccount []*USBankAccount `xml:"us-bank-account"`
}

func (u *USBankAccounts) PaymentMethods() []PaymentMethod {
	if u == nil {
		return nil
	}
	var paymentMethods []PaymentMethod
	for _, u := range u.USBankAccount {
		paymentMethods = append(paymentMethods, u)
	}
	return paymentMethods
}

func (u *USBankAccount) GetCustomerId() string {
	return u.CustomerId
}

func (u *USBankAccount) GetToken() string {
	return u.Token
}

func (u *USBankAccount) IsDefault() bool {
	return u.Default
}

func (u *USBankAccount) GetImageURL() string {
	return u.ImageURL
}

// AllSubscriptions returns all subscriptions for this paypal account, or nil if none present.
func (u *USBankAccount) AllSubscriptions() []*Subscription {
	if u.Subscriptions != nil {
		subs := u.Subscriptions.Subscription
		if len(subs) > 0 {
			a := make([]*Subscription, 0, len(subs))
			for _, s := range subs {
				a = append(a, s)
			}
			return a
		}
	}
	return nil
}
