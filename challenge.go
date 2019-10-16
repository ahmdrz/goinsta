package goinsta

import (
	"encoding/json"
	"strings"
)

type ChallengeStepData struct {
	Choice           string      `json:"choice"`
	FbAccessToken    string      `json:"fb_access_token"`
	BigBlueToken     string      `json:"big_blue_token"`
	GoogleOauthToken string      `json:"google_oauth_token"`
	Email            string      `json:"email"`
	SecurityCode     string      `json:"security_code"`
	ResendDelay      interface{} `json:"resend_delay"`
	ContactPoint     string      `json:"contact_point"`
	FormType         string      `json:"form_type"`
}
type Challenge struct {
	insta        *Instagram
	StepName     string            `json:"step_name"`
	StepData     ChallengeStepData `json:"step_data"`
	LoggedInUser *Account          `json:"logged_in_user,omitempty"`
	UserID       int64             `json:"user_id"`
	NonceCode    string            `json:"nonce_code"`
	Action       string            `json:"action"`
	Status       string            `json:"status"`
}

type challengeResp struct {
	*Challenge
}

func newChallenge(insta *Instagram) *Challenge {
	time := &Challenge{
		insta: insta,
	}
	return time
}

// updateState updates current data from challenge url
func (challenge *Challenge) updateState() error {
	insta := challenge.insta

	data, err := insta.prepareData(map[string]interface{}{
		"guid":      insta.uuid,
		"device_id": insta.dID,
	})
	if err != nil {
		return err
	}

	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: challenge.insta.challengeURL,
			Query:    generateSignature(data),
		},
	)
	if err == nil {
		resp := challengeResp{}
		err = json.Unmarshal(body, &resp)
		if err == nil {
			*challenge = *resp.Challenge
			challenge.insta = insta
		}
	}
	return err
}

// selectVerifyMethod selects a way and verify it (Phone number = 0, email = 1)
func (challenge *Challenge) selectVerifyMethod(choice string, isReplay ...bool) error {
	insta := challenge.insta

	url := challenge.insta.challengeURL
	if len(isReplay) > 0 && isReplay[0] {
		url = strings.Replace(url, "/challenge/", "/challenge/replay/", -1)
	}

	data, err := insta.prepareData(map[string]interface{}{
		"choice":    choice,
		"guid":      insta.uuid,
		"device_id": insta.dID,
	})
	if err != nil {
		return err
	}

	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: url,
			Query:    generateSignature(data),
			IsPost:   true,
		},
	)
	if err == nil {
		resp := challengeResp{}
		err = json.Unmarshal(body, &resp)
		if err == nil {
			*challenge = *resp.Challenge
			challenge.insta = insta
		}
	}
	return err
}

// sendSecurityCode sends the code received in the message
func (challenge *Challenge) SendSecurityCode(code string) error {
	insta := challenge.insta
	url := challenge.insta.challengeURL

	data, err := insta.prepareData(map[string]interface{}{
		"security_code": code,
		"guid":          insta.uuid,
		"device_id":     insta.dID,
	})
	if err != nil {
		return err
	}

	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: url,
			Query:    generateSignature(data),
			IsPost:   true,
		},
	)
	if err == nil {
		resp := challengeResp{}
		err = json.Unmarshal(body, &resp)
		if err == nil {
			*challenge = *resp.Challenge
			challenge.insta = insta
		}
	}
	return err
}

// deltaLoginReview process with choice (It was me = 0, It wasn't me = 1)
func (challenge *Challenge) deltaLoginReview() error {
	return challenge.selectVerifyMethod("0")
}

func (challenge *Challenge) Process(apiURL string) error {
	challenge.insta.challengeURL = apiURL[1:]

	if err := challenge.updateState(); err != nil {
		return err
	}

	switch challenge.StepName {
	case "select_verify_method":
		return challenge.selectVerifyMethod(challenge.StepData.Choice)
	case "delta_login_review":
		return challenge.deltaLoginReview()
	}

	return ErrChallengeProcess{StepName: challenge.StepName}
}
