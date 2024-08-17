namespace go user
include "base.thrift"

// 用户名密码登录请求
struct LoginByPasswordReq {
  1: string username,  // 用户名
  2: string password   // 密码
}

// 手机验证码登录请求
struct LoginByMobileReq {
  1: string mobile,     // 手机号码
  2: string captcha,    // 验证码
  3: string session_id  // 会话ID
}

// 登录响应
struct LoginRes {
  1: string token,  // 令牌
  2: i64 expire     // 过期时间
}

// 发送验证码请求
struct SendCaptchaReq {
  1: string session_id,  // 会话ID
  2: string mobile       // 手机号码
}


// 用户在线状态
struct UserOnlineStatus {
  1: bool status  // 是否在线
}

// 添加好友请求
struct AddFriendReq {
  1: i64 owner,  // 本人id
  2: i64 with   // 聊天 
}

// 用户服务
service UserSrv {
  // 用户名密码登录
  LoginRes LoginByPassword(1: LoginByPasswordReq loginByPasswordReq),
  // 验证码登录
  LoginRes LoginByMobile(1: LoginByMobileReq loginByMobileReq),
  // 发送手机验证码
  base.BoolRes SendCaptcha(1: SendCaptchaReq sendCaptchaReq),
  // 获取单个用户信息
  base.UserInfo GetOneUserInfo(1: base.GetOneReq getOneReq),
  // 获取用户在线状态
  base.BoolRes GetUserOnlineStatus(1: base.GetOneReq getOneReq),
  // 获取用户聊天列表
  list<base.ChatList> GetUserChatList(1: base.GetOneReq getOneReq)
  // 添加好友
  base.BoolRes AddFriend(1: AddFriendReq addFriendReq),
}
