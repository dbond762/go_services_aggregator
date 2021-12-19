package models

type Service struct {
	Ident         string
	Type          string
	UserID        int64
	UserServiceID int64
	Credentials   map[string]string
}
