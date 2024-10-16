package phalanx

const (
	Env_Port = "PORT"

	Env_MinioAccessKey       = "MINIO_ACCESSKey"
	Env_MinioSecretAccessKey = "MINIO_SECRET_ACCESS_Key"
	Env_MinioEndPoint        = "MINIO_END_POINT"

	Env_DatabaseClient   = "DB_CLIENT"
	Env_DatabaseUser     = "DB_USER"
	Env_DatabasePassword = "DB_PASSWORD"
	Env_DatabaseHost     = "DB_HOST"
	Env_DatabasePort     = "DB_PORT"
	Env_Database         = "DB"
)
const (
	Defaul_ServerHost          = "localhost"
	Defaul_ServerPort          = 19622
	Default_DBClient           = "postgres"
	Default_TargetDirectory    = "./Videos"
	Default_ThumbnailDirectory = "./Thumbnails"
)

const (
	Param_VideoId = "/:video_id"
)

const (
	Path_Assets     = "/assets"
	Path_Feed       = "/feed"
	Path_Videos     = "/videos"
	Path_Watch      = "/watch"
	Path_Thumbnails = "/thumbnails"
	Path_Thumbnail  = "/thumbnail"
	Path_Search     = "/search"
	Path_API        = "/api"
	// Path_WebSocket  = "/ws"

	Path_Favicon = "/favicon.ico"
)

const (
	Query_VideoId = "v"
)

const (
	Bucket_Thumbnail = "thumbnails"
)

const (
	Key_Offset          = "offset"
	Key_FeedSeed        = "feedSeed"
	Key_VideoId         = "video_id"
	Key_AccessKey       = "accessKey"
	Key_SecretAccessKey = "secretAccessKey"
	Key_Endpoint        = "endpoint"
)
