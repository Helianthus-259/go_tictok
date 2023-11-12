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
	1: i64 user_id, // 用户id
	2: i64 my_user_id // 当前登录用户
}

struct FavoriteListResponse {
    1: BaseResp base_resp
	2: list<video.Video> video_list // 用户点赞视频列表
}

struct FavoriteActionRequest {
	1: i64 user_id,
	2: i64 video_id, // 视频id
	3: i32 action_type // 1-点赞，2-取消点赞
}

struct FavoriteActionResponse {
    1: BaseResp base_resp
}

struct IsFavoriteVideosRequest {
	1: i64 user_id,
	2: list<i64> video_ids // 视频id
}

struct IsFavoriteVideosResponse {
    1: BaseResp base_resp
	2: list<bool> many_is_favorite // 很多个是否点赞
}

service FavoriteService {
	FavoriteListResponse FavoriteList(1: FavoriteListRequest request),
	FavoriteActionResponse FavoriteAction(1: FavoriteActionRequest request),
	IsFavoriteVideosResponse IsFavoriteVideos(1: IsFavoriteVideosRequest request)
}

