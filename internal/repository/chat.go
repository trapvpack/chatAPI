package repository

import (
	"context"

	"chatAPI/internal/model"

	"gorm.io/gorm"
)

type ChatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) *ChatRepository {
	return &ChatRepository{db: db}
}

func (r *ChatRepository) Create(ctx context.Context, chat *model.Chat) error {
	return r.db.WithContext(ctx).Create(chat).Error
}

func (r *ChatRepository) GetByID(
	ctx context.Context,
	id uint,
	limit int,
) (*model.Chat, error) {

	var chat model.Chat
	err := r.db.WithContext(ctx).
		Preload("Messages", func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at desc").Limit(limit)
		}).
		First(&chat, id).
		Error

	if err != nil {
		return nil, err
	}

	return &chat, nil
}

func (r *ChatRepository) DeleteByID(
	ctx context.Context,
	id uint,
) (bool, error) {

	res := r.db.WithContext(ctx).Delete(&model.Chat{}, id)
	if res.Error != nil {
		return false, res.Error
	}

	return res.RowsAffected > 0, nil
}
