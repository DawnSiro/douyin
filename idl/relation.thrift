namespace go relation

struct douyin_relation_action_request {
  1: required string token // 用户鉴权token
  2: required i64 to_user_id (vt.gt = "0", api.vd="$>0") // 对方用户id
  3: required i8 action_type (vt.in = "1", vt.in = "2", api.vd = "$==1||$==2") // 1-关注，2-取消关注
}

struct douyin_relation_action_response {
  1: required i64 status_code // 状态码，0-成功，其他值-失败
  2: optional string status_msg // 返回状态描述
}

struct douyin_relation_follow_list_request {
  1: required i64 user_id (vt.gt = "0", api.vd="$>0") // 用户id
  2: required string token // 用户鉴权token
}

struct douyin_relation_follow_list_response {
  1: required i64 status_code // 状态码，0-成功，其他值-失败
  2: optional string status_msg // 返回状态描述
  3: list<User> user_list // 用户信息列表
}

struct douyin_relation_follower_list_request {
  1: required i64 user_id (vt.gt = "0", api.vd="$>0")  // 用户id
  2: required string token // 用户鉴权token
}

struct douyin_relation_follower_list_response {
  1: required i64 status_code // 状态码，0-成功，其他值-失败
  2: optional string status_msg  // 返回状态描述
  3: list<User> user_list  // 用户列表
}


struct douyin_relation_friend_list_request {
  1: required i64 user_id (vt.gt = "0", api.vd="$>0")  // 用户id
  2: required string token  // 用户鉴权token
}

struct douyin_relation_friend_list_response {
  1: required i64 status_code  // 状态码，0-成功，其他值-失败
  2: optional string status_msg  // 返回状态描述
  3: list<FriendUser> user_list // 用户列表
}

struct User {
  1: required i64 id // 用户id
  2: required string name  // 用户名称
  3: optional i64 follow_count  // 关注总数
  4: optional i64 follower_count  // 粉丝总数
  5: required bool is_follow  // true-已关注，false-未关注
  6: required string avatar  // 用户头像Url
}

struct FriendUser {
  1: required i64 id // 用户id
  2: required string name  // 用户名称
  3: optional i64 follow_count  // 关注总数
  4: optional i64 follower_count   // 粉丝总数
  5: required bool is_follow  // true-已关注，false-未关注
  6: required string avatar  // 用户头像Url
  7: optional string message // 和该好友的最新聊天消息
  8: required i8 msgType // message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
}


service RelationService {
    douyin_relation_action_response Follow(1: douyin_relation_action_request req) (api.post="/douyin/relation/action/")
    douyin_relation_follow_list_response GetFollowList(1: douyin_relation_follow_list_request req) (api.get="/douyin/relation/follow/list/")
    douyin_relation_follower_list_response GetFollowerList(1: douyin_relation_follower_list_request req) (api.get="/douyin/relation/follower/list/")
    douyin_relation_friend_list_response GetFriendList(1: douyin_relation_friend_list_request req) (api.get="/douyin/relation/friend/list/")
}