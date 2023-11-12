namespace go message

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

struct MessageListRequest {
	1: i64 user_id,
	2: i64 to_user_id, // 对方用户id
	3: i64 pre_msg_time // 上次最新消息的时间（新增字段-apk更新中）
}

struct MessageListResponse {
    1: BaseResp base_resp
	2: list<Message> message_list // 消息列表
}

struct Message {
	1: i64 id, // 消息id
	2: i64 to_user_id, // 该消息接收者的id
	3: i64 from_user_id, // 该消息发送者的id
	4: string content, // 消息内容
	5: i64 create_time // 消息创建时间
}

struct ChatActionRequest {
	1: i64 user_id,
	2: i64 to_user_id, // 对方用户id
	3: i32 action_type, // 1-发送消息
	4: string content // 消息内容
}

struct ChatActionResponse {
    1: BaseResp base_resp
}

struct GetFriendLatestMessageRequest {
	1: i64 user_id,
	2: list<i64> friend_ids
}

struct GetFriendLatestMessageResponse {
    1: BaseResp base_resp
	2: list<string> message_list,
	3: list<i32> msg_type_list
}

service MessageService {
	MessageListResponse MessageList(1: MessageListRequest request),
	ChatActionResponse ChatAction(1: ChatActionRequest request),
	GetFriendLatestMessageResponse GetFriendLatestMessage(1: GetFriendLatestMessageRequest request)
}

