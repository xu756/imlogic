package types

type ActionType string

const (
	// 发送
	Send ActionType = "Send"
	// 接收
	Receive ActionType = "Receive"
	// 撤回
	Recall ActionType = "Recall"
)
