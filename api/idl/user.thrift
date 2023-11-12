namespace go user

enum ErrCode {
    SuccessCode                = 0
    ServiceErrCode             = 10001
    ParamErrCode               = 10002
    UserAlreadyExistErrCode    = 10003
    AuthorizationFailedErrCode = 10004
    UpdateFollowCountFailedErrCode = 10005
    UpdateUserWorkCountFailedErrCode = 10006
    UpdateUserFavoriteCountFailedErrCode = 10007
}

struct BaseResp {
    1: i64 status_code// 表示操作结果的状态码
    2: string status_msg// 提供有关操作状态的额外信息的状态消息
}

// 用户注册请求结构体
struct RegisterRequest {
    1: string username (api.query="username") // 用户期望的用户名
    2: string password (api.query="password") // 用户密码
}

// 用户注册响应结构体
struct RegisterResponse {
    1: BaseResp base_resp
    2: i64 user_id // 注册用户的唯一标识符
    3: string token // 注册成功后提供的身份验证令牌
}

// 用户登录请求结构体
struct LoginRequest {
    1: string username (api.query="username") // 用户登录的用户名
    2: string password (api.query="password") // 用户登录的密码
}

// 用户登录响应结构体
struct LoginResponse {
    1: BaseResp base_resp
    2: i64 user_id // 登录用户的唯一标识符
    3: string token // 登录成功后提供的身份验证令牌
}

// 表示用户数据的结构体
struct User {
    1: i64 id; // 用户的唯一标识符
    2: string name; // 用户的显示名称
    3: optional i64 follow_count; // 用户关注的其他用户数量（可选）
    4: optional i64 follower_count; // 关注当前用户的其他用户数量（可选）
    5: bool is_follow; // 表示当前用户是否正在关注该用户
    6: optional string avatar; // 用户头像图片的URL或路径（可选）
    7: optional string background_image; // 用户背景图片的URL或路径（可选）
    8: optional string signature; // 用户的个人简介或签名（可选）
    9: optional i64 total_favorited; // 用户收藏的项目总数（可选）
    10: optional i64 work_count; // 用户上传的作品数量（可选）
    11: optional i64 favorite_count; // 其他用户收藏该用户作品的数量（可选）
}

// 请求用户个人资料的请求结构体
struct UserIndexRequest {
    1: i64 user_id (api.query="user_id") // 请求个人资料信息的用户ID
    2: string token // 身份验证令牌
}

// 用户个人资料响应结构体
struct UserIndexResponse {
    1: BaseResp base_resp
    2: User user // 用户个人资料信息
}


// 用户服务接口定义
service UserService {
    UserIndexResponse UserIndex(1: UserIndexRequest request) (api.get="/douyin/user/"), // 用户个人资料查询接口
    RegisterResponse UserRegister(1: RegisterRequest request) (api.post="/douyin/user/register/"), // 用户注册接口
    LoginResponse UserLogin(1: LoginRequest request) (api.post="/douyin/user/login/"), // 用户登录接口
}
