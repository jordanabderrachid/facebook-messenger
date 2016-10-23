package messenger

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	contentType = "application/json"
	urlFormat   = "https://graph.facebook.com/v2.6/me/messages?access_token=%s"
)

// Messenger comment
type Messenger struct {
	Token  string
	client *http.Client
}

// NewMessenger comment
func NewMessenger(token string) *Messenger {
	return &Messenger{
		Token:  token,
		client: &http.Client{},
	}
}

// SendTextMessage comment
func (m *Messenger) SendTextMessage(recipientID, text string) (string, error) {
	err := m.checkToken()
	if err != nil {
		return "", err
	}

	tm := newTextMessage(recipientID, text)
	body, err := tm.toBody()
	if err != nil {
		return "", err
	}

	resp, err := m.apiCall(body)
	if err != nil {
		return "", err
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	return string(responseBody), err
}

func (m *Messenger) checkToken() error {
	if !isTokenSet(m.Token) {
		return errors.New("token must be set")
	}

	return nil
}

func isTokenSet(token string) bool {
	return token != ""
}

func (m *Messenger) getURL() string {
	return fmt.Sprintf(urlFormat, m.Token)
}

func (m *Messenger) apiCall(data io.Reader) (*http.Response, error) {
	return m.client.Post(m.getURL(), contentType, data)
}
