package goinsta

type Search struct {
	inst *Instagram
}

type SearchResult struct {
	HasMore    bool   `json:"has_more"`
	Status     string `json:"status"`
	NumResults int    `json:"num_results"`
	RankToken  string `json:"rank_token"`

	// User search results
	Users []User `json:"users"`

	// Tag search results
	Tags []struct {
		Name       string `json:"name"`
		MediaCount int    `json:"media_count"`
		ID         int64  `json:"id"`
	} `json:"results"`

	// Location search result
	RequestID string `json:"request_id"`
	Venues    []struct {
		ExternalIDSource string  `json:"external_id_source"`
		ExternalID       string  `json:"external_id"`
		Lat              float64 `json:"lat"`
		Lng              float64 `json:"lng"`
		Address          string  `json:"address"`
		Name             string  `json:"name"`
	} `json:"venues"`

	// Facebook
	// TODO
}

// newSearch creates new Search structure
func newSearch(inst *Instagram) *Search {
	search := &Search{
		inst: inst,
	}
	return search
}

// User search by username
func (search *Search) User(user string) (*SearchResult, error) {
	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: urlSearchUser,
			Query: map[string]string{
				"ig_sig_key_version": goInstaSigKeyVersion,
				"is_typeahead":       "true",
				"query":              user,
				"rank_token":         search.insta.rankToken,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	res := &SearchResult{}
	err = json.Unmarshal(body, res)
	return res, err
}

// Tags search by tag
func (search *Search) Tags(tag string) (*SearchResult, error) {
	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: urlSearchTag,
			Query: map[string]string{
				"is_typeahead": "true",
				"rank_token":   search.insta.rankToken,
				"q":            tag,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	res := &SearchResult{}
	err = json.Unmarshal(body, res)
	return res, err
}

// Location search by location.
//
// Lat and Lng (Latitude & Longitude) cannot be ""
func (search *Search) Location(lat, lng, search string) (*SearchResult, error) {
	q := map[string]string{
		"rank_token":     search.inst.rankToken,
		"latitude":       lat,
		"longitude":      lng,
		"ranked_content": "true",
	}

	if search != "" {
		q["search_query"] = search
	} else {
		q["timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	}

	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: urlSearchLocation,
			Query:    q,
		},
	)
	if err != nil {
		return nil, err
	}

	res := &SearchResult{}
	err = json.Unmarshal(body, res)
	return res, err
}

func (search *Search) Facebook(user string) (*SearchResult, error) {
	// TODO
	return nil, nil
}
