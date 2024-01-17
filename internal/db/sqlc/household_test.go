package db

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nicobh15/HomeBuddy-Backend/internal/util"
	"github.com/stretchr/testify/require"
)

func TestCreateHousehold(t *testing.T) {
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

}

func TestFetchHousehold(t *testing.T) {
	parsedID, err := uuid.Parse("5bf1fbcd-b4f1-4c60-a770-9b18a2b2fa9a")
	if err != nil {
		t.Fatalf("Failed to parse UUID: %v", err)
	}
	HouseholdID := pgtype.UUID{Bytes: parsedID, Valid: true}

	household, err := testQueries.FetchHousehold(context.Background(), HouseholdID)
	require.NoError(t, err)
	require.NotEmpty(t, household)
	require.NotEmpty(t, household.HouseholdName)
	require.NotEmpty(t, household.Address)
	require.NotZero(t, household.CreatedAt)
	require.NotZero(t, household.UpdatedAt)
	require.Equal(t, HouseholdID, household.HouseholdID)
}

func TestUpdateHousehold(t *testing.T) {

}

func TestDeleteHousehold(t *testing.T) {

}

func TestListHouseholds(t *testing.T) {

}
