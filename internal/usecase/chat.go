package usecase

import (
	"context"
	"errors"
	"strings"

	"chatAPI/internal/model"
)

var ErrInvalidTitle = errors.New("invalid title")
var ErrChatNotFound = errors.New("chat not found")

type ChatRepository interface {
	Create(ctx context.Context, chat *model.Chat) error
	GetByID(ctx context.Context, id uint, limit int) (*model.Chat, error)
	DeleteByID(ctx context.Context, id uint) (bool, error)
}

type ChatUsecase struct {
	repo ChatRepository
}

func NewChatUsecase(repo ChatRepository) *ChatUsecase {
	return &ChatUsecase{repo: repo}
}

func (u *ChatUsecase) CreateChat(ctx context.Context, title string) (*model.Chat, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return nil, ErrInvalidTitle
	}

	chat := &model.Chat{
		Title: title,
	}

	if err := u.repo.Create(ctx, chat); err != nil {
		return nil, err
	}

	return chat, nil
}

func (u *ChatUsecase) GetChat(
	ctx context.Context,
	id uint,
	limit int,
) (*model.Chat, error) {

	if limit <= 0 || limit > 100 {
		limit = 20
	}

	chat, err := u.repo.GetByID(ctx, id, limit)
	if err != nil {
		return nil, ErrChatNotFound
	}

	return chat, nil
}

func (u *ChatUsecase) DeleteChat(
	ctx context.Context,
	id uint,
) error {

	ok, err := u.repo.DeleteByID(ctx, id)
	if err != nil {
		return err
	}
	if !ok {
		return ErrChatNotFound
	}

	return nil
}
