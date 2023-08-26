package db

import (
	"context"
	"testing"


	"github.com/bontusss/gobank/utils"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, from_account, to_account Account) Transfer {
	args := CreatetransferParams {
		FromAccountID: from_account.ID,
		ToAccountID: to_account.ID,
		Amount: utils.RandomBalance(),
	}

	transfer, err := testQueries.Createtransfer(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, args.FromAccountID, transfer.FromAccountID)
	require.Equal(t, args.ToAccountID, transfer.ToAccountID)
	require.NotEmpty(t, transfer.ID)
	require.NotEmpty(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	createRandomTransfer(t, account1, account2)
}

func TestGetTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	tx := createRandomTransfer(t, account1, account2)

	transfer, err := testQueries.GetTransfer(context.Background(), tx.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, tx.CreatedAt, transfer.CreatedAt)
	require.Equal(t, tx.FromAccountID, transfer.FromAccountID)
	require.Equal(t, tx.ToAccountID, transfer.ToAccountID)
	require.Equal(t, tx.ID, transfer.ID)
}

func TestListTransfers(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i:=0;i<10;i++ {
		createRandomTransfer(t, account1, account2)
	}
	args := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID: account2.ID,
		Limit: 5,
		Offset: 5,
	}
	transfers, err := testQueries.ListTransfers(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}