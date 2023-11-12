namespace go user

enum ErrCode {
    SuccessCode                = 0,
    ServiceErrCode             = 1,
    ParamErrCode               = 2,
    MysqlErrCode               = 3,
    RedisErrCode               = 4,

    UserIsNotExistErrCode      = 10005,
    PasswordIsNotVerifiedCode  = 10006,
    AuthorizationFailedErrCode = 10007,
    UserAlreadyExistErrCode    = 10008,
    UpdateUserCountFailedErrCode = 10009,

    FollowRelationAlreadyExistErrCode = 20010,
    FollowRelationNotExistErrCode      = 20011,

    FavoriteActionErrCode            = 30012,
    FavoriteAddFailedCode             = 30013,
    FavoriteRelationAlreadyExistErrCode = 30014,
    FavoriteRelationNotExistErrCode = 30015,

    ChatActionErrCode          = 40016,
    MessageAddFailedErrCode    = 40017,
    FriendListNoPermissionErrCode = 40018,

    VideoFeedErrCode              = 50019,
    VideoIsNotExistErrCode        = 50020,
    UpdateVideoCountFailedErrCode = 50021,

    CommentActionErrCode     = 60022,
    CommentIsNotExistErrCode  = 60023,
    CommentAddFailedErrCode   = 60024,
}

struct BaseResp {
    1: i64 status_code// 表示注册操作结果的状态码
    2: string status_msg// 提供有关注册状态的额外信息的状态消息
}

// 用户注册请求结构体
struct RegisterRequest {
    1: string username (vt.min_size = "1")// 用户期望的用户名
    2: string password (vt.min_size = "1")// 用户密码
}

// 用户注册响应结构体
struct RegisterResponse {
    1: BaseResp base_resp
    2: i64 user_id // 注册用户的唯一标识符
    3: string token // 注册成功后提供的身份验证令牌
}

// 用户登录请求结构体
struct LoginRequest {
    1: string username (vt.min_size = "1")// 用户登录的用户名
    2: string password (vt.min_size = "1")// 用户登录的密码
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
    1: i64 user_id (vt.min_size = "1")// 请求个人资料信息的用户ID
    2: i64 my_user_id (vt.min_size = "1") // 请求者的用户ID（可选）
}

// 用户个人资料响应结构体
struct UserIndexResponse {
    1: BaseResp base_resp
    2: User user; // 用户个人资料信息
}

// 增加关注数请求结构体
struct AddFollowCountRequest {
    1: i64 user_id; // 当前用户的用户ID
    2: i64 target_id; // 要关注的用户的用户ID
}

// 增加关注数响应结构体
struct AddFollowCountResponse {
    1: BaseResp base_resp
}

// 减少关注数请求结构体
struct SubFollowCountRequest {
    1: i64 user_id; // 当前用户的用户ID
    2: i64 target_id; // 要取消关注的用户的用户ID
}

// 减少关注数响应结构体
struct SubFollowCountResponse {
    1: BaseResp base_resp
}

// 获取用户列表请求结构体
struct GetUserListRequest {
    1: i64 user_id; // 用户ID，用于获取用户列表
    2: list<i64> target_id; // 要获取用户列表的目标用户ID列表
}

// 获取用户列表响应结构体
struct GetUserListResponse {
    1: BaseResp base_resp
    2: list<User> user_list; // 用户列表信息
}

// 增加用户作品数请求结构体
struct AddUserWorkCountRequest {
    1: i64 user_id; // 用户ID，用于增加用户作品数
}

// 增加用户作品数响应结构体
struct AddUserWorkCountResponse {
    1: BaseResp base_resp
}

// 更新用户收藏数请求结构体
struct UpdateUserFavoriteCountRequest {
    1: i64 user_id; // 用户ID，用于更新用户收藏数
    2: i64 author_id; // 作者的用户ID
    3: i32 action_type; // 操作类型：1-喜欢，2-取消喜欢
}

// 更新用户收藏数响应结构体
struct UpdateUserFavoriteCountResponse {
    1: BaseResp base_resp
}

// 用户服务接口定义
service UserService {
    RegisterResponse Register(1: RegisterRequest request), // 用户注册接口
    LoginResponse Login(1: LoginRequest request), // 用户登录接口
    UserIndexResponse UserIndex(1: UserIndexRequest request); // 用户个人资料查询接口
    AddFollowCountResponse AddFollowCount(1: AddFollowCountRequest request), // 增加关注数接口
    SubFollowCountResponse SubFollowCount(1: SubFollowCountRequest request), // 减少关注数接口
    GetUserListResponse GetUserList(1: GetUserListRequest request), // 获取用户列表接口
    AddUserWorkCountResponse AddUserWorkCount(1: AddUserWorkCountRequest request), // 增加用户作品数接口
    UpdateUserFavoriteCountResponse UpdateUserFavoriteCount(1: UpdateUserFavoriteCountRequest request); // 更新用户收藏数接口
}
