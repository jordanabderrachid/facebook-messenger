package messenger

import (
	"encoding/json"
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

type sendable interface {
	toBody() (io.Reader, error)
}

// NewMessenger comment
func NewMessenger(token string) *Messenger {
	return &Messenger{
		Token:  token,
		client: &http.Client{},
	}
}

// SendTextMessage comment
func (m *Messenger) SendTextMessage(recipientID, text string) (Response, error) {
	err := m.checkToken()
	if err != nil {
		return Response{}, err
	}

	tm := newTextMessage(recipientID, text)
	return m.send(tm)
}

// SendImageAttachment comment
func (m *Messenger) SendImageAttachment(recipientID, imageURL string) (Response, error) {
	err := m.checkToken()
	if err != nil {
		return Response{}, err
	}

	ia := newImageAttachment(recipientID, imageURL)
	return m.send(ia)
}

// SendAudioAttachment comment
func (m *Messenger) SendAudioAttachment(recipientID, audioURL string) (Response, error) {
	err := m.checkToken()
	if err != nil {
		return Response{}, err
	}

	aa := newAudioAttachment(recipientID, audioURL)
	return m.send(aa)
}

// SendVideoAttachment comment
func (m *Messenger) SendVideoAttachment(recipientID, videoURL string) (Response, error) {
	err := m.checkToken()
	if err != nil {
		return Response{}, err
	}

	va := newVideoAttachment(recipientID, videoURL)
	return m.send(va)
}

// SendFileAttachment comment
func (m *Messenger) SendFileAttachment(recipientID, fileURL string) (Response, error) {
	err := m.checkToken()
	if err != nil {
		return Response{}, err
	}

	fa := newFileAttachment(recipientID, fileURL)
	return m.send(fa)
}

func (m *Messenger) send(data sendable) (Response, error) {
	body, err := data.toBody()
	if err != nil {
		return Response{}, err
	}

	resp, err := m.apiCall(body)
	defer resp.Body.Close()
	if err != nil {
		return Response{}, err
	}

	return handleReponse(resp)
}

func (m *Messenger) getURL() string {
	return fmt.Sprintf(urlFormat, m.Token)
}

func (m *Messenger) apiCall(body io.Reader) (*http.Response, error) {
	return m.client.Post(m.getURL(), contentType, body)
}

func (m *Messenger) checkToken() error {
	if !isTokenSet(m.Token) {
		return errors.New("token must be set")
	}

	return nil
}

func handleReponse(resp *http.Response) (Response, error) {
	statusCode := resp.StatusCode
	switch resp.StatusCode {
	case 200:
		// OK
		return parseOKResponse(resp.Body)
	case 400:
		// Bad request
		return Response{}, parseBadRequestResponse(resp.Body)
	case 500, 502, 503, 504:
		// Internal server error, Bad gateway , Service unavailable and Gateway timeout
		return Response{}, fmt.Errorf("Error on remote server, status code: %d", statusCode)
	default:
		return Response{}, fmt.Errorf("Unhandled response status code: %d", statusCode)
	}
}

func parseOKResponse(body io.ReadCloser) (Response, error) {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return Response{}, err
	}

	r := Response{}
	err = json.Unmarshal(data, &r)
	return r, err
}

func parseBadRequestResponse(body io.ReadCloser) error {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	brr := badRequestResponse{}
	err = json.Unmarshal(data, &brr)
	if err != nil {
		return err
	}
	return brr.toError()
}

func isTokenSet(token string) bool {
	return token != ""
}
