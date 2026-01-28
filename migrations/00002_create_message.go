package migrations

import (
	database "chatAPI/internal/db"
	"chatAPI/internal/model"
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateMessage, downCreateMessage)
}

func upCreateMessage(ctx context.Context, tx *sql.Tx) error {
	return database.DB_MIGRATOR.CreateTable(&model.Message{})
}

func downCreateMessage(ctx context.Context, tx *sql.Tx) error {
	return database.DB_MIGRATOR.DropTable(&model.Message{})
}
