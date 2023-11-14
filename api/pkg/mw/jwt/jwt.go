package jwt

import (
	demo "api/biz/model/user"
	"api/biz/rpc"
	"api/pkg/constants"
	"common-components/errno"
	"context"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"github.com/hertz-contrib/jwt"
	"net/http"
	rpcUser "rpc/kitex_gen/user"
	"time"
)

//	type AuthToken struct {
//		Token string `json:"token" form:"token" query:"token"`
//	}
//
// var Maker token.Maker
//
//	func InitJwt() {
//		Maker = token.NewJWTMaker("cy&july")
//	}
//
// // AuthenticationMiddleware  Token verification middleware HandleFunction
//
//	func AuthenticationMiddleware() app.HandlerFunc {
//		return func(ctx context.Context, c *app.RequestContext) {
//			var authToken AuthToken
//			err := c.BindAndValidate(&authToken)
//			if logger.CheckError(err, "AuthToken BindAndValidate Failed") {
//				c.JSON(http.StatusOK, errno.AuthorizationFailedErr)
//				c.Abort()
//				return
//			}
//			CurrentUserName, CurrentUserId, ok := Maker.VerifyToken(authToken.Token)
//			if !ok {
//				zap.L().Error("TokenVerify err:", zap.Error(errors.New(fmt.Sprintf("illegal token, host:%s", c.Host()))))
//				c.JSON(http.StatusOK, errno.AuthorizationFailedErr)
//				c.Abort()
//				return
//			}
//			c.Set("current_user_id", CurrentUserId)
//			c.Set("current_user_name", CurrentUserName)
//			c.Next(ctx)
//		}
//	}

var JwtMiddleware *jwt.HertzJWTMiddleware

func InitJWT() {
	JwtMiddleware, _ = jwt.New(&jwt.HertzJWTMiddleware{
		Key:           []byte(constants.SecretKey),
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		IdentityKey:   constants.IdentityKey,
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			userid, _ := claims[constants.IdentityKey].(json.Number).Int64()
			return &demo.User{
				ID: userid,
			}
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					constants.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var err error
			var req demo.LoginRequest
			if err = c.BindAndValidate(&req); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			if len(req.Username) == 0 || len(req.Password) == 0 {
				return "", jwt.ErrMissingLoginValues
			}
			return rpc.Login(context.Background(), &rpcUser.LoginRequest{
				Username: req.Username,
				Password: req.Password,
			})
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			id, _ := c.Get("id")
			c.JSON(http.StatusOK, utils.H{
				"status_code": errno.Success.ErrCode,
				"status_msg":  "Success",
				"user_id":     id,
				"token":       token,
				//"expire": expire.Format(time.RFC3339),
			})
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"status_code": errno.AuthorizationFailedErr.ErrCode,
				"status_msg":  message,
				"user_id":     nil,
				"token":       nil,
			})
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			switch t := e.(type) {
			case errno.ErrNo:
				return t.ErrMsg
			default:
				return t.Error()
			}
		},
		ParseOptions: []jwtv4.ParserOption{jwtv4.WithJSONNumber()},
	})
}
