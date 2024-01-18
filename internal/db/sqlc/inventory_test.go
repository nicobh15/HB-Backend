package db

import (
	"context"
	"testing"

	"github.com/nicobh15/HomeBuddy-Backend/internal/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomInventoryItem(t *testing.T) Inventory {
	household := CreateRandomHousehold(t)

	args := CreateInventoryItemParams{
		HouseholdID: household.HouseholdID,
		Category:    util.RandomCategory(),
		Name:        util.RandomName(),
		Quantity:    util.RandomInt32(1, 100),
		Location:    util.RandomLocation(),
	}

	inventory, err := testQueries.CreateInventoryItem(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, inventory)
	require.Equal(t, args.HouseholdID, inventory.HouseholdID)
	require.Equal(t, args.Category, inventory.Category)
	require.Equal(t, args.Name, inventory.Name)
	require.Equal(t, args.Quantity, inventory.Quantity)
	require.Equal(t, args.Location, inventory.Location)
	require.NotZero(t, inventory.PurchaseDate)
	require.NotZero(t, inventory.CreatedAt)
	require.NotZero(t, inventory.UpdatedAt)
	require.NotZero(t, inventory.ItemID)

	return inventory
}

func CreateRandomInventoryItems(t *testing.T, num int) []Inventory {
	household := CreateRandomHousehold(t)
	category := util.RandomCategory()
	location := util.RandomLocation()

	var inventory []Inventory
	for i := 0; i < num; i++ {
		args := CreateInventoryItemParams{
			HouseholdID: household.HouseholdID,
			Category:    category,
			Name:        util.RandomName(),
			Quantity:    util.RandomInt32(1, 100),
			Location:    location,
		}
		inventoryItem, err := testQueries.CreateInventoryItem(context.Background(), args)
		require.NoError(t, err)
		require.NotEmpty(t, inventoryItem)
		require.Equal(t, args.HouseholdID, inventoryItem.HouseholdID)
		require.Equal(t, args.Category, inventoryItem.Category)
		require.Equal(t, args.Name, inventoryItem.Name)
		require.Equal(t, args.Quantity, inventoryItem.Quantity)
		require.Equal(t, args.Location, inventoryItem.Location)
		require.NotZero(t, inventoryItem.PurchaseDate)
		require.NotZero(t, inventoryItem.CreatedAt)
		require.NotZero(t, inventoryItem.UpdatedAt)
		require.NotZero(t, inventoryItem.ItemID)
		inventory = append(inventory, inventoryItem)
	}
	return inventory
}

func TestCreateInventory(t *testing.T) {
	CreateRandomInventoryItem(t)
}

func TestDeleteInventoryItem(t *testing.T) {
	item1 := CreateRandomInventoryItem(t)
	_, err := testQueries.DeleteInventoryItem(context.Background(), item1.ItemID)

	item2, err2 := testQueries.FetchInventoryItem(context.Background(), item1.ItemID)

	require.NoError(t, err)
	require.Error(t, err2)
	require.EqualError(t, err2, "no rows in result set")
	require.Empty(t, item2)
}

func TestListInventoryItems(t *testing.T) {

	inventories := CreateRandomInventoryItems(t, 10)
	householdID := inventories[0].HouseholdID

	arg := ListInventoryItemsParams{
		Limit:       5,
		Offset:      5,
		HouseholdID: householdID,
	}

	items, err := testQueries.ListInventoryItems(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, items, 5)

	for _, item := range items {
		require.NotEmpty(t, item)
	}
}

func TestFetchInventoryItem(t *testing.T) {
	item1 := CreateRandomInventoryItem(t)

	item2, err := testQueries.FetchInventoryItem(context.Background(), item1.ItemID)

	require.NoError(t, err)
	require.NotEmpty(t, item2)
	require.Equal(t, item1.ItemID, item2.ItemID)
	require.Equal(t, item1.HouseholdID, item2.HouseholdID)
	require.Equal(t, item1.Category, item2.Category)
	require.Equal(t, item1.Name, item2.Name)
	require.Equal(t, item1.Quantity, item2.Quantity)
	require.Equal(t, item1.Location, item2.Location)
	require.Equal(t, item1.PurchaseDate, item2.PurchaseDate)
	require.Equal(t, item1.CreatedAt, item2.CreatedAt)
	require.Equal(t, item1.UpdatedAt, item2.UpdatedAt)
}

func TestUpdateInventoryItem(t *testing.T) {
	item1 := CreateRandomInventoryItem(t)

	args := UpdateInventoryItemParams{
		Category: util.RandomCategory(),
		Name:     util.RandomName(),
		Quantity: util.RandomInt32(1, 100),
		Location: util.RandomLocation(),
		ItemID:   item1.ItemID,
	}

	item2, err := testQueries.UpdateInventoryItem(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, item2)
	require.Equal(t, args.Category, item2.Category)
	require.Equal(t, args.Name, item2.Name)
	require.Equal(t, args.Quantity, item2.Quantity)
	require.Equal(t, args.Location, item2.Location)
	require.NotZero(t, item2.PurchaseDate)
	require.NotZero(t, item2.CreatedAt)
	require.NotZero(t, item2.UpdatedAt)
	require.Equal(t, item1.ItemID, item2.ItemID)
	require.Equal(t, item1.HouseholdID, item2.HouseholdID)
}

func TestListInventoryItemsByCategory(t *testing.T) {

	inventories := CreateRandomInventoryItems(t, 10)
	householdID := inventories[0].HouseholdID
	category := inventories[0].Category

	arg := ListInventoryItemsByCategoryParams{
		HouseholdID: householdID,
		Category:    category,
		Limit:       5,
		Offset:      5,
	}

	items, err := testQueries.ListInventoryItemsByCategory(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, items, 5)

	for _, item := range items {
		require.NotEmpty(t, item)
	}
}

func TestListInventoryItemsByLocation(t *testing.T) {

	inventories := CreateRandomInventoryItems(t, 10)
	householdID := inventories[0].HouseholdID
	location := inventories[0].Location

	arg := ListInventoryItemsByLocationParams{
		HouseholdID: householdID,
		Location:    location,
		Limit:       5,
		Offset:      5,
	}

	items, err := testQueries.ListInventoryItemsByLocation(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, items, 5)

	for _, item := range items {
		require.NotEmpty(t, item)
	}
}
