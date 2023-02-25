package model

import "time"

type TalkRecords struct {
	Id         int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`           // 聊天记录ID
	MsgId      string    `gorm:"column:msg_id;NOT NULL" json:"msg_id"`                     // 消息唯一ID
	Sequence   int64     `gorm:"column:sequence;default:0;NOT NULL" json:"sequence"`       // 消息时序ID
	TalkType   int       `gorm:"column:talk_type;default:1;NOT NULL" json:"talk_type"`     // 对话类型[1:私信;2:群聊;]
	MsgType    int       `gorm:"column:msg_type;default:0;NOT NULL" json:"msg_type"`       // 消息类型[0:系统消息;1:文本消息;2:文件消息;3:会话消息;4:代码消息;5:投票消息;6:群公告;7:好友申请;8:登录通知;9:入群消息/退群消息;]
	UserId     int       `gorm:"column:user_id;default:0;NOT NULL" json:"user_id"`         // 发送者ID（用户ID），0代表系统消息
	ReceiverId int       `gorm:"column:receiver_id;default:0;NOT NULL" json:"receiver_id"` // 接收者ID（用户ID 或 群ID）
	IsRevoke   int       `gorm:"column:is_revoke;default:0;NOT NULL" json:"is_revoke"`     // 是否撤回消息[0:否;1:是;]
	IsMark     int       `gorm:"column:is_mark;default:0;NOT NULL" json:"is_mark"`         // 是否重要消息[0:否;1:是;]
	IsRead     int       `gorm:"column:is_read;default:0;NOT NULL" json:"is_read"`         // 是否已读[0:否;1:是;]
	QuoteId    string    `gorm:"column:quote_id;NOT NULL" json:"quote_id"`                 // 引用消息ID
	Content    string    `gorm:"column:content" json:"content"`                            // 文本消息
	Extra      string    `gorm:"column:extra;default:{}" json:"extra"`                     // 扩展信信息
	CreatedAt  time.Time `gorm:"column:created_at;NOT NULL" json:"created_at"`             // 创建时间
	UpdatedAt  time.Time `gorm:"column:updated_at;NOT NULL" json:"updated_at"`             // 更新时间
}

func (TalkRecords) TableName() string {
	return "talk_records"
}

type TalkRecordExtraCode struct {
	Lang string `json:"lang"`
	Code string `json:"code"`
}

type TalkRecordExtraLocation struct {
	Longitude   string `json:"longitude"`
	Latitude    string `json:"latitude"`
	Description string `json:"description"`
}

type TalkRecordExtraForward struct {
	MsgIds  []int            `json:"msg_ids"` // 消息ID列表
	Records []map[string]any `json:"records"` // 消息列表快照
}

type TalkRecordExtraLogin struct {
	IpAddress string `json:"ip"`
	Address   string `json:"address"`
	Agent     string `json:"agent"`
	CreatedAt string `json:"created_at"`
	Platform  string `json:"platform"`
	Reason    string `json:"reason"`
}

type TalkRecordExtraCard struct {
	UserId int `json:"user_id"`
}

type TalkRecordExtraContact struct {
	Action int    `json:"action"`  // 操作方式 1联系人申请 2被申请人处理结果
	UserId int    `json:"user_id"` // 联系人ID
	Remark string `json:"remark"`  // 申请备注
}

type TalkRecordExtraFile struct {
	Type         int    `json:"type"`
	Drive        int    `json:"drive"`
	Suffix       string `json:"suffix"`
	Size         int    `json:"size"`
	Path         string `json:"path"`
	Url          string `json:"url"`
	OriginalName string `json:"original_name"`
}

type TalkRecordExtraGroupJoin struct {
	Action   int   `json:"action"`  // 操作方式 1邀请入群 2自动退群 3管理员踢人
	Operator int   `json:"user_id"` // 操作人
	Uids     []int `json:"uids"`    // 被操作人
}

type QueryTalkRecordsItem struct {
	Id         int       `json:"id"`
	MsgId      string    `json:"msg_id"`
	Sequence   int64     `json:"sequence"`
	TalkType   int       `json:"talk_type"`
	MsgType    int       `json:"msg_type"`
	UserId     int       `json:"user_id"`
	ReceiverId int       `json:"receiver_id"`
	IsRevoke   int       `json:"is_revoke"`
	IsMark     int       `json:"is_mark"`
	IsRead     int       `json:"is_read"`
	QuoteId    int       `json:"quote_id"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
	Nickname   string    `json:"nickname"`
	Avatar     string    `json:"avatar"`
}
