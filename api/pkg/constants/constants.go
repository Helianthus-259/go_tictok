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
	NoteTableName    = "note"
	UserTableName    = "user"
	SecretKey        = "secret key"
	IdentityKey      = "id"
	Total            = "total"
	Notes            = "notes"
	ApiServiceName   = "demoapi"
	VideoServiceName = "demovideo"
	UserServiceName  = "demouser"
	MySQLDefaultDSN  = "gorm:gorm@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	TCP              = "tcp"
	UserServiceAddr  = ":9000"
	NoteServiceAddr  = ":10000"
	ExportEndpoint   = ":4317"
	ETCDAddress      = "127.0.0.1:2379"
	DefaultLimit     = 10
)
