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

// 媒体类型
struct MediaType {
  1: string uid,
  2: string url
}

// 消息
struct Message {
  1: string link_id (go.tag = "json:\"link_id,omitempty\""),
  2: string msg_id (go.tag = "json:\"msg_id,omitempty\""),
  3: i64 timestamp (go.tag = "json:\"timestamp,omitempty\""),
  4: ChatType chat_type (go.tag = "json:\"chat_type,omitempty\""),
  5: i64 sender (go.tag = "json:\"sender,omitempty\""),
  6: i64 receiver (go.tag = "json:\"receiver,omitempty\""),
  7: string content (go.tag = "json:\"content,omitempty\""),
  8: MsgType msg_type (go.tag = "json:\"msg_type,omitempty\""),
  9: list<MediaType> media (go.tag = "json:\"media,omitempty\""),
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
  1: i64 id
}


// BoolRes
struct BoolRes {
  1: bool ok
}