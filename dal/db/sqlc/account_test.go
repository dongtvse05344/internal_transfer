package db

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	createTestAccount(t)
}

func createTestAccount(t *testing.T) GetAccountRow {
	arg := CreateAccountParams{
		ID:      time.Now().Unix(),
		Balance: float64((rand.Int63n(2000) + time.Now().Unix()) % 10000),
	}
	result, err := testStore.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmptyf(t, result, "result should not be empty")

	account, err := testStore.GetAccount(context.Background(), arg.ID)
	require.NoError(t, err)
	require.NotEmptyf(t, account, "account should not be empty")
	return account
}
