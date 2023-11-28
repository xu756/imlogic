package db

var _ Model = (*customModel)(nil)

type Model interface {
	CreateTable() error
}
