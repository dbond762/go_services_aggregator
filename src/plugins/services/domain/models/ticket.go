package models

type Ticket struct {
	ID            int64
	UserServiceID int64
	Name          string
	Type          string
	Project       string
	Caption       string
	Status        string
	Priority      string
	Assignee      string
	Creator       string
}
