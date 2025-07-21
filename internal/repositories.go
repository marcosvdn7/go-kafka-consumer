package internal

import (
	"consumer/cmd/gen/sqlc/sqlc"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MessageLogRepository struct {
	dbConn *pgxpool.Pool
}

func NewMessageLogRepository(dbConn *pgxpool.Pool) *MessageLogRepository {
	return &MessageLogRepository{dbConn: dbConn}
}

func (mlr *MessageLogRepository) InsertLog(ctx context.Context, ip sqlc.InsertLogParams) error {
	q := sqlc.New(mlr.dbConn)
	if _, err := q.InsertLog(ctx, ip); err != nil {
		return err
	}
	
	return nil
}
