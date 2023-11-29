package db

var _ Model = (*customModel)(nil)

type Model interface {
	dbUserModel
}
