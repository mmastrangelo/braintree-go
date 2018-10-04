package braintree

import (
	"encoding/xml"
	"time"
)

type ACHMandate struct {
	XMLName    xml.Name   `xml:"ach-mandate"`
	Text       string     `xml:"text"`
	AcceptedAt *time.Time `xml:"accepted-at"`
}
