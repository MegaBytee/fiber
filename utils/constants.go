package utils

const (
	FIBER_SERVER  = "fiber_server"
	FIBER_CACHE   = "fiber_cache"
	FIBER_STORAGE = "fiber_storage"
	SESSIONS      = "sessions"

	CACHE_CONTROL = "Cache-Control"
	NO_CACHE      = "no-cache"

	UA             = "USER-AGENT"
	PLATFORM       = "Sec-CH-UA-Platform"
	X_RECV_WINDOW  = "X-RECV-WINDOW"
	X_KAPI_SIGN    = "X-KAPI-SIGN"
	X_TIMESTAMP    = "X-TIMESTAMP"
	X_AUTH_EXPIRE  = "X-AUTH-EXPIRE"
	X_AUTH_USER    = "X-AUTH-USER" //username
	X_AUTH_KEY     = "X-AUTH-KEY"
	X_USER_SESSION = "x_user_session"
	WEEK_IN_MILLI  = int64(604800000)

	METHOD_GET  = "GET"
	METHOD_POST = "POST"
	DEV_ENV     = true
)
