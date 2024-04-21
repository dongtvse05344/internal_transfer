package app

import (
	"context"
	"database/sql"

	"github.com/golang/protobuf/proto"

	db "github.com/internal_transfer/dal/db/sqlc"
	"github.com/internal_transfer/pb"
	"github.com/internal_transfer/utils"
)

func (a *App) Transfer(ctx context.Context, req *pb.TransferRequest) (*pb.TransferResponse, error) {
	status := utils.GetStatus(nil)
	if req == nil || req.FromAccountId <= 0 || req.ToAccountId <= 0 || req.Amount <= 0 {
		err := utils.NewInvalidParamsError("invalid params")
		status = utils.GetStatus(err)
		return &pb.TransferResponse{
			Code:    proto.Int64(status.Code),
			Message: proto.String(status.Message),
		}, nil
	}
	_, err := a.Store.TransferTx(ctx, db.TransferTxParams{
		FromAccountID: req.FromAccountId,
		ToAccountID:   req.ToAccountId,
		Amount:        req.Amount,
	})

	if err != nil {
		if err == sql.ErrNoRows {
			err = utils.NewDbRowNotFound("valid account not found")
		}
		status = utils.GetStatus(err)
	}

	return &pb.TransferResponse{
		Code:    proto.Int64(status.Code),
		Message: proto.String(status.Message),
	}, nil
}
