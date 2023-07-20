package models

type Organization struct {
	ID      int64
	BIN     string
	Name    string
	Address string
	Workers int
}
