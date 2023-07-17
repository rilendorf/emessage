package emessage

import (
	"gopkg.in/resty.v1"

	"encoding/json"
)

type SendRufRequest struct {
	JWT string `json:"-"`

	Identifier  string `json:"identifier"`
	MessageText string `json:"messageText"`
}

func (*SendRufRequest) apiRequest() {}

type SendRufResponse struct {
	Status     string            `json:"status"`
	TrackingID string            `json:"trackingId"`
	Recipients []RecipientStatus `json:"recipients"`
}

type RecipientStatus struct {
	Identifier        string `json:"identifier"`
	StatusSendMessage string `json:"statusSendMessage"`
}

func (*SendRufResponse) apiResponseData() {}

func (s *SendRufRequest) Send() (res *SendRufResponse, err error) {
	body, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+s.JWT).
		SetBody(string(body)).
		Post("https://api.emessage.de/rs/ruf/send")

	r := &APIResponse{
		Data: new(SendRufResponse),
	}

	err = json.Unmarshal(resp.Body(), r)
	if err != nil {
		return nil, err
	}

	if r.ApiStatusCode != 200 {
		return nil, &ErrStatusCode{r.ApiStatusCode}
	}

	return r.Data.(*SendRufResponse), nil
}
