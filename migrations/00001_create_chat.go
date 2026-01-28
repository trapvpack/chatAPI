package migrations

import (
	database "chatAPI/internal/db"
	"chatAPI/internal/model"
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateChat, downCreateChat)
}

func upCreateChat(ctx context.Context, tx *sql.Tx) error {
	return database.DB_MIGRATOR.CreateTable(&model.Chat{})
}

func downCreateChat(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return database.DB_MIGRATOR.DropTable(&model.Chat{})
}
