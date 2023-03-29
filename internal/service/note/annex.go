package note

import (
	"context"

	"go-chat/internal/pkg/filesystem"
	"go-chat/internal/pkg/timeutil"
	"go-chat/internal/repository/model"
	"go-chat/internal/repository/repo"
	"go-chat/internal/repository/repo/note"
)

type ArticleAnnexService struct {
	*repo.Source
	annex      *note.ArticleAnnex
	fileSystem *filesystem.Filesystem
}

func NewArticleAnnexService(source *repo.Source, dao *note.ArticleAnnex, fileSystem *filesystem.Filesystem) *ArticleAnnexService {
	return &ArticleAnnexService{Source: source, annex: dao, fileSystem: fileSystem}
}

func (s *ArticleAnnexService) Dao() *note.ArticleAnnex {
	return s.annex
}

func (s *ArticleAnnexService) Create(ctx context.Context, data *model.ArticleAnnex) error {
	return s.annex.Create(ctx, data)
}

// UpdateStatus 更新附件状态
func (s *ArticleAnnexService) UpdateStatus(ctx context.Context, uid int, id int, status int) error {

	data := map[string]any{
		"status": status,
	}

	if status == 2 {
		data["deleted_at"] = timeutil.DateTime()
	}

	return s.Db().Model(&model.ArticleAnnex{}).Where("id = ? and user_id = ?", id, uid).Updates(data).Error
}

// ForeverDelete 永久删除笔记附件
func (s *ArticleAnnexService) ForeverDelete(ctx context.Context, uid int, id int) error {

	annex, err := s.annex.FindByWhere(ctx, "id = ? and user_id = ?", id, uid)
	if err != nil {
		return err
	}

	switch annex.Drive {
	case 1:
		_ = s.fileSystem.Local.Delete(annex.Path)
	case 2:
		_ = s.fileSystem.Cos.Delete(annex.Path)
	}

	return s.Db().Delete(&model.ArticleAnnex{}, id).Error
}
