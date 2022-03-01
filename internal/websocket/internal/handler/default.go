package handler

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/tidwall/gjson"

	"go-chat/config"
	"go-chat/internal/cache"
	"go-chat/internal/entity"
	"go-chat/internal/pkg/im"
	"go-chat/internal/pkg/jsonutil"
	"go-chat/internal/pkg/jwt"
	"go-chat/internal/service"
	"go-chat/internal/websocket/internal/dto"
)

type DefaultWebSocket struct {
	rds                *redis.Client
	conf               *config.Config
	cache              *service.ClientService
	room               *cache.Room
	groupMemberService *service.GroupMemberService
}

func NewDefaultWebSocket(
	rds *redis.Client,
	conf *config.Config,
	client *service.ClientService,
	room *cache.Room,
	groupMemberService *service.GroupMemberService,
) *DefaultWebSocket {
	handler := &DefaultWebSocket{rds: rds, conf: conf, cache: client, room: room, groupMemberService: groupMemberService}

	im.Sessions.Default.SetHandler(handler)

	return handler
}

// Connect 初始化连接
func (c *DefaultWebSocket) Connect(ctx *gin.Context) {
	conn, err := im.NewWebsocket(ctx)
	if err != nil {
		log.Printf("websocket connect error: %s", err)
		return
	}

	// 创建客户端
	client := im.NewClient(conn, &im.ClientOptions{
		Channel: im.Sessions.Default,
		Uid:     jwt.GetUid(ctx),
		Storage: c.cache,
	})

	// 推送心跳检测信息
	data, _ := jsonutil.JsonEncodeByte(im.Message{
		Event: "connect",
		Content: gin.H{
			"ping_interval": 20,
			"ping_timeout":  20 * 3,
		},
	})

	_ = client.Write(data)
}

// Open 连接成功回调事件
func (c *DefaultWebSocket) Open(client *im.Client) {
	// 1.查询用户群列表
	ids := c.groupMemberService.Dao().GetUserGroupIds(client.Uid())

	// 2.客户端加入群房间
	for _, id := range ids {
		_ = c.room.Add(context.Background(), &cache.RoomOption{
			Channel:  im.Sessions.Default.Name(),
			RoomType: entity.RoomGroupChat,
			Number:   strconv.Itoa(id),
			Sid:      c.conf.ServerId(),
			Cid:      client.ClientId(),
		})
	}

	// 推送上线消息
	c.rds.Publish(context.Background(), entity.IMGatewayAll, entity.JsonText{
		"event": entity.EventOnlineStatus,
		"data": entity.JsonText{
			"user_id": client.Uid(),
			"status":  1,
		}.Json(),
	}.Json())
}

// Message 消息接收回调事件
func (c *DefaultWebSocket) Message(message *im.ReceiveContent) {
	// fmt.Printf("[%s]消息通知 Client:%d，Content: %s \n", message.Client.Channel().Name(), message.Client.ClientId(), message.Content)

	event := gjson.Get(message.Content, "event").String()

	switch event {
	case "event_keyboard":
		var m *dto.KeyboardMessage
		if err := json.Unmarshal([]byte(message.Content), &m); err == nil {
			c.rds.Publish(context.Background(), entity.IMGatewayAll, entity.JsonText{
				"event": entity.EventKeyboard,
				"data": entity.JsonText{
					"sender_id":   m.Data.SenderID,
					"receiver_id": m.Data.ReceiverID,
				}.Json(),
			}.Json())
		}
	}
}

// Close 客户端关闭回调事件
func (c *DefaultWebSocket) Close(client *im.Client, code int, text string) {
	// 1.判断用户是否是多点登录

	// 2.查询用户群列表
	ids := c.groupMemberService.Dao().GetUserGroupIds(client.Uid())

	// 3.客户端退出群房间
	for _, id := range ids {
		_ = c.room.Del(context.Background(), &cache.RoomOption{
			Channel:  im.Sessions.Default.Name(),
			RoomType: entity.RoomGroupChat,
			Number:   strconv.Itoa(id),
			Sid:      c.conf.ServerId(),
			Cid:      client.ClientId(),
		})
	}

	// 推送下线消息
	c.rds.Publish(context.Background(), entity.IMGatewayAll, entity.JsonText{
		"event": entity.EventOnlineStatus,
		"data": entity.JsonText{
			"user_id": client.Uid(),
			"status":  0,
		}.Json(),
	}.Json())
}
