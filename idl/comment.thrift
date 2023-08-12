namespace go comment

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

struct Comment {
    1: i64 id               // 视频评论id
    2: User user            // 评论用户信息
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

service CommentService {
    // 评论操作
    CommentActionResponse CommentAction(1: CommentActionRequest request) (api.post = "/douyin/comment/action/")
    // 评论列表
    CommentListResponse CommentList(1: CommentListRequest request) (api.get = "/douyin/comment/list/")
}