package errno

import (
	"errors"
	"fmt"
)

const (
	SuccessCode    = 0
	ServiceErrCode = iota + 10000
	ParamErrCode
	AuthorizationFailedErrCode

	UserAlreadyExistErrCode
	UserIsNotExistErrCode

	FollowRelationAlreadyExistErrCode
	FollowRelationNotExistErrCode

	FavoriteRelationAlreadyExistErrCode
	FavoriteRelationNotExistErrCode
	FavoriteActionErrCode

	MessageAddFailedErrCode
	FriendListNoPermissionErrCode

	VideoIsNotExistErrCode
	CommentIsNotExistErrCode
)

const (
	SuccessMsg               = "Success"
	ServerErrMsg             = "Service is unable to start successfully"
	ParamErrMsg              = "Wrong Parameter has been given"
	UserIsNotExistErrMsg     = "user is not exist"
	PasswordIsNotVerifiedMsg = "username or password not verified"
	FavoriteActionErrMsg     = "favorite add failed"

	MessageAddFailedErrMsg    = "message add failed"
	FriendListNoPermissionMsg = "You can't query his friend list"
	VideoIsNotExistErrMsg     = "video is not exist"
	CommentIsNotExistErrMsg   = "comment is not exist"
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
	Success                         = NewErrNo(SuccessCode, SuccessMsg)
	ServiceErr                      = NewErrNo(ServiceErrCode, ServerErrMsg)
	ParamErr                        = NewErrNo(ParamErrCode, ParamErrMsg)
	UserAlreadyExistErr             = NewErrNo(UserAlreadyExistErrCode, "User already exists")
	AuthorizationFailedErr          = NewErrNo(AuthorizationFailedErrCode, "Authorization failed")
	UserIsNotExistErr               = NewErrNo(UserIsNotExistErrCode, UserIsNotExistErrMsg)
	PasswordIsNotVerified           = NewErrNo(AuthorizationFailedErrCode, PasswordIsNotVerifiedMsg)
	FollowRelationAlreadyExistErr   = NewErrNo(FollowRelationAlreadyExistErrCode, "Follow Relation already exist")
	FollowRelationNotExistErr       = NewErrNo(FollowRelationNotExistErrCode, "Follow Relation does not exist")
	FavoriteRelationAlreadyExistErr = NewErrNo(FavoriteRelationAlreadyExistErrCode, "Favorite Relation already exist")
	FavoriteRelationNotExistErr     = NewErrNo(FavoriteRelationNotExistErrCode, "FavoriteRelationNotExistErr")
	FavoriteActionErr               = NewErrNo(FavoriteActionErrCode, FavoriteActionErrMsg)

	MessageAddFailedErr       = NewErrNo(MessageAddFailedErrCode, MessageAddFailedErrMsg)
	FriendListNoPermissionErr = NewErrNo(FriendListNoPermissionErrCode, FriendListNoPermissionMsg)
	VideoIsNotExistErr        = NewErrNo(VideoIsNotExistErrCode, VideoIsNotExistErrMsg)
	CommentIsNotExistErr      = NewErrNo(CommentIsNotExistErrCode, CommentIsNotExistErrMsg)
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