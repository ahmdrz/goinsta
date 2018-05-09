package goinsta

import (
	"fmt"
	"net/http"
)

// Instagram represent the main API handler
//
// ...
type Instagram struct {
	user string
	pass string
	// device id
	dID string
	// uuid
	uuid string
	// rankToken
	rankToken string
	// token
	token string
	// phone id
	pid string

	// Instagram objects

	// Users is the user interaction
	Users *Users
	// Account stores all personal data of the user and his/her options.
	Account *Account
	// Search performs searching of multiple things (users, locations...)
	Search *Search

	logged bool

	c *http.Client
}

// School is void structure (yet).
type School struct {
}

// PicURLInfo repre
type PicURLInfo struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

type instaError struct {
	Message   string `json:"message"`
	Status    string `json:"status"`
	ErrorType string `json:"error_type"`
}

func instaToErr(ierr instaError) error {
	return fmt.Errorf("%s: %s | %s", ierr.Status, ierr.Message, ierr.ErrorType)
}

type Nametag struct {
	Mode          int    `json:"mode"`
	Gradient      int    `json:"gradient"`
	Emoji         string `json:"emoji"`
	SelfieSticker int    `json:"selfie_sticker"`
}
