package goinsta

// Comments allows user to interact with media (item) comments.
// You can Add or Delete by index or by user name.
type Comments struct {
	inst *Instagram

	media    Media
	endpoint string
	err      error

	Items                          []Comment `json:"comments"`
	CommentCount                   int       `json:"comment_count"`
	Caption                        Caption   `json:"caption"`
	CaptionIsEdited                bool      `json:"caption_is_edited"`
	HasMoreComments                bool      `json:"has_more_comments"`
	HasMoreHeadloadComments        bool      `json:"has_more_headload_comments"`
	MediaHeaderDisplay             string    `json:"media_header_display"`
	DisplayRealtimeTypingIndicator bool      `json:"display_realtime_typing_indicator"`
	NextID                         string    `json:"next_max_id"`
	LastID                         string    `json:"next_min_id"`
	Status                         string    `json:"status"`

	//PreviewComments                []Comment   `json:"preview_comments"`
}

func newComments(media Media) *Comments {
	c := &Comments{
		media: media,
	}
	return c
}

func (comments Comments) Error() error {
	return comments.err
}

// Next allows comment pagination.
//
// This function support concurrency methods to get comments using Last and Next ID
//
// New comments are stored in Comments.Items
func (comments *Comments) Next() bool {
	if comments.err != nil {
		return false
	}

	insta := comments.media.instagram()
	data, err := insta.prepareData(
		map[string]interface{}{
			"can_support_threading": true,
			"max_id":                comments.NextID,
			"min_id":                comments.LastID,
		},
	)
	if err != nil {
		comments.err = err
		return false
	}

	media := comments.media
	endpoint := comments.endpoint

	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: endpoint,
			Query:    generateSignature(data),
			IsPost:   true,
		},
	)
	if err == nil {
		c := Comments{}
		err = json.Unmarshal(body, &c)
		if err == nil {
			*comments = c
			comments.endpoint = endpoint
			comments.media = media
			if !comments.HasMoreComments || comments.NextID == "" {
				comments.err = ErrNoMore
			}
			return true
		}
	}
	comments.err = err
	return false
}

// Sync prepare Comments to receive comments.
// Use Next to receive comments.
func (comments *Comments) Sync() {
	media := comments.media
	endpoint := fmt.Sprintf(urlCommentSync, comments.media.ID())
	comments.media = media
	comments.endpoint = endpoint
	return
}

// Add push a comment in media.
//
// If parent media is a Story this function will send a private message
// replying the Instagram story. TODO
func (comments *Comments) Add(msg string) error {
	insta := comments.media.instagram()
	data, err := insta.prepareData(
		map[string]interface{}{
			"comment_text": msg,
		},
	)
	if err != nil {
		return err
	}

	_, err = insta.sendRequest(
		&reqOptions{
			Endpoint: fmt.Sprintf(urlCommentAdd, comments.media.ID()),
			Query:    generateSignature(data),
			IsPost:   true,
		},
	)
	return err
}

// Del deletes comment.
func (comments *Comments) Del(comment *Comment) error {
	insta := comments.media.instagram()

	data, err := insta.prepareData()
	if err != nil {
		return err
	}
	id := comment.getid()

	_, err = insta.sendRequest(
		&reqOptions{
			Endpoint: fmt.Sprintf(urlCommentDelete, comments.media.ID(), id),
			Query:    generateSignature(data),
			IsPost:   true,
		},
	)
	return err
}

// DelByID removes comment using id.
//
// See example: examples/media/commentsDelByID.go
func (comments *Comments) DelByID(id string) error {
	return comments.Del(&Comment{idstr: id})
}

// DelMine removes all of your comments limited by parsed parameter.
//
// If limit is <= 0 DelMine will delete all your comments.
//
// See example: examples/media/commentsDelMine.go
func (comments *Comments) DelMine(limit int) error {
	i := 0
	if limit <= 0 {
		i = limit - 1
	}
	comments.Sync()

	insta := comments.media.instagram()
floop:
	for comments.Next() {
		for _, c := range comments.Items {
			if c.UserID == insta.Account.ID || c.User.ID == insta.Account.ID {
				if i >= limit {
					break floop
				}
				comments.Del(&c)
				i++
			}
		}
	}
	if err := comments.Error(); err != nil && err != ErrNoMore {
		return err
	}
	return nil
}
