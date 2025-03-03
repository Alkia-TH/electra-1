package keeper_test

import (
	"strconv"
	"testing"

	keepertest "electra/testutil/keeper"
	"electra/testutil/nullify"
	"electra/x/meter/keeper"
	"electra/x/meter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNCustomerbillingline(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Customerbillingline {
	items := make([]types.Customerbillingline, n)
	for i := range items {
		items[i].CustomerDeviceID = strconv.Itoa(i)
		items[i].CycleID = uint64(i)
		items[i].Lineid = uint64(i)

		keeper.SetCustomerbillingline(ctx, items[i])
	}
	return items
}

func TestCustomerbillinglineGet(t *testing.T) {
	keeper, ctx := keepertest.MeterKeeper(t)
	items := createNCustomerbillingline(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetCustomerbillingline(ctx,
			item.CustomerDeviceID,
			item.CycleID,
			item.Lineid,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestCustomerbillinglineRemove(t *testing.T) {
	keeper, ctx := keepertest.MeterKeeper(t)
	items := createNCustomerbillingline(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCustomerbillingline(ctx,
			item.CustomerDeviceID,
			item.CycleID,
			item.Lineid,
		)
		_, found := keeper.GetCustomerbillingline(ctx,
			item.CustomerDeviceID,
			item.CycleID,
			item.Lineid,
		)
		require.False(t, found)
	}
}

func TestCustomerbillinglineGetAll(t *testing.T) {
	keeper, ctx := keepertest.MeterKeeper(t)
	items := createNCustomerbillingline(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllCustomerbillingline(ctx)),
	)
}
