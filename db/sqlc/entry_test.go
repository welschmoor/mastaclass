package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/welschmoor/simplebank/util"
)

func createRandomEntry(t *testing.T) Entry {
	originalAcc := createRandomAccount(t)
	require.NotEmpty(t, originalAcc)

	arg := CreateEntryParams{
		AccountID: originalAcc.ID,
		Amount: 5000,
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}


func TestGetEntry(t *testing.T) {
  createEntryResponse := createRandomEntry(t)

	getEntryResponse, err := testQueries.GetEntry(context.Background(), createEntryResponse.ID)
	require.NoError(t, err)
	require.NotEmpty(t, getEntryResponse)

	require.Equal(t, createEntryResponse.Amount, getEntryResponse.Amount)
}

func TestGetEntryByUserId(t *testing.T) {
	originalAcc := createRandomAccount(t)
	require.NotEmpty(t, originalAcc)

	arg := CreateEntryParams{
		AccountID: originalAcc.ID,
		Amount: util.RandomBalance(),
	}

	for i :=0; i < 5; i++ {
		entry, err := testQueries.CreateEntry(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, entry)
	}

	arg2 := ListEntriesByUserIdParams{
		AccountID: originalAcc.ID,
		Limit: 10,
		Offset: 0,
	}

  list, err :=	testQueries.ListEntriesByUserId(context.Background(), arg2)
	require.NoError(t, err)
	require.NotEmpty(t, list)
	require.Equal(t, 5, len(list))
}