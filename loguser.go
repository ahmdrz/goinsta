package goinsta

type LoggedUser struct {
	Username string
	Password string
	Feed     *UserFeed
	Device   DeviceInfo
}
