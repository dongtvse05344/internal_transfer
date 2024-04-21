package app

import (
	"context"
	"database/sql"

	"github.com/golang/protobuf/proto"

	db "github.com/internal_transfer/dal/db/sqlc"
	"github.com/internal_transfer/pb"
	"github.com/internal_transfer/utils"
)

func (a *App) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	status := utils.GetStatus(nil)
	if req == nil || req.Id == 0 || req.Balance < 0 {
		err := utils.NewInvalidParamsError("invalid params")
		status = utils.GetStatus(err)
		return &pb.CreateAccountResponse{
			Code:    proto.Int64(status.Code),
			Message: proto.String(status.Message),
		}, nil
	}

	_, err := a.Store.CreateAccount(ctx, db.CreateAccountParams{
		ID:      req.Id,
		Balance: req.Balance,
	})

	if err != nil {
		err := utils.NewDbError(err.Error())
		status = utils.GetStatus(err)
	}
	return &pb.CreateAccountResponse{
		Code:    proto.Int64(status.Code),
		Message: proto.String(status.Message),
	}, nil
}

func (a *App) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	status := utils.GetStatus(nil)
	if req == nil || req.GetId() <= 0 {
		err := utils.NewInvalidParamsError("invalid params")
		status = utils.GetStatus(err)
		return &pb.GetAccountResponse{
			Code:    proto.Int64(status.Code),
			Message: proto.String(status.Message),
		}, nil
	}

	account, err := a.Store.GetAccount(ctx, req.GetId())
	if err != nil {
		if err == sql.ErrNoRows {
			err = utils.NewDbRowNotFound(err.Error())
		}
		status = utils.GetStatus(err)
		return &pb.GetAccountResponse{
			Code:    proto.Int64(status.Code),
			Message: proto.String(status.Message),
		}, nil
	}
	return &pb.GetAccountResponse{
		Code:    proto.Int64(status.Code),
		Message: proto.String(status.Message),
		Data: &pb.GetAccountResponseData{
			Account: &pb.Account{
				Id:      account.ID,
				Balance: proto.Float64(account.Balance),
			},
		},
	}, nil
}
