namespace go favorite

struct User {
    1: i64 id                           // 用户id
    2: string name                      // 用户名称
    3: optional i64 follow_count        // 关注总数
    4: optional i64 follower_count      // 粉丝总数
    5: bool is_follow                   // true-已关注，false-未关注
    6: optional string avatar           // 用户头像
    7: optional string background_image // 用户个人页顶部大图
    8: optional string signature        // 个人简介
    9: optional i64 total_favorited     // 获赞数量
    10: optional i64 work_count         // 作品数量
    11: optional i64 favorite_count     // 点赞数量
}

struct Video {
    1: i64 id               // 视频唯一标识
    2: User author          // 视频作者信息
    3: string play_url      // 视频播放地址
    4: string cover_url     // 视频封面地址
    5: i64 favorite_count   // 视频的点赞总数
    6: i64 comment_count    // 视频的评论总数
    7: bool is_favorite     // true-已点赞，false-未点赞
    8: string title         // 视频标题
}

struct FavoriteActionRequest {
    1: string token     // 用户鉴权token
    2: i64 video_id     // 视频id
    3: i32 action_type  // 1-点赞，2-取消点赞
}

struct FavoriteActionResponse {
    1: i32 status_code (go.tag="json:\"status_code\"")  // 状态码，0-成功，其他值-失败
    2: optional string status_msg                       // 返回状态描述
}

struct FavoriteListRequest {
    1: i64 user_id  // 用户id
    2: string token // 用户鉴权token
}

struct FavoriteListResponse {
    1: i32 status_code (go.tag="json:\"status_code\"")  // 状态码，0-成功，其他值-失败
    2: optional string status_msg                       // 返回状态描述
    3: list<Video> video_list                           // 用户点赞视频列表
}

service FavoriteService {
    FavoriteActionResponse FavoriteAction (1: FavoriteListRequest request)
    FavoriteListResponse FavoriteList (1: FavoriteListRequest request)
}

