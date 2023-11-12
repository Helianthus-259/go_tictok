// 引入用户相关的Thrift定义
include "user.thrift"

namespace go relation

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
    1: i64 status_code, // 表示操作结果的状态码
    2: string status_msg // 提供有关操作状态的额外信息的状态消息
}

struct FollowActionRequest {
    1: string token (api.query="token", api.vd="len($) > 0"), // 用户鉴权token
    2: i64 to_user_id (api.query="to_user_id"), // 对方用户id
    3: i32 action_type (api.query="action_type"), // 1-关注，2-取消关注
}

struct FollowActionResponse {
    1: BaseResp base_resp
}

struct FollowListRequest {
    1: i64 user_id (api.query="user_id"), // 用户id
    2: string token (api.query="token", api.vd="len($) > 0"), // 用户鉴权token
}

struct FollowListResponse {
    1: BaseResp base_resp
    2: list<user.User> user_list // 用户信息列表
}

struct FansListRequest {
    1: i64 user_id (api.query="user_id"), // 用户id
    2: string token (api.query="token", api.vd="len($) > 0"), // 用户鉴权token
}

struct FansListResponse {
    1: BaseResp base_resp
    2: list<user.User> user_list // 用户列表
}

struct FriendListRequest {
    1: i64 user_id (api.query="user_id"), // 用户id
    2: string token (api.query="token", api.vd="len($) > 0"), // 用户鉴权token
}

struct FriendListResponse {
    1: BaseResp base_resp
    2: list<FriendUser> user_list // 用户列表
}

struct FriendUser {
    1: i64 id, // 用户id
    2: string name, // 用户名称
    3: optional i64 follow_count, // 关注总数
    4: optional i64 follower_count, // 粉丝总数
    5: bool is_follow, // true-已关注，false-未关注
    6: optional string avatar, // 用户头像
    7: optional string background_image, // 用户个人页顶部大图
    8: optional string signature, // 个人简介
    9: optional i64 total_favorited, // 获赞数量
    10: optional i64 work_count, // 作品数量
    11: optional i64 favorite_count, // 点赞数量
    12: optional string chat_message, // 和该好友的最新聊天消息
    13: i32 msg_type // message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
}

service RelationService {
    FollowActionResponse FollowAction(1: FollowActionRequest request) (api.post="/douyin/relation/action/"),
    FollowListResponse FollowList(1: FollowListRequest request) (api.get="/douyin/relation/follow/list/"),
    FansListResponse FansList(1: FansListRequest request) (api.get="/douyin/relation/follower/list/"),
    FriendListResponse FriendList(1: FriendListRequest request) (api.get="/douyin/relation/friend/list/")
}
