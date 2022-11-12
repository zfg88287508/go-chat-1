package service

import (
	"context"
	"time"

	"go-chat/internal/entity"
	"go-chat/internal/pkg/timeutil"
	"go-chat/internal/repository/model"
	"go-chat/internal/repository/repo"
)

type GroupNoticeService struct {
	repo *repo.GroupNotice
}

func NewGroupNoticeService(repo *repo.GroupNotice) *GroupNoticeService {
	return &GroupNoticeService{
		repo: repo,
	}
}

func (s *GroupNoticeService) Dao() *repo.GroupNotice {
	return s.repo
}

type GroupNoticeEditOpt struct {
	UserId    int
	GroupId   int
	NoticeId  int
	Title     string
	Content   string
	IsTop     int
	IsConfirm int
}

// Create 创建群公告
func (s *GroupNoticeService) Create(ctx context.Context, opts *GroupNoticeEditOpt) error {
	return s.repo.Db.Create(&model.GroupNotice{
		GroupId:      opts.GroupId,
		CreatorId:    opts.UserId,
		Title:        opts.Title,
		Content:      opts.Content,
		IsTop:        opts.IsTop,
		IsConfirm:    opts.IsConfirm,
		ConfirmUsers: "{}",
	}).Error
}

// Update 更新群公告
func (s *GroupNoticeService) Update(ctx context.Context, opts *GroupNoticeEditOpt) error {
	_, err := s.repo.BaseUpdate(&model.GroupNotice{}, entity.MapStrAny{
		"id":       opts.NoticeId,
		"group_id": opts.GroupId,
	}, entity.MapStrAny{
		"title":      opts.Title,
		"content":    opts.Content,
		"is_top":     opts.IsTop,
		"is_confirm": opts.IsConfirm,
		"updated_at": time.Now(),
	})

	return err
}

func (s *GroupNoticeService) Delete(ctx context.Context, groupId, noticeId int) error {
	_, err := s.repo.BaseUpdate(&model.GroupNotice{}, entity.MapStrAny{
		"id":       noticeId,
		"group_id": groupId,
	}, entity.MapStrAny{
		"is_delete":  1,
		"deleted_at": timeutil.DateTime(),
	})

	return err
}
