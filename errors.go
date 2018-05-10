package goinsta

import (
	"encoding/json"
	"fmt"
	"log"
)

// ErrNotFoundMsg is message text when the request responds with a 404 status code
// i.e a non existent user
var ErrNotFoundMsg = "The specified data wasn't found."

// ErrBadRequestMsg is message for when the request responds with a 400 status code
var ErrBadRequestMsg = "The account might be logged out. "

// IGAPIError ...
type IGAPIError struct {
	Message      string
	StatusCode   int64
	ResponseBody string
	ResponseData *IGAPIErrorResponse
}

func (err *IGAPIError) Error() string {
	if err.StatusCode > 0 {
		return fmt.Sprintf("%s [code=%d]", err.Message, err.StatusCode)
	}
	return err.Message
}

// NewIGAPIError ...
func NewIGAPIError(body string, code int64) *IGAPIError {
	var data IGAPIErrorResponse
	errmsg := body

	e := json.Unmarshal([]byte(body), &data)
	if e != nil {
		//unmarshal ng
		log.Printf("Error: %s, body=%s", e.Error(), body)

		switch code {
		case 400:
			//NOTE: IG API return 400 not only when logged out.
			//for excample, "The password you entered is incorrect. Please try again."
			//we are taking care of such errors by IGAPIErrorResponse.
			errmsg = ErrBadRequestMsg
		case 404:
			errmsg = ErrNotFoundMsg
		default:
			errmsg = fmt.Sprintf("Invalid status code %s", body)
		}
	} else {
		//unmarshal ok
		if data.Message != "" {
			// having data.message.
			errmsg = data.Message
		}
	}

	return &IGAPIError{
		Message:      errmsg,
		StatusCode:   code,
		ResponseBody: body,
		ResponseData: &data,
	}
}

// IGAPIErrorResponse ...
type IGAPIErrorResponse struct {
	Message            string `json:"message"`
	InvalidCredentials bool   `json:"invalid_credentials"`
	ErrorTitle         string `json:"error_title"`
	Buttons            []struct {
		Title  string `json:"title"`
		Action string `json:"action"`
	} `json:"buttons"`

	Challenge struct {
		URL               string `json:"url"`
		APIPath           string `json:"api_path"`
		HideWebviewHeader bool   `json:"hide_webview_header"`
		Lock              bool   `json:"lock"`
		Logout            bool   `json:"logout"`
		NativeFlow        bool   `json:"native_flow"`
	} `json:"challenge"`

	Status    string `json:"status"`
	ErrorType string `json:"error_type"`
}
