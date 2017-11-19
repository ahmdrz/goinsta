package goinsta

var (
	APIUrl       = "https://i.instagram.com/api/v1/"
	APIUserAgent = ""
	APISigKey    = ""
	APISigKeyV   = "4"
)

type Instagram struct {
	HttpInfo *Conn
	User     *LoggedUser
	Stories  *Stories
	Messages *Messages
}
