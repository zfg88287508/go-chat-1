// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/google/wire"
	"go-chat/config"
	"go-chat/internal/provider"
	"go-chat/internal/repository/cache"
	"go-chat/internal/repository/dao"
	"go-chat/internal/service"
	"go-chat/internal/websocket/internal/event"
	"go-chat/internal/websocket/internal/event/chat"
	"go-chat/internal/websocket/internal/handler"
	"go-chat/internal/websocket/internal/process"
	"go-chat/internal/websocket/internal/process/consume"
	"go-chat/internal/websocket/internal/process/server"
	"go-chat/internal/websocket/internal/router"
)

// Injectors from wire.go:

func Initialize(ctx context.Context, conf *config.Config) *AppProvider {
	client := provider.NewRedisClient(ctx, conf)
	sidServer := cache.NewSid(client)
	clientStorage := cache.NewClientStorage(client, conf, sidServer)
	clientService := service.NewClientService(clientStorage)
	roomStorage := cache.NewRoomStorage(client)
	db := provider.NewMySQLClient(conf)
	baseService := service.NewBaseService(db, client)
	baseDao := dao.NewBaseDao(db, client)
	relation := cache.NewRelation(client)
	groupMemberDao := dao.NewGroupMemberDao(baseDao, relation)
	groupMemberService := service.NewGroupMemberService(baseService, groupMemberDao)
	chatHandler := chat.NewHandler(client, groupMemberService)
	defaultEvent := event.NewDefaultEvent(client, conf, clientService, roomStorage, groupMemberService, chatHandler)
	defaultChannel := handler.NewDefaultChannel(clientService, defaultEvent)
	exampleEvent := event.NewExampleEvent()
	exampleChannel := handler.NewExampleChannel(clientService, exampleEvent)
	handlerHandler := &handler.Handler{
		Default: defaultChannel,
		Example: exampleChannel,
	}
	sessionStorage := cache.NewSessionStorage(client)
	engine := router.NewRouter(conf, handlerHandler, sessionStorage)
	websocketServer := provider.NewWebsocketServer(conf, engine)
	healthSubscribe := server.NewHealthSubscribe(conf, sidServer)
	talkVote := cache.NewTalkVote(client)
	talkRecordsVoteDao := dao.NewTalkRecordsVoteDao(baseDao, talkVote)
	talkRecordsDao := dao.NewTalkRecordsDao(baseDao)
	talkRecordsService := service.NewTalkRecordsService(baseService, talkVote, talkRecordsVoteDao, groupMemberDao, talkRecordsDao)
	contactRemark := cache.NewContactRemark(client)
	contactDao := dao.NewContactDao(baseDao, contactRemark, relation)
	contactService := service.NewContactService(baseService, contactDao)
	defaultSubscribe := consume.NewDefaultSubscribe(conf, clientStorage, roomStorage, talkRecordsService, contactService)
	exampleSubscribe := consume.NewExampleSubscribe()
	messageSubscribe := server.NewMessageSubscribe(conf, client, sidServer, defaultSubscribe, exampleSubscribe)
	subServers := &process.SubServers{
		HealthSubscribe:  healthSubscribe,
		MessageSubscribe: messageSubscribe,
	}
	processServer := process.NewServer(subServers)
	appProvider := &AppProvider{
		Config:    conf,
		Server:    websocketServer,
		Coroutine: processServer,
		Handler:   handlerHandler,
	}
	return appProvider
}

// wire.go:

var providerSet = wire.NewSet(provider.NewMySQLClient, provider.NewRedisClient, provider.NewWebsocketServer, router.NewRouter, wire.Struct(new(process.SubServers), "*"), process.NewServer, server.NewHealthSubscribe, server.NewMessageSubscribe, consume.NewDefaultSubscribe, consume.NewExampleSubscribe, cache.NewSessionStorage, cache.NewSid, cache.NewRedisLock, cache.NewClientStorage, cache.NewRoomStorage, cache.NewTalkVote, cache.NewRelation, cache.NewContactRemark, dao.NewBaseDao, dao.NewTalkRecordsDao, dao.NewTalkRecordsVoteDao, dao.NewGroupMemberDao, dao.NewContactDao, chat.NewHandler, event.NewDefaultEvent, event.NewExampleEvent, service.NewBaseService, service.NewTalkRecordsService, service.NewClientService, service.NewGroupMemberService, service.NewContactService, handler.NewDefaultChannel, handler.NewExampleChannel, wire.Struct(new(handler.Handler), "*"), wire.Struct(new(AppProvider), "*"))
