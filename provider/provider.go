package provider

import (
	"github.com/go-redis/redis/v8"
	"go-chat/app/process"
	"go-chat/config"
	"net/http"
)

type Services struct {
	Config     *config.Config
	HttpServer *http.Server
	Process    *process.Process
	Redis      *redis.Client
}
