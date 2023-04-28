package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/welschmoor/mastaclass/util"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner: util.RandomOwner(),
		Balance: util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}

	acc, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err) //fail the test if err is not nil
	require.NotEmpty(t, acc)

	require.Equal(t, arg.Owner, acc.Owner)
	require.Equal(t, arg.Balance, acc.Balance)
	require.Equal(t, arg.Currency, acc.Currency)

	require.NotZero(t, acc.ID)
	require.NotZero(t, acc.CreatedAt)
	return acc
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	originalAcc := createRandomAccount(t)

	readAcc, err := testQueries.GetAccount(context.Background(), originalAcc.ID)
	require.NoError(t, err)
	require.NotEmpty(t, readAcc)

	require.Equal(t, originalAcc.ID, readAcc.ID)
	require.Equal(t, originalAcc.Owner, readAcc.Owner)
	require.Equal(t, originalAcc.Balance, readAcc.Balance)
	require.Equal(t, originalAcc.Currency, readAcc.Currency)
	require.WithinDuration(t, originalAcc.CreatedAt, readAcc.CreatedAt, time.Second)	//account creating date must be within one second
}

func TestUpdateAccount(t *testing.T) {
  originalAcc :=	createRandomAccount(t)

	arg := UpdateAccountParams {
		ID: originalAcc.ID,
		Balance: util.RandomBalance(),
	}
	updatedAcc, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAcc)
	require.Equal(t, originalAcc.ID, updatedAcc.ID)
	require.Equal(t, originalAcc.Owner, updatedAcc.Owner)
	require.Equal(t, arg.Balance, updatedAcc.Balance) //different!!!
	require.Equal(t, originalAcc.Currency, updatedAcc.Currency)
	require.WithinDuration(t, originalAcc.CreatedAt, updatedAcc.CreatedAt, time.Second)	
}

func TestDeleteAccount(t *testing.T) {
  originalAcc :=	createRandomAccount(t)

  err := testQueries.DeleteAccount(context.Background(), originalAcc.ID)
	require.NoError(t, err)

	deletedAcc, err := testQueries.GetAccount(context.Background(), originalAcc.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, deletedAcc)
}

func TestListAccounts(t *testing.T) {
	for i:= 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit: 5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}



