// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"context"
	"github.com/google/wire"
	"go-chat/app/cache"
	"go-chat/app/dao"
	"go-chat/app/http/handler"
	"go-chat/app/http/handler/api/v1"
	"go-chat/app/http/handler/open"
	"go-chat/app/http/handler/ws"
	"go-chat/app/http/router"
	"go-chat/app/pkg/filesystem"
	"go-chat/app/service"
	"go-chat/config"
	"go-chat/provider"
)

import (
	_ "go-chat/app/validator"
)

// Injectors from wire.go:

func Initialize(ctx context.Context, conf *config.Config) *provider.Services {
	client := provider.RedisConnect(ctx, conf)
	smsCodeCache := &cache.SmsCodeCache{
		Redis: client,
	}
	smsService := service.NewSmsService(smsCodeCache)
	db := provider.MysqlConnect(conf)
	userDao := &dao.UserDao{
		Db: db,
	}
	common := v1.NewCommonHandler(conf, smsService, userDao)
	userService := service.NewUserService(userDao)
	authTokenCache := &cache.AuthTokenCache{
		Redis: client,
	}
	redisLock := &cache.RedisLock{
		Redis: client,
	}
	auth := v1.NewAuthHandler(conf, userService, smsService, authTokenCache, redisLock)
	user := v1.NewUserHandler(userService, smsService)
	talkRecordsDao := &dao.TalkRecordsDao{}
	talkRecordsCodeDao := &dao.TalkRecordsCodeDao{}
	talkRecordsLoginDao := &dao.TalkRecordsLoginDao{}
	talkRecordsFileDao := &dao.TalkRecordsFileDao{}
	talkRecordsVoteDao := &dao.TalkRecordsVoteDao{}
	talkMessageService := service.NewTalkMessageService(conf, talkRecordsDao, talkRecordsCodeDao, talkRecordsLoginDao, talkRecordsFileDao, talkRecordsVoteDao)
	talkMessage := v1.NewTalkMessageHandler(talkMessageService)
	download := v1.NewDownloadHandler()
	filesystemFilesystem := filesystem.NewFilesystem(conf)
	emoticon := v1.NewEmoticonHandler(filesystemFilesystem)
	upload := v1.NewUploadHandler(conf, filesystemFilesystem)
	index := open.NewIndexHandler()
	wsClient := &cache.WsClient{
		Redis: client,
		Conf:  conf,
	}
	clientService := service.NewClientService(wsClient)
	webSocket := ws.NewWebSocketHandler(clientService)
	groupMemberService := service.NewGroupMemberService(db)
	groupService := service.NewGroupService(db, groupMemberService)
	talkListService := service.NewTalkListService(db)
	group := v1.NewGroupHandler(groupService, groupMemberService, talkListService, userDao, redisLock)
	groupNoticeService := service.NewGroupNoticeService(db)
	groupNotice := v1.NewGroupNoticeHandler(groupNoticeService, groupMemberService)
	handlerHandler := &handler.Handler{
		Common:      common,
		Auth:        auth,
		User:        user,
		TalkMessage: talkMessage,
		Download:    download,
		Emoticon:    emoticon,
		Upload:      upload,
		Index:       index,
		Ws:          webSocket,
		Group:       group,
		GroupNotice: groupNotice,
	}
	engine := router.NewRouter(conf, handlerHandler)
	server := provider.NewHttp(conf, engine)
	serverRunID := cache.NewServerRun(client)
	socketService := &service.SocketService{
		Conf:        conf,
		ServerRunID: serverRunID,
	}
	services := &provider.Services{
		HttpServer:   server,
		SocketServer: socketService,
	}
	return services
}

// wire.go:

var providerSet = wire.NewSet(provider.NewLogger, provider.RedisConnect, provider.MysqlConnect, provider.NewHttp, router.NewRouter, filesystem.NewFilesystem, cache.NewServerRun, wire.Struct(new(cache.WsClient), "*"), wire.Struct(new(cache.AuthTokenCache), "*"), wire.Struct(new(cache.SmsCodeCache), "*"), wire.Struct(new(cache.RedisLock), "*"), wire.Struct(new(dao.UserDao), "*"), wire.Struct(new(dao.TalkRecordsDao), "*"), wire.Struct(new(dao.TalkRecordsCodeDao), "*"), wire.Struct(new(dao.TalkRecordsLoginDao), "*"), wire.Struct(new(dao.TalkRecordsFileDao), "*"), wire.Struct(new(dao.TalkRecordsVoteDao), "*"), service.NewUserService, service.NewSmsService, service.NewTalkMessageService, service.NewClientService, service.NewGroupService, service.NewGroupMemberService, service.NewGroupNoticeService, service.NewTalkListService, wire.Struct(new(service.SocketService), "*"), v1.NewAuthHandler, v1.NewCommonHandler, v1.NewUserHandler, v1.NewGroupHandler, v1.NewGroupNoticeHandler, v1.NewTalkHandler, v1.NewTalkMessageHandler, v1.NewUploadHandler, v1.NewDownloadHandler, v1.NewEmoticonHandler, open.NewIndexHandler, ws.NewWebSocketHandler, wire.Struct(new(handler.Handler), "*"), wire.Struct(new(provider.Services), "*"))
