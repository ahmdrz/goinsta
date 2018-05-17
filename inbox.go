package goinsta

import (
	"encoding/json"
)

// Inbox is the direct message inbox.
type Inbox struct {
	inst *Instagram

	Conversations []Conversation `json:"threads"`

	HasOlder            bool  `json:"has_older"`
	UnseenCount         int   `json:"unseen_count"`
	UnseenCountTs       int64 `json:"unseen_count_ts"`
	BlendedInboxEnabled bool  `json:"blended_inbox_enabled"`
	// this fields are copied from response
	SeqID                int   `json:"seq_id"`
	PendingRequestsTotal int   `json:"pending_requests_total"`
	SnapshotAtMs         int64 `json:"snapshot_at_ms"`
}

type inboxResp struct {
	Inbox                Inbox  `json:"inbox"`
	SeqID                int    `json:"seq_id"`
	PendingRequestsTotal int    `json:"pending_requests_total"`
	SnapshotAtMs         int64  `json:"snapshot_at_ms"`
	Status               string `json:"status"`
}

func newInbox(inst *Instagram) *Inbox {
	return &Inbox{inst: inst}
}

// Sync updates inbox messages.
func (inbox *Inbox) Sync() error {
	insta := inbox.inst
	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: urlInbox,
			Query: map[string]string{
				"persistentBadging": "true",
				"use_unified_inbox": "true",
			},
		},
	)
	if err == nil {
		resp := inboxResp{}
		err = json.Unmarshal(body, &resp)
		if err == nil {
			*inbox = resp.Inbox
			inbox.inst = insta
			inbox.SeqID = resp.Inbox.SeqID
			inbox.PendingRequestsTotal = resp.Inbox.PendingRequestsTotal
			inbox.SnapshotAtMs = resp.Inbox.SnapshotAtMs
		}
	}
	return err
}
