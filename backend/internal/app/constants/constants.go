package constants

import "time"

var EnvKeys = envKeys{
	Env:               "ENV",
	ServerAddress:     "SERVER_ADDRESS",
	CorsAllowedOrigin: "CORS_ALLOWED_ORIGIN",
	DBDriver:          "DB_DRIVER",
	DBSource:          "DB_SOURCE",
	JWTSecret:         "JWT_SECRET",
}

var Headers = headers{
	Origin:        "Origin",
	ContentLength: "Content-Length",
	Authorization: "Authorization",
	ContentType:   "Content-Type",
}

var MaxAge = 12 * time.Hour

type envKeys struct {
	Env               string
	ServerAddress     string
	CorsAllowedOrigin string
	DBDriver          string
	DBSource          string
	JWTSecret         string
}

type headers struct {
	Origin        string
	ContentLength string
	Authorization string
	ContentType   string
}
