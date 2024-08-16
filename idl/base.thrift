namespace go base

// 聊天类型
enum ChatType {
  PrivateChat = 0,  // 私聊 （处理私聊消息）
  GroupChat = 1,    // 群聊 （处理群聊消息）
  SystemMessage = 2,// 系统消息 （接收）
  SystemNotice = 3  // 系统通知 （接收）
}

// 消息类型
enum MsgType {
  Text = 0,
  Image = 1,
  File = 2,
  Audio = 3,
  Video = 4
}

// 媒体类型
struct MediaType {
  1: string uid,
  2: string url
}

// 消息
struct Message {
  1: string link_id,
  2: string msg_id,
  3: i64 timestamp,
  4: ChatType chat_type,
  5: i64 sender,
  6: i64 receiver,
  7: i64 chat_id,
  8: i64 group_id,
  9: string content,
  10: MsgType msg_type,
  11: list<MediaType> media
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
    1:string uuid,
    2:ChatType chat_type,
    3:i64 group_id,
    4:i64 chat_id,
    5:i64 user1_id,
    6:i64 user2_id,
    7:Message last_msg,
    8:i64 timestamp
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
  1: bool success
}