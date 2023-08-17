namespace go feed

include "common.thrift"

struct FeedRequest {
    1: optional i64 latest_timestamp    // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
    2: optional string token            // 可选参数，登录用户设置
}

struct FeedResponse {
    1: i32 status_code (go.tag="json:\"status_code\"")  // 状态码，0-成功，其他值-失败
    2: optional string status_msg                       // 返回状态描述
    3: list<common.Video> video_list                           // 视频列表
    4: optional i64 next_time                           // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

service FeedService {
    // 视频流接口
    FeedResponse GetFeed(1: FeedRequest request) (api.get = "/douyin/feed/")
}
