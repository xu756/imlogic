namespace go base

// 聊天类型
enum ChatType {
  // 私聊 （处理私聊消息）
  PrivateChat = 0,
  // 群聊 （处理群聊消息）
  GroupChat = 1,
  // 系统消息 （接收）
  SystemMessage = 2,
  // 系统通知 （接收）
  SystemNotice = 3
}

// 消息类型
enum MsgType {
  // 事件
  Event = 0
  /**文字*/
  Text = 1,
  // 图片
  Image = 2,
  // 文件
  File = 3,
  // 语音
  Audio = 4,
  // 视频
  Video = 5
}

// 事件类型
enum EventType {
  // 消息已读
  Read = 0,
  // 撤回消息
  Recall = 1,
  // 消息发送失败
  SendFail = 2
  // 不是好友
  NotFriend = 3
  // 不在群
  NotInGroup = 4
  // 群加入消息
  GroupJoin = 5,
  // 群聊退群
  GroupQuit = 6,

    //  // 群加入消息
    //  GroupJoin = 3
    //  // 群聊退群
    //  GroupQuit = 4
    //  // 群聊踢人
    //  GroupKick = 5
    //  // 群聊解散
    //  GroupDismiss = 6
    //  // 群聊禁言
    //  GroupBan = 7
    //  // 群聊解除禁言
    //  GroupUnBan = 8

}

// 媒体类型
struct MediaType {
  1: string uid,
  2: string url
}
// 事件 (撤回、消息发送失败、消息已读)
struct Event{
    1:EventType event_type (go.tag = "json:\"event_type,omitempty\""),
    2:i64 user_id (go.tag = "json:\"user_id,omitempty\""),
}

// 消息
struct Message {
  1: string msg_id (go.tag = "json:\"msg_id,omitempty\""),
  2: i64 timestamp ,
  3: ChatType chat_type,
  4: i64 sender (go.tag = "json:\"sender,omitempty\""),
  5: i64 receiver (go.tag = "json:\"receiver,omitempty\""),
  6: string content (go.tag = "json:\"content,omitempty\""),
  7: MsgType msg_type,
  8: list<MediaType> media (go.tag = "json:\"media,omitempty\""),
  9: Event event (go.tag = "json:\"event,omitempty\""),
}


// websocket状态
enum WsStatus {
  Connect = 0,
  Disconnect = 1,
  Heartbeat = 2
}

// 元消息
struct MetaMsg {
  1: string link_id,
  2: i64 user_id,
  3: WsStatus status,
  4: string host_name,
  5: string device
}

// 聊天列表
struct ChatList{
    1:string chat_id,
    2:ChatType chat_type,
    4:i64 with_id,
    5:Message last_msg,
    6:i64 timestamp
}





// 用户信息
struct UserInfo {
  1: i64 id,         // 用户ID
  2: string username, // 用户名
  3: string mobile,   // 手机号码
  4: string avatar,   // 头像
  5: string desc      // 描述
}





// 获取一个请求
struct GetOneReq {
  1: i64 id (go.tag = "json:\"id,required\""),
}


// BoolRes
struct BoolRes {
  1: bool ok
}