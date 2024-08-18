namespace go im
include "base.thrift"



// 发送消息给单个用户
struct SendMsgToOneReq {
  1: string link_id,
  2: base.Message message
}

// 发送消息给多个用户
struct SendMsgToGroupReq {
  1: list<string> link_ids,
  2: base.Message message
}

// 发送消息给多个用户失败的用户
struct SendMsgToGroupRes {
  1: list<string> link_ids
}

// 返回消息
struct MessageRes {
  1: bool success,
  2: base.Message message
}

// 处理websocket消息
service ImHandler {
  // 元消息
  MessageRes MetaMessage(1: base.MetaMsg metaMsg),
  // 处理私聊消息
  MessageRes HandlerPrivateMessage(1: base.Message message),
  // 处理群聊消息
  MessageRes HandlerGroupMessage(1: base.Message message)
}

// 发送消息
service ImServer {
  // 发送消息给单个用户
  MessageRes SendMsgToOne(1: SendMsgToOneReq sendMsgToOneReq),
  // 发送消息给多个用户
  SendMsgToGroupRes SendMsgToGroup(1: SendMsgToGroupReq sendMsgToGroupReq),
  // 发送消息给所有用户
  MessageRes SendMsgToAll(1: base.Message message)
}
