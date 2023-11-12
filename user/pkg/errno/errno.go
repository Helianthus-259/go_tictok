package errno

import (
	"errors"
	"fmt"
)

//const (
//	SuccessCode    = 0
//	ServiceErrCode = iota + 10000
//	ParamErrCode
//	MysqlErrCode
//	RedisErrCode
//
//	UserIsNotExistErrCode
//	PasswordIsNotVerifiedCode
//	AuthorizationFailedErrCode
//	UserAlreadyExistErrCode
//	UpdateUserCountFailedErrCode
//
//	FollowRelationAlreadyExistErrCode
//	FollowRelationNotExistErrCode
//
//	FavoriteActionErrCode
//	FavoriteAddFailedCode
//	FavoriteRelationAlreadyExistErrCode
//	FavoriteRelationNotExistErrCode
//
//	ChatActionErrCode
//	MessageAddFailedErrCode
//	FriendListNoPermissionErrCode
//
//	VideoFeedErrCode
//	VideoIsNotExistErrCode
//	UpdateVideoCountFailedErrCode
//
//	CommentActionErrCode
//	CommentIsNotExistErrCode
//	CommentAddFailedErrCode
//)

const (
	SuccessCode = iota
	ServiceErrCode
	ParamErrCode
	MysqlErrCode
	RedisErrCode

	UserIsNotExistErrCode = iota + 10000
	PasswordIsNotVerifiedCode
	AuthorizationFailedErrCode
	UserAlreadyExistErrCode
	UpdateUserCountFailedErrCode

	FollowRelationAlreadyExistErrCode = iota + 20000
	FollowRelationNotExistErrCode

	FavoriteActionErrCode = iota + 30000
	FavoriteAddFailedCode
	FavoriteRelationAlreadyExistErrCode
	FavoriteRelationNotExistErrCode

	ChatActionErrCode = iota + 40000
	MessageAddFailedErrCode
	FriendListNoPermissionErrCode

	VideoFeedErrCode = iota + 50000
	VideoIsNotExistErrCode
	UpdateVideoCountFailedErrCode

	CommentActionErrCode = iota + 60000
	CommentIsNotExistErrCode
	CommentAddFailedErrCode
)

const (
	SuccessMsg   = "Success"
	ServerErrMsg = "Service is unable to start successfully"
	ParamErrMsg  = "Wrong Parameter has been given"
	MysqlErrMsg  = "MySQL error"
	RedisErrMsg  = "Redis error"

	UserIsNotExistErrMsg        = "User is not exist"
	PasswordIsNotVerifiedMsg    = "Username or password not verified"
	AuthorizationFailedErrMsg   = "Authorization failed"
	UserAlreadyExistErrMsg      = "User already exists"
	UpdateUserCountFailedErrMsg = "Update user count failed"

	FollowRelationAlreadyExistMsg = "Follow relation already exists"
	FollowRelationNotExistMsg     = "Follow relation does not exist"

	FavoriteActionErrMsg            = "Favorite action failed"
	FavoriteAddFailedErrMsg         = "Favorite add failed"
	FavoriteRelationAlreadyExistMsg = "Favorite relation already exists"
	FavoriteRelationNotExistMsg     = "Favorite relation does not exist"

	ChatActionErrMsg          = "Chat action failed"
	MessageAddFailedErrMsg    = "Message add failed"
	FriendListNoPermissionMsg = "You can't query his friend list"

	VideoFeedErrMsg              = "Video feed error"
	UpdateVideoCountFailedErrMsg = "Update video count failed"
	VideoIsNotExistErrMsg        = "Video is not exist"

	CommentActionErrMsg     = "Comment action error"
	CommentIsNotExistErrMsg = "Comment is not exist"
	CommentAddFailedErrMsg  = "Comment add error"
)

type ErrNo struct {
	ErrCode int64  `json:"status_code"`          // 状态码，0-成功，其他值-失败
	ErrMsg  string `json:"status_msg,omitempty"` // 返回状态描述
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{code, msg}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success    = NewErrNo(SuccessCode, SuccessMsg)
	ServiceErr = NewErrNo(ServiceErrCode, ServerErrMsg)
	ParamErr   = NewErrNo(ParamErrCode, ParamErrMsg)
	MysqlErr   = NewErrNo(MysqlErrCode, MysqlErrMsg)
	RedisErr   = NewErrNo(RedisErrCode, RedisErrMsg)

	UserIsNotExistErr        = NewErrNo(UserIsNotExistErrCode, UserIsNotExistErrMsg)
	PasswordIsNotVerifiedErr = NewErrNo(PasswordIsNotVerifiedCode, PasswordIsNotVerifiedMsg)
	AuthorizationFailedErr   = NewErrNo(AuthorizationFailedErrCode, AuthorizationFailedErrMsg)
	UserAlreadyExistErr      = NewErrNo(UserAlreadyExistErrCode, UserAlreadyExistErrMsg)
	UpdateUserCountFailedErr = NewErrNo(UpdateUserCountFailedErrCode, UpdateUserCountFailedErrMsg)

	FollowRelationAlreadyExistErr = NewErrNo(FollowRelationAlreadyExistErrCode, FollowRelationAlreadyExistMsg)
	FollowRelationNotExistErr     = NewErrNo(FollowRelationNotExistErrCode, FollowRelationNotExistMsg)

	FavoriteActionErr               = NewErrNo(FavoriteActionErrCode, FavoriteActionErrMsg)
	FavoriteAddFailedErr            = NewErrNo(FavoriteAddFailedCode, FavoriteAddFailedErrMsg)
	FavoriteRelationAlreadyExistErr = NewErrNo(FavoriteRelationAlreadyExistErrCode, FavoriteRelationAlreadyExistMsg)
	FavoriteRelationNotExistErr     = NewErrNo(FavoriteRelationNotExistErrCode, FavoriteRelationNotExistMsg)

	ChatActionErr             = NewErrNo(ChatActionErrCode, ChatActionErrMsg)
	MessageAddFailedErr       = NewErrNo(MessageAddFailedErrCode, MessageAddFailedErrMsg)
	FriendListNoPermissionErr = NewErrNo(FriendListNoPermissionErrCode, FriendListNoPermissionMsg)

	VideoFeedErr              = NewErrNo(VideoFeedErrCode, VideoFeedErrMsg)
	VideoIsNotExistErr        = NewErrNo(VideoIsNotExistErrCode, VideoIsNotExistErrMsg)
	UpdateVideoCountFailedErr = NewErrNo(UpdateVideoCountFailedErrCode, UpdateVideoCountFailedErrMsg)

	CommentActionErr     = NewErrNo(CommentActionErrCode, CommentActionErrMsg)
	CommentIsNotExistErr = NewErrNo(CommentIsNotExistErrCode, CommentIsNotExistErrMsg)
	CommentAddFailedErr  = NewErrNo(CommentAddFailedErrCode, CommentAddFailedErrMsg)
)

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}

//// WriteData Write data into ErrNo
//func WriteData(Err *ErrNo, data interface{}) *ErrNo {
//	Err.Data = data
//	return Err
//}
//
//// WriteDataList Write data list into ErrNo
//func WriteDataList(Err *ErrNo, list response.List) *ErrNo {
//	Err.Data = list
//	return Err
//}
