namespace go comment

include "common.thrift"

struct Comment {
    1: i64 id               // 视频评论id
    2: common.User user            // 评论用户信息
    3: string content       // 评论内容
    4: string create_date   // 评论发布日期，格式 mm-dd
}

struct CommentActionRequest {
    1: string token                 // 用户鉴权token
    2: i64 video_id                 // 视频id
    3: i32 action_type              // 1-发布评论，2-删除评论
    4: optional string comment_text // 用户填写的评论内容，在action_type=1的时候使用
    5: optional i64 comment_id      // 要删除的评论id，在action_type=2的时候使用
}

struct CommentActionResponse {
    1: i32 status_code (go.tag="json:\"status_code\"")  // 状态码，0-成功，其他值-失败
    2: optional string status_msg                       // 返回状态描述
    3: optional Comment comment                         // 评论成功返回评论内容，不需要重新拉取整个列表
}

struct CommentListRequest {
    1: string token // 用户鉴权token
    2: i64 video_id // 视频id
}

struct CommentListResponse {
    1: i32 status_code (go.tag="json:\"status_code\"")  // 状态码，0-成功，其他值-失败
    2: optional string status_msg                       // 返回状态描述
    3: list<Comment> comment_list                       // 评论列表
}

struct GetVideoCommentCountRequest {
    1: i64 video_id                                     // 视频id
}

struct GetVideoCommentCountResponse {
    1: i32 status_code (go.tag="json:\"status_code\"")  // 状态码，0-成功，其他值-失败
    2: optional string status_msg                       // 返回状态描述
    3: i64 comment_count                                // 评论总数
}

struct GetVideoCommentCountListRequest {
    1: list<i64> video_id_list                          // 视频id列表
}

struct GetVideoCommentCountListResponse {
    1: i32 status_code (go.tag="json:\"status_code\"")  // 状态码，0-成功，其他值-失败
    2: optional string status_msg                       // 返回状态描述
    3: list<i64> comment_count_list                     // 评论总数列表
}

service CommentService {
    // 评论操作
    CommentActionResponse CommentAction(1: CommentActionRequest request) (api.post = "/douyin/comment/action/")
    // 评论列表
    CommentListResponse CommentList(1: CommentListRequest request) (api.get = "/douyin/comment/list/")
    // 获取视频评论总数
    GetVideoCommentCountResponse GetVideoCommentCount(1: GetVideoCommentCountRequest request)
    // 根据视频ID列表获取评论数列表
    GetVideoCommentCountListResponse GetVideoCommentListCount(1: GetVideoCommentCountListRequest request)
}