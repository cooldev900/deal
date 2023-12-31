package deal_test

import (
	"testing"

	keepertest "github.com/cooldev900/deal/testutil/keeper"
	"github.com/cooldev900/deal/testutil/nullify"
	"github.com/cooldev900/deal/x/deal"
	"github.com/cooldev900/deal/x/deal/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DealKeeper(t)
	deal.InitGenesis(ctx, *k, genesisState)
	got := deal.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
