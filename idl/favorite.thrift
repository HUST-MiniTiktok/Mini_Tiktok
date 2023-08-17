namespace go favorite

include "common.thrift"

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
    3: list<common.Video> video_list                           // 用户点赞视频列表
}

struct GetVideoFavoriteCountRequest {
    1: i64 video_id // 视频id
}

struct GetVideoFavoriteCountResponse {
    1: i32 status_code (go.tag="json:\"status_code\"")  // 状态码，0-成功，其他值-失败
    2: optional string status_msg                       // 返回状态描述
    3: i64 favorite_count                               // 视频点赞总数
}

struct CheckIsFavoriteRequest {
    1: i64 video_id // 视频id
    2: i64 user_id  // 用户id
}

struct CheckIsFavoriteResponse {
    1: i32 status_code (go.tag="json:\"status_code\"")  // 状态码，0-成功，其他值-失败
    2: optional string status_msg                       // 返回状态描述
    3: bool is_favorite                                 // true-已点赞，false-未点赞
}

service FavoriteService {
    // 赞操作
    FavoriteActionResponse FavoriteAction (1: FavoriteActionRequest request) (api.post = "/douyin/favorite/action/")
    // 喜欢列表
    FavoriteListResponse FavoriteList (1: FavoriteListRequest request) (api.get = "/douyin/favorite/list/")

    // 获取视频点赞总数
    GetVideoFavoriteCountResponse GetVideoFavoriteCount (1: GetVideoFavoriteCountRequest request)
    // 检查视频是否已点赞
    CheckIsFavoriteResponse CheckIsFavorite (1: CheckIsFavoriteRequest request)
}

