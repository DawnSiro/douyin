namespace go message

struct douyin_message_chat_request {
  1: required string token // 用户鉴权token
  2: required i64 to_user_id (vt.gt = "0", api.vd="$>0")  // 对方用户id
  3: optional i64 pre_msg_time //上次最新消息的时间
}

struct douyin_message_chat_response {
  1: required i64 status_code // 状态码，0-成功，其他值-失败
  2: optional string status_msg // 返回状态描述
  3: list<Message> message_list // 消息列表
}

struct Message {
  1: required i64 id // 消息id
  2: required i64 to_user_id // 该消息接收者的id
  3: required i64 from_user_id // 该消息发送者的id
  4: required string content // 消息内容
  5: optional i64 create_time // 消息创建时间
}

struct douyin_message_action_request {
  1: required string token // 用户鉴权token
  2: required i64 to_user_id (vt.gt = "0", api.vd="$>0") // 对方用户id
  3: required i8 action_type (vt.in = "1", api.vd="$==1") // 1-发送消息
  4: required string content (vt.min_size = "1", vt.max_size = "255", api.vd = "len($)>0&&len($)<256") // 消息内容
}

struct douyin_message_action_response {
  1: required i64 status_code // 状态码，0-成功，其他值-失败
  2: optional string status_msg // 返回状态描述
}


service MessageService {
    douyin_message_action_response SendMessage(1: douyin_message_action_request req) (api.post="/douyin/message/action/")
    douyin_message_chat_response GetMessageChat(1: douyin_message_chat_request req) (api.get="/douyin/message/chat/")
}