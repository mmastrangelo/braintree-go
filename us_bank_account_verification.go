package braintree

import (
	"encoding/xml"
	"time"
)

type USBankAccountVerification struct {
	XMLName                  xml.Name                        `xml:"us-bank-account-verification"`
	ID                       string                          `xml:"id"`
	Status                   USBankAccountStatus             `xml:"status"`
	VerificationMethod       USBankAccountVerificationMethod `xml:"verification-method"`
	ProcessorResponseCode    string                          `xml:"processor-response-code"`
	ProcessorResponseText    string                          `xml:"processor-response-text"`
	VerificationDeterminedAt *time.Time                      `xml:"verification-determined-at"`
	CreatedAt                *time.Time                      `xml:"created-at"`
	GatewayRejectionReason   GatewayRejectionReason          `xml:"gateway-rejection-reason"`
	USBankAccount            *USBankAccount                  `xml:"us-bank-account"`
}

type USBankAccountVerifications struct {
	USBankAccountVerification []*USBankAccountVerification `xml:"us-bank-account-verification"`
}

type USBankAccountVerificationMethod string

const (
	VerificationMethodTokenizedCheck   USBankAccountVerificationMethod = "tokenized_check"
	VerificationMethodNetworkCheck     USBankAccountVerificationMethod = "network_check"
	VerificationMethodIndependentCheck USBankAccountVerificationMethod = "independent_check"
	VerificationMethodUnrecognized     USBankAccountVerificationMethod = "unrecognized"
	VerificationMethodMicroTransfers   USBankAccountVerificationMethod = "micro_transfers"
)

type USBankAccountStatus string

const (
	USBankAccountStatusFailed            USBankAccountStatus = "failed"
	USBankAccountStatusGatewayRejected   USBankAccountStatus = "gateway_rejected"
	USBankAccountStatusProcessorDeclined USBankAccountStatus = "processor_declined"
	USBankAccountStatusUnrecognized      USBankAccountStatus = "unrecognized"
	USBankAccountStatusVerified          USBankAccountStatus = "verified"
	USBankAccountStatusPending           USBankAccountStatus = "pending"
)
