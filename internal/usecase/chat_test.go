package usecase

import (
	"chatAPI/internal/model"
	"context"
	"testing"
)

type fakeChatRepo struct {
	createdChat *model.Chat
}

func (f *fakeChatRepo) Create(ctx context.Context, chat *model.Chat) error {
	f.createdChat = chat
	chat.ID = 1
	return nil
}

func (f *fakeChatRepo) GetByID(ctx context.Context, id uint, limit int) (*model.Chat, error) {
	return nil, nil
}

func (f *fakeChatRepo) DeleteByID(ctx context.Context, id uint) (bool, error) {
	return false, nil
}

func TestChatUsecase_CreateChat_Success(t *testing.T) {
	repo := &fakeChatRepo{}
	uc := NewChatUsecase(repo)

	chat, err := uc.CreateChat(context.Background(), " Test chat ")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if chat == nil {
		t.Fatal("chat must not be nil")
	}

	if chat.Title != "Test chat" {
		t.Fatalf("expected trimmed title, got %q", chat.Title)
	}

	if repo.createdChat == nil {
		t.Fatal("repository Create was not called")
	}
}

func TestChatUsecase_CreateChat_InvalidTitle(t *testing.T) {
	repo := &fakeChatRepo{}
	uc := NewChatUsecase(repo)

	_, err := uc.CreateChat(context.Background(), "   ")

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err != ErrInvalidTitle {
		t.Fatalf("expected ErrInvalidTitle, got %v", err)
	}

	if repo.createdChat != nil {
		t.Fatal("repository must not be called on invalid input")
	}
}
