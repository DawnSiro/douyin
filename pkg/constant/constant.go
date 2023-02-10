package constant

const (
	CommentTableName            = "comment"
	MessageTableName            = "message"
	RelationTableName           = "relation"
	UserFavoriteVideosTableName = "user_favorite_videos"
	VideoTableName              = "video"
	UserTableName               = "user"
	SecretKey                   = "CloudWeRun"
	IdentityKey                 = "id"
	ApiServiceName              = "api"
	CommentServiceName          = "comment"
	FavoriteServiceName         = "Favorite"
	FeedServiceName             = "feed"
	MessageServiceName          = "message"
	PublishServiceName          = "publish"
	RelationServiceName         = "relation"
	UserServiceName             = "user"
	MySQLDefaultDSN             = "cloud:cloud@tcp(localhost:3306)/douyin?charset=utf8&parseTime=True&loc=Local"
	TCP                         = "tcp"
	UserServiceAddr             = ":30110"
	NoteServiceAddr             = ":30120"
	ExportEndpoint              = ":4317"
	ETCDAddress                 = ":2379"
	DefaultLimit                = 10
)
