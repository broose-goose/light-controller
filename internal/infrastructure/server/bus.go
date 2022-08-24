package server

type Bus interface {
	AuthorizeLogin() (string, error)
}
