package models

type Projects struct {
	ID           int64
	Name         string
	DueDate      string
	Owner        int
	Participants map[int]string
}
