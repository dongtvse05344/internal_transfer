package app

import (
	"context"
	"database/sql"
	"net/http"

	db "github.com/internal_transfer/dal/db/sqlc"
	"github.com/internal_transfer/pb"
)

type InitFunc func(app *App) error

type App struct {
	pb.UnimplementedInternalTransferServer
	*http.Server

	Db    *sql.DB
	Store *db.Store
}

func (a *App) Ping(context.Context, *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{
		Message: "[pong]",
	}, nil
}
