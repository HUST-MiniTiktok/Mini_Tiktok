namespace go user

include "common.thrift"

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
    1: i32 status_code (go.tag="json:\"status_code\"")      // 状态码，0-成功，其他值-失败
    2: optional string status_msg                           // 返回状态描述
    3: common.User user                                     // 用户信息
}

struct CheckUserIsExistRequest {
    1: i64 user_id  // 用户id
}

struct CheckUserIsExistResponse {
    1: i32 status_code (go.tag="json:\"status_code\"")      // 状态码，0-成功，其他值-失败
    2: optional string status_msg                           // 返回状态描述
    3: bool is_exist                                        // 用户是否存在
}

service UserService {
    // 用户信息
    UserResponse User(1: UserRequest request) (api.get = "/douyin/user/")
    // 用户注册
    UserRegisterResponse Register(1: UserRegisterRequest request) (api.post = "/douyin/user/register/")
    // 用户登录
    UserLoginResponse Login(1: UserLoginRequest request) (api.post = "/douyin/user/login/")
    // 检查用户是否存在
    CheckUserIsExistResponse CheckUserIsExist(1: CheckUserIsExistRequest request)
}