package types

type MsgType int32

const (
	// 文本消息
	Text MsgType = 0
	// 图片消息
	Image MsgType = 1
	// 文件消息
	File MsgType = 2
	// 音频消息
	Audio MsgType = 3
	// 视频消息
	Video MsgType = 4
)
