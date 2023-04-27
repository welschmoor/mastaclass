package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateTransfer(t *testing.T) {
  accont1 :=	createRandomAccount(t)
  accont2 :=	createRandomAccount(t)


	arg := CreateTransferParams{
		FromAccountID: accont1.ID,
		ToAccountID: accont2.ID,
		Amount: 1337,
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

}