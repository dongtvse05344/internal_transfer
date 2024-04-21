package db

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTranferTx(t *testing.T) {
	account1 := createTestAccount(t)
	time.Sleep(1 * time.Second)
	account2 := createTestAccount(t)
	t.Log("account1:", account1)
	t.Log("account2:", account2)
	n := 20
	balance1 := account1.Balance
	balance2 := account2.Balance
	amount := float64(10)
	var result TransferTxResult
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		var arg TransferTxParams
		if i%2 == 0 {
			arg = TransferTxParams{
				FromAccountID: account2.ID,
				ToAccountID:   account1.ID,
				Amount:        amount,
			}
		} else {
			arg = TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			}
		}
		wg.Add(1)
		go func(arg TransferTxParams) {
			if arg.FromAccountID == account1.ID {
				t.Log("transfer from account 1 to account 2")
			} else {
				t.Log("transfer from account 2 to account 1")
			}
			var err error
			result, err = testStore.TransferTx(context.Background(), arg)
			require.NoError(t, err)
			require.NotEmpty(t, result)
			wg.Done()
		}(arg)
	}

	wg.Wait()
	dbAccount1, err := testStore.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	dbAccount2, err := testStore.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)
	require.Equal(t, balance1, dbAccount1.Balance)
	require.Equal(t, balance2, dbAccount2.Balance)
}

func TestTranferTx2(t *testing.T) {
	account1 := createTestAccount(t)
	time.Sleep(1 * time.Second)
	account2 := createTestAccount(t)
	t.Log("account1:", account1)
	t.Log("account2:", account2)
	n := 10
	balance1 := account1.Balance
	balance2 := account2.Balance
	amount := float64(10)
	var result TransferTxResult
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		var arg TransferTxParams
		arg = TransferTxParams{
			FromAccountID: account1.ID,
			ToAccountID:   account2.ID,
			Amount:        amount,
		}
		wg.Add(1)
		go func(arg TransferTxParams) {
			var err error
			result, err = testStore.TransferTx(context.Background(), arg)
			require.NoError(t, err)
			require.NotEmpty(t, result)
			wg.Done()
		}(arg)
	}

	wg.Wait()
	dbAccount1, err := testStore.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	dbAccount2, err := testStore.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)
	require.Equal(t, balance1-(float64(n)*amount), dbAccount1.Balance)
	require.Equal(t, balance2+(float64(n)*amount), dbAccount2.Balance)
}
