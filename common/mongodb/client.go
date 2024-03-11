package mongodb

var _ Client = (*customClient)(nil)

type Client interface {
	MsgApi
}
