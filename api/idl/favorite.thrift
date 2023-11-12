// 引入用户相关的Thrift定义
include "video.thrift"
namespace go favorite

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


struct FavoriteListRequest {
    1: i64 user_id (api.query="user_id"), // 用户id
    2: string token (api.query="token", api.vd="len($) > 0"), // 用户鉴权token
}

struct FavoriteListResponse {
    1: BaseResp base_resp
    2: list<video.Video> video_list // 用户点赞视频列表
}

struct FavoriteActionRequest {
    1: string token (api.query="token", api.vd="len($) > 0"), // 用户鉴权token
    2: i64 video_id (api.query="video_id"), // 视频id
    3: i32 action_type (api.query="action_type")// 1-点赞，2-取消点赞
}

struct FavoriteActionResponse {
    1: BaseResp base_resp
}

service FavoriteService {
    FavoriteListResponse FavoriteList(1: FavoriteListRequest request) (api.post="/douyin/favorite/action/"),
    FavoriteActionResponse FavoriteAction(1: FavoriteActionRequest request) (api.get="/douyin/favorite/list/")
}
