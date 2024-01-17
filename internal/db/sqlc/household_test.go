package db

import (
	"context"
	"testing"

	"github.com/nicobh15/HomeBuddy-Backend/internal/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomHousehold(t *testing.T) Household {
	args := CreateHouseholdParams{
		HouseholdName: util.RandomName(),
		Address:       util.RandomAddress(),
	}

	household, err := testQueries.CreateHousehold(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, household)
	require.Equal(t, args.HouseholdName, household.HouseholdName)
	require.Equal(t, args.Address, household.Address)
	require.NotZero(t, household.CreatedAt)
	require.NotZero(t, household.UpdatedAt)
	require.NotZero(t, household.HouseholdID)

	return household
}

func TestCreateHousehold(t *testing.T) {
	CreateRandomHousehold(t)
}

func TestFetchHousehold(t *testing.T) {
	household1 := CreateRandomHousehold(t)

	household2, err := testQueries.FetchHousehold(context.Background(), household1.HouseholdID)
	require.NoError(t, err)
	require.NotEmpty(t, household2)
	require.NotEmpty(t, household2.HouseholdName)
	require.NotEmpty(t, household2.Address)
	require.NotZero(t, household2.CreatedAt)
	require.NotZero(t, household2.UpdatedAt)
	require.Equal(t, household1.HouseholdID, household2.HouseholdID)
}

func TestUpdateHousehold(t *testing.T) {
	household1 := CreateRandomHousehold(t)
	UpdateHouseholdParams := UpdateHouseholdParams{util.RandomName(), util.RandomAddress(), household1.HouseholdID}
	household2, err := testQueries.UpdateHousehold(context.Background(), UpdateHouseholdParams)
	require.NoError(t, err)
	require.NotEmpty(t, household2)
	require.NotEmpty(t, household2.HouseholdName)
	require.NotEmpty(t, household2.Address)
	require.NotZero(t, household2.CreatedAt)
	require.NotZero(t, household2.UpdatedAt)
	require.NotEqual(t, household2.HouseholdName, household1.HouseholdName)
	require.NotEqual(t, household2.Address, household1.Address)
	require.Equal(t, household1.HouseholdID, household2.HouseholdID)

}

func TestDeleteHousehold(t *testing.T) {
	household1 := CreateRandomHousehold(t)
	household2, err := testQueries.DeleteHousehold(context.Background(), household1.HouseholdID)

	require.NoError(t, err)
	require.NotEmpty(t, household2)
	require.NotEmpty(t, household2.HouseholdName)
	require.NotEmpty(t, household2.Address)
	require.NotZero(t, household2.CreatedAt)
	require.NotZero(t, household2.UpdatedAt)
	require.Equal(t, household1.HouseholdID, household2.HouseholdID)
	require.Equal(t, household1.HouseholdName, household2.HouseholdName)
	require.Equal(t, household1.Address, household2.Address)
}

func TestListHouseholds(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomHousehold(t)
	}

	arg := ListHouseholdsParams{
		Limit:  5,
		Offset: 5,
	}

	households, err := testQueries.ListHouseholds(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, households, 5)

	for _, household := range households {
		require.NotEmpty(t, household)
		require.NotEmpty(t, household.HouseholdName)
		require.NotEmpty(t, household.Address)
		require.NotZero(t, household.CreatedAt)
		require.NotZero(t, household.UpdatedAt)
		require.NotZero(t, household.HouseholdID)
	}

}
