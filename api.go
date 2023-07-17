package emessage

import (
	"fmt"
)

type APIResponseData interface {
	apiResponseData()
}

type APIRequest interface {
	apiRequest()
}

type APIResponse struct {
	ApiStatusCode int             `json:"apiStatusCode"`
	Data          APIResponseData `json:"data"`
}

type ErrStatusCode struct {
	code int
}

func (s *ErrStatusCode) StatusCode() int {
	return s.code
}

func (s *ErrStatusCode) Error() string {
	return fmt.Sprintf("status code error, code %d)", s.code)
}

const (
	// Default credentials extracted from webinterface (yes they are constant and are still sent)
	DefaultUsername     = "sendRUF"
	DefaultPasswordHash = "38ec3d06496c26288290e5cd129c9cda3be82c4ce327618c573f70fae74c7370"
)

// SendMessage is a simple wrapper function that sends a message to one identifier
func SendMessage(identifier, messageText string) (*SendRufResponse, error) {
	login2 := &Login2Request{
		Username:     DefaultUsername,
		PasswordHash: DefaultPasswordHash,
	}

	l2res, err := login2.Send()
	if err != nil {
		return nil, err
	}

	sendRuf := &SendRufRequest{
		JWT: l2res.JWT,

		Identifier:  identifier,
		MessageText: messageText,
	}

	srres, err := sendRuf.Send()
	return srres, err
}
