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
	result, err := a.Store.TransferTx(ctx, db.TransferTxParams{
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
	transferId, _ := result.Transfer.LastInsertId()
	return &pb.TransferResponse{
		Code:       proto.Int64(status.Code),
		Message:    proto.String(status.Message),
		TransferId: proto.Int64(transferId),
	}, nil
}

func (a *App) GetTransfer(ctx context.Context, req *pb.GetTransferRequest) (*pb.GetTransferResponse, error) {
	status := utils.GetStatus(nil)
	if req == nil || req.GetId() <= 0 {
		err := utils.NewInvalidParamsError("invalid params")
		status = utils.GetStatus(err)
		return &pb.GetTransferResponse{
			Code:    proto.Int64(status.Code),
			Message: proto.String(status.Message),
		}, nil
	}

	result, err := a.Store.GetTransferById(ctx, req.GetId())
	if err != nil {
		if err == sql.ErrNoRows {
			err = utils.NewDbRowNotFound("transfer not found")
		}
		status = utils.GetStatus(err)
	}
	return &pb.GetTransferResponse{
		Code:    proto.Int64(status.Code),
		Message: proto.String(status.Message),
		Transfer: &pb.Transfer{
			TransferId:    proto.Int64(result.ID),
			FromAccountId: proto.Int64(result.FromAccountID.Int64),
			ToAccountId:   proto.Int64(result.ToAccountID.Int64),
			Amount:        proto.Float64(result.Amount),
		},
	}, nil
}
