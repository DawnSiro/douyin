namespace go user

enum ErrCode {
    SuccessCode                = 0
    ServiceErrCode             = 10001
    ParamErrCode               = 10002
    UserAlreadyExistErrCode    = 10003
    AuthorizationFailedErrCode = 10004
}

struct douyin_user_register_request {
  1: required string username (vt.min_size = "2", vt.max_size = "32") // 注册用户名，最长32个字符
  2: required string password (vt.min_size = "6", vt.max_size = "32", vt.pattern = "[0-9A-Za-z]+") // 密码，最长32个字符
}

struct douyin_user_register_response {
  1: required i64 status_code // 状态码，0-成功，其他值-失败
  2: optional string status_msg // 返回状态描述
  3: required i64 user_id (vt.gt = "0") // 用户id
  4: required string token // 用户鉴权token
}

struct douyin_user_login_request {
  1: required string username (vt.min_size = "2", vt.max_size = "32") // 登录用户名
  2: required string password (vt.min_size = "6", vt.max_size = "32") // 登录密码
}

struct douyin_user_login_response {
  1: required i64 status_code // 状态码，0-成功，其他值-失败
  2: optional string status_msg // 返回状态描述
  3: required i64 user_id (vt.gt = "0") // 用户id
  4: required string token // 用户鉴权token
}

struct douyin_user_request {
  1: required i64 user_id (vt.gt = "0") // 用户id
  2: required string token // 用户鉴权token
}

struct douyin_user_response {
  1: required i64 status_code // 状态码，0-成功，其他值-失败
  2: optional string status_msg // 返回状态描述
  3: required User user // 用户信息
}

struct User {
  1: required i64 id (vt.gt = "0") // 用户id
  2: required string name  // 用户名称
  3: optional i64 follow_count (vt.gt = "0")  // 关注总数
  4: optional i64 follower_count (vt.gt = "0")  // 粉丝总数
  5: required bool is_follow  // true-已关注，false-未关注
  6: required string avatar  // 用户头像Url
}

service UserService {
    douyin_user_register_response Register(1: douyin_user_register_request req)
    douyin_user_login_response Login(1: douyin_user_login_request req)
    douyin_user_response GetUserInfo(1: douyin_user_request req)
}