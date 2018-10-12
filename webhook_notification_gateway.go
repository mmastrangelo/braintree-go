package braintree

import (
	"encoding/base64"
	"encoding/xml"
	"net/http"

	"github.com/mmastrangelo/braintree-go/xmlnil"
)

type WebhookNotificationGateway struct {
	*Braintree
	apiKey apiKey
}

func (w *WebhookNotificationGateway) ParseRequest(r *http.Request) (*WebhookNotification, error) {
	signature := r.PostFormValue("bt_signature")
	payload := r.PostFormValue("bt_payload")
	return w.Parse(signature, payload)
}

func (w *WebhookNotificationGateway) VerifySignature(signature, payload string) (bool, error) {
	hmacer := newHmacer(w.apiKey.publicKey, w.apiKey.privateKey)
	return hmacer.verifySignature(signature, payload)
}

func (w *WebhookNotificationGateway) Parse(signature, payload string) (*WebhookNotification, error) {
	if verified, err := w.VerifySignature(signature, payload); err != nil {
		return nil, err
	} else if !verified {
		return nil, SignatureError{}
	}

	xmlNotification, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		return nil, err
	}

	strippedBuf, err := xmlnil.StripNilElements(xmlNotification)
	if err != nil {
		return nil, err
	}

	var n WebhookNotification
	err = xml.Unmarshal(strippedBuf, &n)
	if err != nil {
		return nil, err
	}
	return &n, nil
}

func (w *WebhookNotificationGateway) Verify(challenge string) (string, error) {
	hmacer := newHmacer(w.apiKey.publicKey, w.apiKey.privateKey)
	digest, err := hmacer.hmac(challenge)
	if err != nil {
		return ``, err
	}
	return hmacer.publicKey + `|` + digest, nil
}
