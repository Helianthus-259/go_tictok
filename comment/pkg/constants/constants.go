package constants

const (
	RpcUser        = "rpc-user"
	RpcVideo       = "rpc-video"
	RpcInteraction = "rpc-interaction"
	RpcChat        = "rpc-chat"
	RpcFavor       = "rpc-favor"
	RpcComment     = "rpc-comment"
)

const (
	TokTikApi         = "toktik-api"
	TokTikUser        = "toktik-user"
	TokTikInteraction = "toktik-interaction"
	TokTikVideo       = "toktik-video"
	TokTikChat        = "toktik-chat"
	TokTikFavor       = "toktik-favor"
	TokTikComment     = "toktik-comment"
)

const (
	SentinelApi = "api"
)

const (
	UserTableName       = "user"
	UserCountTableName  = "user_count"
	SecretKey           = "secret key"
	IdentityKey         = "id"
	Total               = "total"
	Notes               = "notes"
	ApiServiceName      = "demoApi"
	VideoServiceName    = "demoVideo"
	UserServiceName     = "demoUser"
	RelationServiceName = "demoRelation"
	MessageServiceName  = "demoMessage"
	FavoriteServiceName = "demoFavorite"
	CommentServiceName  = "demoComment"
	MySQLDefaultDSN     = "root:a123@tcp(127.0.0.1:3306)/tictok_db?charset=utf8mb4&parseTime=True&loc=Local"
	TCP                 = "tcp"
	UserServiceAddr     = "127.0.0.1:9000"
	VideoServiceAddr    = "127.0.0.1:9001"
	RelationServiceAddr = "127.0.0.1:9002"
	MessageServiceAddr  = "127.0.0.1:9003"
	FavoriteServiceAddr = "127.0.0.1:9004"
	CommentServiceAddr  = "127.0.0.1:9004"
	NoteServiceAddr     = ":10000"
	ExportEndpoint      = ":4317"
	ETCDAddress         = "127.0.0.1:2379"
	DefaultLimit        = 10
	RedisDefaultAddr    = "localhost:6379"
)

const (
	FollowCount    = "follow_count"
	FollowerCount  = "follower_count"
	TotalFavorited = "total_favorited"
	WorkCount      = "work_count"
	FavoriteCount  = "favorite_count"
)

const (
	FAVORITE = iota + 1
	CANCELFAVORITE
)
