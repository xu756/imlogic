namespace go user
include "base.thrift"

// 用户名密码登录请求
struct LoginByPasswordReq {
  1: string username (go.tag = "json:\"username,required\""),  // 用户名
  2: string password  (go.tag = "json:\"password,required\"")  // 密码
}

// 手机验证码登录请求
struct LoginByMobileReq {
  1: string mobile (go.tag = "json:\"mobile,required\""),  // 手机号码
  2: string captcha (go.tag = "json:\"captcha,required\"")  // 验证码
}

// 登录响应
struct LoginRes {
  1: string token,  // 令牌
  2: i64 expire     // 过期时间
}

// 发送验证码请求
struct SendCaptchaReq {
  1: string mobile (go.tag = "json:\"mobile,required\"")  // 手机号码
}

// 好友
struct Friend {
  1: i64 user_id,         // 用户ID
  2: string alias   // 别名
  3: string owner_desc  // 所属描述
  4: i64 created_at // 添加时间
}



// 添加好友请求
struct AddFriendReq {
  1: i64 owner (go.tag = "json:\"owner,required\""),  // 用户ID
  2: i64 with_id (go.tag = "json:\"with_id,required\"")  // 好友ID
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
  // 获取用户的所有好友
  list<Friend> GetUserFriendList(1: base.GetOneReq getOneReq),
  // 添加好友
  base.BoolRes AddFriend(1: AddFriendReq addFriendReq),
}
