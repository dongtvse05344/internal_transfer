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

	n := 100
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
			var err error
			result, err = testStore.TransferTx(context.Background(), arg)
			require.NoError(t, err)
			require.NotEmpty(t, result)
			wg.Done()
		}(arg)
	}
	wg.Wait()
	require.Equal(t, balance1, result.FromAccount.Balance)
	require.Equal(t, balance2, result.ToAccount.Balance)
}

func TestTranferTx2(t *testing.T) {
	account1 := createTestAccount(t)
	time.Sleep(1 * time.Second)
	account2 := createTestAccount(t)

	n := 10
	balance1 := account1.Balance
	balance2 := account2.Balance
	amount := float64(10)
	var result TransferTxResult
	for i := 0; i < n; i++ {
		var arg TransferTxParams
		arg = TransferTxParams{
			FromAccountID: account1.ID,
			ToAccountID:   account2.ID,
			Amount:        amount,
		}

		go func() {
			var err error
			result, err = testStore.TransferTx(context.Background(), arg)
			require.NoError(t, err)
			require.NotEmpty(t, result)
		}()
	}
	require.Equal(t, balance1-100, result.FromAccount.Balance)
	require.Equal(t, balance2+100, result.ToAccount.Balance)
}
