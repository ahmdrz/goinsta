package endpoint

var (
	FetchHeader = 0
	Login       = 1

	endpoints = map[int]string{
		FetchHeader: "si/fetch_headers/",
		Login:       "acccounts/login",
	}
)
