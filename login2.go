package emessage

import (
	"gopkg.in/resty.v1"

	"encoding/json"
)

type Login2Request struct {
	Username     string `json:"username"`
	PasswordHash string `json:"passwordHash"`
}

func (*Login2Request) apiRequest() {}

type Login2Response struct {
	Username string `json:"username"`
	JWT      string `json:"jwt"`
}

func (*Login2Response) apiResponseData() {}

func (s *Login2Request) Send() (res *Login2Response, err error) {
	body, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(string(body)).
		Post("https://api.emessage.de/auth/login/2")

	r := &APIResponse{
		Data: new(Login2Response),
	}

	err = json.Unmarshal(resp.Body(), r)
	if err != nil {
		return nil, err
	}

	if r.ApiStatusCode != 200 {
		return nil, &ErrStatusCode{r.ApiStatusCode}
	}

	return r.Data.(*Login2Response), nil
}
