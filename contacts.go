package goinsta

import (
	"encoding/json"
	"strconv"
)

type Contacts struct {
	inst *Instagram
}

type Contact struct {
	Numbers []string `json:"phone_numbers"`
	Emails  []string `json:"email_addresses"`
	Name    string   `json:"first_name"`
}

type SyncAnswer struct {
	Users []struct {
		Pk                         int64  `json:"pk"`
		Username                   string `json:"username"`
		FullName                   string `json:"full_name"`
		IsPrivate                  bool   `json:"is_private"`
		ProfilePicURL              string `json:"profile_pic_url"`
		ProfilePicID               string `json:"profile_pic_id"`
		IsVerified                 bool   `json:"is_verified"`
		HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
		ReelAutoArchive            string `json:"reel_auto_archive"`
		AddressbookName            string `json:"addressbook_name"`
	} `json:"users"`
	Warning string `json:"warning"`
	Status  string `json:"status"`
}

func newContacts(inst *Instagram) *Contacts {
	return &Contacts{inst: inst}
}

func (c *Contacts) SyncContacts(contacts *[]Contact) (*SyncAnswer, error) {
	acquireContacts := &reqOptions{
		Endpoint: "address_book/acquire_owner_contacts/",
		IsPost:   true,
		Login:    true,
		UseV2:    false,
		Query: map[string]string{
			"phone_id": c.inst.pid,
			"me":       `{"phone_numbers":[],"email_addresses":[]}`,
		},
	}
	body, err := c.inst.sendRequest(acquireContacts)
	if err != nil {
		return nil, err
	}

	byteContacts, err := json.Marshal(contacts)
	if err != nil {
		return nil, err
	}

	syncContacts := &reqOptions{
		Endpoint: `address_book/link/`,
		IsPost:   true,
		Login:    true,
		UseV2:    false,
		Query: map[string]string{
			"_uuid":      c.inst.uuid,
			"_csrftoken": c.inst.token,
			"contacts":   string(byteContacts),
		},
	}

	body, err = c.inst.sendRequest(syncContacts)
	if err != nil {
		return nil, err
	}

	answ := &SyncAnswer{}
	json.Unmarshal(body, answ)
	return answ, nil
}

func (c *Contacts) UnlinkContacts() error {
	toSign := map[string]string{
		"_csrftoken": c.inst.token,
		"_uid":       strconv.Itoa(int(c.inst.Account.ID)),
		"_uuid":      c.inst.uuid,
	}

	bytesS, _ := json.Marshal(toSign)

	unlinkBody := &reqOptions{
		Endpoint: "address_book/unlink/",
		IsPost:   true,
		Login:    true,
		UseV2:    false,
		Query:    generateSignature(string(bytesS)),
	}

	_, err := c.inst.sendRequest(unlinkBody)
	if err != nil {
		return err
	}
	return nil
}
