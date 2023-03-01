namespace go comment

struct douyin_comment_action_request {
  1: required string token       // 用户鉴权token
  2: required i64 video_id (vt.gt = "0", api.vd="$>0")      // 视频id
  3: required i8 action_type (vt.in = "1", vt.in = "2", api.vd = "$==1||$==2")   // 1-发布评论，2-删除评论
  4: optional string comment_text (vt.min_size = "1", vt.max_size = "255", api.vd = "$=nil||(len($)>0&&len($)<256)") // 用户填写的评论内容，在action_type=1的时候使用
  5: optional i64 comment_id (vt.gt = "0", api.vd="$=nil||$>0")   // 要删除的评论id，在action_type=2的时候使用
}

struct douyin_comment_action_response {
  1: required i64 status_code // 状态码，0-成功，其他值-失败
  2: optional string status_msg // 返回状态描述
  3: optional Comment comment // 评论成功返回评论内容，不需要重新拉取整个列表
}

struct douyin_comment_list_request {
  1: required string token // 用户鉴权token
  2: required i64 video_id (vt.gt = "0", api.vd="$>0") // 视频id
}

struct douyin_comment_list_response {
  1: required i64 status_code // 状态码，0-成功，其他值-失败
  2: optional string status_msg // 返回状态描述
  3: list<Comment> comment_list // 评论列表
}

struct Comment {
  1: required i64 id // 视频评论id
  2: required User user // 评论用户信息
  3: required string content // 评论内容
  4: required string create_date // 评论发布日期，格式 mm-dd
}

struct User {
  1: required i64 id // 用户id
  2: required string name  // 用户名称
  3: optional i64 follow_count  // 关注总数
  4: optional i64 follower_count  // 粉丝总数
  5: required bool is_follow  // true-已关注，false-未关注
  6: required string avatar  // 用户头像Url
}

service CommentService {
    douyin_comment_action_response CommentAction(1: douyin_comment_action_request req) (api.post="/douyin/comment/action/")
    douyin_comment_list_response GetCommentList(1: douyin_comment_list_request req) (api.get="/douyin/comment/list/")
}