// 引入用户相关的Thrift定义
include "user.thrift"

namespace go video

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

// 请求视频动态流的结构体
struct VideoFeedRequest {
    1: optional i64 latest_time; // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
    2: optional string token (api.query="token"); // 可选参数，登录用户设置
    3: i64 user_id; // 用户ID
}

// 视频动态流响应结构体
struct VideoFeedResponse {
    1: BaseResp base_resp
    2: list<Video> video_list; // 视频列表
    3: optional i64 next_time; // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

// 视频结构体
struct Video {
    1: i64 id; // 视频唯一标识
    2: user.User author; // 视频作者信息
    3: string play_url; // 视频播放地址
    4: string cover_url; // 视频封面地址
    5: i64 favorite_count; // 视频的点赞总数
    6: i64 comment_count; // 视频的评论总数
    7: bool is_favorite; // true-已点赞，false-未点赞
    8: string title; // 视频标题
}

// 发布视频请求结构体
struct VideoPublishRequest {
    1: i64 user_id; // 用户ID
    2: binary data; // 视频数据
    3: string title; // 视频标题
}

// 发布视频响应结构体
struct VideoPublishResponse {
    1: BaseResp base_resp
}

// 获取用户发布的视频列表请求结构体
struct PublishListRequest {
    1: i64 user_id; // 用户id
    2: i64 my_user_id; // 当前登录用户的id
}

// 获取用户发布的视频列表响应结构体
struct PublishListResponse {
    1: BaseResp base_resp
    2: list<Video> video_list; // 用户发布的视频列表
}

// 获取视频信息请求结构体
struct GetVideoInfoRequest {
    1: i64 user_id; // 用户ID
    2: i64 video_id; // 视频ID
}

// 获取视频信息响应结构体
struct GetVideoInfoResponse {
    1: BaseResp base_resp
    2: Video video_info; // 视频信息
}

// 批量获取视频信息请求结构体
struct GetManyVideoInfosRequest {
    1: i64 user_id; // 用户ID
    2: list<i64> video_ids; // 视频ID列表
}

// 批量获取视频信息响应结构体
struct GetManyVideoInfosResponse {
    1: BaseResp base_resp
    2: list<Video> video_infos; // 视频信息列表
}

// 增加视频点赞数请求结构体
struct AddVideoFavoriteCountRequest {
    1: i64 video_id; // 视频ID
}

// 增加视频点赞数响应结构体
struct AddVideoFavoriteCountResponse {
    1: BaseResp base_resp
}

// 减少视频点赞数请求结构体
struct SubVideoFavoriteCountRequest {
    1: i64 video_id; // 视频ID
}

// 减少视频点赞数响应结构体
struct SubVideoFavoriteCountResponse {
    1: BaseResp base_resp
}

// 增加视频评论数请求结构体
struct AddVideoCommentCountRequest {
    1: i64 video_id; // 视频ID
}

// 增加视频评论数响应结构体
struct AddVideoCommentCountResponse {
    1: BaseResp base_resp
}

// 减少视频评论数请求结构体
struct SubVideoCommentCountRequest {
    1: i64 video_id; // 视频ID
}

// 减少视频评论数响应结构体
struct SubVideoCommentCountResponse {
    1: BaseResp base_resp
}

// 视频服务接口定义
service VideoService {
    VideoFeedResponse VideoFeed(1: VideoFeedRequest request), // 视频动态流接口
    VideoPublishResponse VideoPublish(1: VideoPublishRequest request), // 发布视频接口
    PublishListResponse PublishList(1: PublishListRequest request), // 获取用户发布的视频列表接口
    GetVideoInfoResponse GetVideoInfo(1: GetVideoInfoRequest request), // 获取视频信息接口
    GetManyVideoInfosResponse GetManyVideoInfos(1: GetManyVideoInfosRequest request), // 批量获取视频信息接口
    AddVideoFavoriteCountResponse AddVideoFavoriteCount(1: AddVideoFavoriteCountRequest request), // 增加视频点赞数接口
    SubVideoFavoriteCountResponse SubVideoFavoriteCount(1: SubVideoFavoriteCountRequest request), // 减少视频点赞数接口
    AddVideoCommentCountResponse AddVideoCommentCount(1: AddVideoCommentCountRequest request), // 增加视频评论数接口
    SubVideoCommentCountResponse SubVideoCommentCount(1: SubVideoCommentCountRequest request); // 减少视频评论数接口
}
