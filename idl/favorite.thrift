namespace go favorite

struct douyin_favorite_action_request {
  1: required string token  // 用户鉴权token
  2: required i64 video_id (vt.gt = "0", api.vd="$>0")  // 视频id
  3: required i8 action_type (vt.in = "1", vt.in = "2", api.vd = "$==1||$==2") // 1-点赞，2-取消点赞
}

struct douyin_favorite_action_response {
  1: required i64 status_code // 状态码，0-成功，其他值-失败
  2: optional string status_msg // 返回状态描述
}

struct douyin_favorite_list_request {
  1: required i64 user_id (vt.gt = "0", api.vd="$>0") // 用户id
  2: required string token // 用户鉴权token
}

struct douyin_favorite_list_response {
  1: required i64 status_code // 状态码，0-成功，其他值-失败
  2: optional string status_msg // 返回状态描述
  3: list<Video> video_list // 用户点赞视频列表
}

struct Video {
  1: required i64 id // 视频唯一标识
  2: required UserInfo author // 视频作者信息
  3: required string play_url // 视频播放地址
  4: required string cover_url // 视频封面地址
  5: required i64 favorite_count // 视频的点赞总数
  6: required i64 comment_count // 视频的评论总数
  7: required bool is_favorite // true-已点赞，false-未点赞
  8: required string title // 视频标题
}

struct UserInfo {
  1: required i64 id // 用户id
  2: required string name  // 用户名称
  3: required i64 follow_count  // 关注总数
  4: required i64 follower_count  // 粉丝总数
  5: required bool is_follow  // true-已关注，false-未关注
  6: required string avatar  // 用户头像Url
  7: required string background_image //用户个人页顶部大图
  8: required string signature //个人简介
  9: required i64 total_favorited //获赞数量
  10: required i64 work_count  // 用户作品数
  11: required i64 favorite_count  // 用户点赞的视频数
}


service FavoriteService {
    douyin_favorite_action_response FavoriteVideo(1: douyin_favorite_action_request req) (api.post="/douyin/favorite/action/")
    douyin_favorite_list_response GetFavoriteList(1: douyin_favorite_list_request req) (api.get="/douyin/favorite/list/")
}