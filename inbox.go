package goinsta

import (
	"encoding/json"
	"fmt"
)

// InboxItem is any conversation message.
type InboxItem struct {
	ID            string `json:"item_id"`
	UserID        int64  `json:"user_id"`
	Timestamp     int64  `json:"timestamp"`
	ClientContext string `json:"client_context"`

	// Type there are a few types:
	// text, like, raven_media
	Type string `json:"item_type"`

	// Text is message text.
	Text string `json:"text"`

	// InboxItemLike is the heart that your girlfriend send to you.
	// (or in my case: the heart that my fans sends to me hehe)

	Like string `json:"like"`

	// Media is image or video
	Media struct {
		ID                   int64  `json:"id"`
		Images               Images `json:"image_versions2"`
		OriginalWidth        int    `json:"original_width"`
		OriginalHeight       int    `json:"original_height"`
		MediaType            int    `json:"media_type"`
		MediaID              int64  `json:"media_id"`
		PlaybackDurationSecs int    `json:"playback_duration_secs"`
		URLExpireAtSecs      int    `json:"url_expire_at_secs"`
		OrganicTrackingToken string `json:"organic_tracking_token"`
	}
}

// Inbox is the direct message inbox.
//
// Inbox contains Conversations. Each conversation has InboxItems.
// InboxItems are the message of the chat.
type Inbox struct {
	inst *Instagram
	err  error

	Conversations []Conversation `json:"threads"`

	HasNewer            bool   `json:"has_newer"` // TODO
	HasOlder            bool   `json:"has_older"`
	Cursor              string `json:"oldest_cursor"`
	UnseenCount         int    `json:"unseen_count"`
	UnseenCountTs       int64  `json:"unseen_count_ts"`
	BlendedInboxEnabled bool   `json:"blended_inbox_enabled"`
	// this fields are copied from response
	SeqID                int64 `json:"seq_id"`
	PendingRequestsTotal int   `json:"pending_requests_total"`
	SnapshotAtMs         int64 `json:"snapshot_at_ms"`
}

type inboxResp struct {
	Inbox                Inbox  `json:"inbox"`
	SeqID                int64  `json:"seq_id"`
	PendingRequestsTotal int    `json:"pending_requests_total"`
	SnapshotAtMs         int64  `json:"snapshot_at_ms"`
	Status               string `json:"status"`
}

func newInbox(inst *Instagram) *Inbox {
	return &Inbox{inst: inst}
}

func (inbox *Inbox) sync(pending bool, params map[string]string) error {
	endpoint := urlInbox
	if pending {
		endpoint = urlInboxPending
	}

	insta := inbox.inst
	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: endpoint,
			Query:    params,
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
			for i := range inbox.Conversations {
				inbox.Conversations[i].inst = insta
				inbox.Conversations[i].firstRun = true
			}
		}
	}
	return err
}

func (inbox *Inbox) next(pending bool, params map[string]string) bool {
	endpoint := urlInbox
	if pending {
		endpoint = urlInboxPending
	}
	if inbox.err != nil {
		return false
	}
	insta := inbox.inst
	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: endpoint,
			Query:    params,
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
			for i := range inbox.Conversations {
				inbox.Conversations[i].inst = insta
				inbox.Conversations[i].firstRun = true
			}
			if inbox.Cursor == "" || !inbox.HasOlder {
				inbox.err = ErrNoMore
			}
			return true
		}
	}
	inbox.err = err
	return false
}

// Sync updates inbox messages.
//
// See example: examples/inbox/sync.go
func (inbox *Inbox) Sync() error {
	return inbox.sync(false, map[string]string{
		"persistentBadging": "true",
		"use_unified_inbox": "true",
	})
}

// SyncPending updates inbox pending messages.
//
// See example: examples/inbox/sync.go
func (inbox *Inbox) SyncPending() error {
	return inbox.sync(true, map[string]string{})
}

// New initialises a new conversation with a user, for further messages you should use Conversation.Send
//
// See example: examples/inbox/newconversation.go
func (inbox *Inbox) New(user *User, text string) error {
	insta := inbox.inst
	to, err := prepareRecipients(user.ID)
	if err != nil {
		return err
	}

	data := insta.prepareDataQuery(
		map[string]interface{}{
			"recipient_users": to,
			"client_context":  generateUUID(),
			"thread_ids":      `["0"]`,
			"action":          "send_item",
			"text":            text,
		},
	)
	_, err = insta.sendRequest(
		&reqOptions{
			Connection: "keep-alive",
			Endpoint:   urlInboxSend,
			Query:      data,
			IsPost:     true,
		},
	)
	return err
}

// Reset sets inbox cursor at the beginning.
func (inbox *Inbox) Reset() {
	inbox.Cursor = ""
}

// Next allows pagination over messages.
//
// See example: examples/inbox/next.go
func (inbox *Inbox) Next() bool {
	return inbox.next(false, map[string]string{
		"persistentBadging": "true",
		"use_unified_inbox": "true",
		"cursor":            inbox.Cursor,
	})
}

// NextPending allows pagination over pending messages.
//
// See example: examples/inbox/next.go
func (inbox *Inbox) NextPending() bool {
	return inbox.next(true, map[string]string{
		"cursor": inbox.Cursor,
	})
}

// Conversation is the representation of an instagram already established conversation through direct messages.
type Conversation struct {
	inst     *Instagram
	err      error
	firstRun bool

	ID   string `json:"thread_id"`
	V2ID string `json:"thread_v2_id"`
	// Items can be of many types.
	Items                     []InboxItem `json:"items"`
	Title                     string      `json:"thread_title"`
	Users                     []User      `json:"users"`
	LeftUsers                 []User      `json:"left_users"`
	Pending                   bool        `json:"pending"`
	PendingScore              int64       `json:"pending_score"`
	ReshareReceiveCount       int         `json:"reshare_receive_count"`
	ReshareSendCount          int         `json:"reshare_send_count"`
	ViewerID                  int64       `json:"viewer_id"`
	ValuedRequest             bool        `json:"valued_request"`
	LastActivityAt            int64       `json:"last_activity_at"`
	Muted                     bool        `json:"muted"`
	IsPin                     bool        `json:"is_pin"`
	Named                     bool        `json:"named"`
	ThreadType                string      `json:"thread_type"`
	ExpiringMediaSendCount    int         `json:"expiring_media_send_count"`
	ExpiringMediaReceiveCount int         `json:"expiring_media_receive_count"`
	Inviter                   User        `json:"inviter"`
	HasOlder                  bool        `json:"has_older"`
	HasNewer                  bool        `json:"has_newer"`
	LastSeenAt                map[string]struct {
		Timestamp string `json:"timestamp"`
		ItemID    string `json:"item_id"`
	} `json:"last_seen_at"`
	NewestCursor      string `json:"newest_cursor"`
	OldestCursor      string `json:"oldest_cursor"`
	IsSpam            bool   `json:"is_spam"`
	LastPermanentItem Item   `json:"last_permanent_item"`
}

func (c Conversation) Error() error {
	return c.err
}

func (c Conversation) lastItemID() string {
	n := len(c.Items)
	if n == 0 {
		return ""
	}
	return c.Items[n-1].ID
}

// Like sends heart to the conversation
//
// See example: examples/media/likeAll.go
func (c *Conversation) Like() error {
	insta := c.inst
	to, err := prepareRecipients(c)
	if err != nil {
		return err
	}

	thread, err := json.Marshal([]string{c.ID})
	if err != nil {
		return err
	}

	data := insta.prepareDataQuery(
		map[string]interface{}{
			"recipient_users": to,
			"client_context":  generateUUID(),
			"thread_ids":      b2s(thread),
			"action":          "send_item",
		},
	)
	_, err = insta.sendRequest(
		&reqOptions{
			Connection: "keep-alive",
			Endpoint:   urlInboxSendLike,
			Query:      data,
			IsPost:     true,
		},
	)
	return err
}

// Send sends message in conversation
//
// See example: examples/inbox/sms.go
func (c *Conversation) Send(text string) error {
	insta := c.inst
	// I DON'T KNOW WHY BUT INSTAGRAM WANTS A DOUBLE SLICE OF INTS FOR ONE ID.
	to, err := prepareRecipients(c)
	if err != nil {
		return err
	}

	// I DONT KNOW WHY BUT INSTAGRAM WANTS SLICE OF STRINGS FOR ONE ID
	thread, err := json.Marshal([]string{c.ID})
	if err != nil {
		return err
	}

	data := insta.prepareDataQuery(
		map[string]interface{}{
			"recipient_users": to,
			"client_context":  generateUUID(),
			"thread_ids":      b2s(thread),
			"action":          "send_item",
			"text":            text,
		},
	)
	_, err = insta.sendRequest(
		&reqOptions{
			Connection: "keep-alive",
			Endpoint:   urlInboxSend,
			Query:      data,
			IsPost:     true,
		},
	)
	return err
}

// Write is like Send but being compatible with io.Writer.
func (c *Conversation) Write(b []byte) (int, error) {
	n := len(b)
	return n, c.Send(b2s(b))
}

// Next loads next set of private messages.
//
// See example: examples/inbox/conversation.go
func (c *Conversation) Next() bool {
	if c.err != nil {
		return false
	}
	if c.firstRun {
		c.firstRun = false
		return true
	}

	insta := c.inst
	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: fmt.Sprintf(urlInboxThread, c.ID),
			Query: map[string]string{
				"cursor":            c.lastItemID(),
				"direction":         "older", // go to upper
				"use_unified_inbox": "true",
			},
		},
	)
	if err == nil {
		resp := threadResp{}
		err = json.Unmarshal(body, &resp)
		if err == nil {
			*c = resp.Conversation
			c.inst = insta
			if !c.HasOlder {
				c.err = ErrNoMore
			}
			return true
		}
	}
	c.err = err
	return false
}
