namespace go user

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

struct UserRegisterRequest {
    1: string username (vt.min_size = "1", vt.max_size = "32") // 注册用户名，最长32个字符
    2: string password (vt.min_size = "1", vt.max_size = "32") // 注册密码，最长32个字符
}

struct UserRegisterResponse {
    1: i32 status_code (go.tag="json:\"status_code\"")   // 状态码，0-成功，其他值-失败
    2: optional string status_msg // 返回状态描述
    3: i64 user_id                // 用户id
    4: string token               // 用户鉴权token
}

struct UserLoginRequest {
    1: string username (vt.min_size = "1", vt.max_size = "32") // 登录用户名，最长32个字符
    2: string password (vt.min_size = "1", vt.max_size = "32") // 登录密码，最长32个字符
}

struct UserLoginResponse {
    1: i32 status_code (go.tag="json:\"status_code\"")  // 状态码，0-成功，其他值-失败
    2: optional string status_msg                       // 返回状态描述
    3: i64 user_id                                      // 用户id
    4: string token                                     // 用户鉴权token
}

struct UserRequest {
    1: i64 user_id  // 用户id
    2: string token // 用户鉴权token
}

struct UserResponse {
    1: i32 status_code (go.tag="json:\"status_code\"")  // 状态码，0-成功，其他值-失败
    2: optional string status_msg                       // 返回状态描述
    3: User user                                        // 用户信息
}

service UserService {
    // 用户信息
    UserResponse User(1: UserRequest request) (api.get = "/douyin/user/")
    // 用户注册
    UserRegisterResponse Register(1: UserRegisterRequest request) (api.post = "/douyin/user/register/")
    // 用户登录
    UserLoginResponse Login(1: UserLoginRequest request) (api.post = "/douyin/user/login/")
}