package app

import (
	"context"
	"database/sql"

	db "github.com/internal_transfer/dal/db/sqlc"
	"github.com/internal_transfer/pb"
)

type InitFunc func(app *App) error

type App struct {
	pb.UnimplementedInternalTransferServer

	Db    *sql.DB
	Store *db.Store
}

func (a *App) Close(ctx context.Context) error {
	if a.Db != nil {
		return a.Db.Close()
	}
	return nil
}

func (a *App) Ping(context.Context, *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{
		Message: "[pong]",
	}, nil
}
