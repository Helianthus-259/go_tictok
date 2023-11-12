// 引入用户相关的Thrift定义
include "user.thrift"

namespace go comment

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

struct CommentActionRequest {
	1: i64 user_id,
	2: i64 video_id, // 视频id
	3: i32 action_type, // 1-发布评论，2-删除评论
	4: optional string comment_text, // 用户填写的评论内容，在action_type=1的时候使用
	5: optional i64 comment_id // 要删除的评论id，在action_type=2的时候使用
}

struct CommentActionResponse {
    1: BaseResp base_resp
	2: optional Comment comment // 评论成功返回评论内容，不需要重新拉取整个列表
}

struct Comment {
	1: i64 id, // 视频评论id
	2: user.User user, // 评论用户信息
	3: string content, // 评论内容
	4: string create_date // 评论发布日期，格式 mm-dd
}

struct CommentListRequest {
	1: i64 user_id,
	2: i64 video_id // 视频id
}

struct CommentListResponse {
    1: BaseResp base_resp
	2: list<Comment> comment_list // 评论列表
}

service CommentService {
	CommentActionResponse CommentAction(1: CommentActionRequest request),
	CommentListResponse CommentList(1: CommentListRequest request)
}
