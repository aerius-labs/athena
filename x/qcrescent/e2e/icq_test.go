package e2e

import (
	"context"
	"testing"

	"github.com/strangelove-ventures/interchaintest/v7"
	"github.com/strangelove-ventures/interchaintest/v7/ibc"
	"github.com/strangelove-ventures/interchaintest/v7/relayer"
	"github.com/strangelove-ventures/interchaintest/v7/testreporter"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

func TestInterchainQueries(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	t.Parallel()

	client, network := interchaintest.DockerSetup(t)

	rep := testreporter.NewNopReporter()
	eRep := rep.RelayerExecReporter(t)

	ctx := context.Background()

	numVals := 1
	numNodes := 0

	dockerImage := ibc.DockerImage{
		Repository: "ghcr.io/strangelove-ventures/heighliner/icqd",
		Version:    "latest",
		UidGid:     "1025:1025",
	}

	cf := interchaintest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*interchaintest.ChainSpec{
		{
			ChainName:     "controller",
			NumValidators: &numVals,
			NumFullNodes:  &numNodes,
			ChainConfig: ibc.ChainConfig{
				Type:           "cosmos",
				Name:           "controller",
				ChainID:        "controller",
				Images:         []ibc.DockerImage{dockerImage},
				Bin:            "icq",
				Bech32Prefix:   "cosmos",
				Denom:          "atom",
				GasPrices:      "0.00atom",
				TrustingPeriod: "300h",
				GasAdjustment:  1.1,
			}},
		{
			ChainName:     "host",
			NumValidators: &numVals,
			NumFullNodes:  &numNodes,
			ChainConfig: ibc.ChainConfig{
				Type:           "cosmos",
				Name:           "host",
				ChainID:        "host",
				Images:         []ibc.DockerImage{dockerImage},
				Bin:            "icq",
				Bech32Prefix:   "cosmos",
				Denom:          "atom",
				GasPrices:      "0.00atom",
				TrustingPeriod: "300h",
				GasAdjustment:  1.1,
				ModifyGenesis:  modifyGenesisAllowICQQueries([]string{"/cosmos.bank.v1beta1.Query/AllBalances"}), // Add the whitelisted queries to the host chain
			}},
	})

	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)

	controllerChain, hostChain := chains[0], chains[1]

	r := interchaintest.NewBuiltinRelayerFactory(
		ibc.CosmosRly,
		zaptest.NewLogger(t),
		relayer.StartupFlags("-b", "100"),
	).Build(t, client, network)

	const pathName = "host-controller"
	const relayerName = "relayer"

	ic := interchaintest.NewInterchain().
		AddChain(controllerChain).
		AddChain(hostChain).
		AddRelayer(r, relayerName).
		AddLink(interchaintest.InterchainLink{
			Chain1:  controllerChain,
			Chain2:  hostChain,
			Relayer: r,
			Path:    pathName,
			CreateChannelOpts: ibc.CreateChannelOptions{
				SourcePortName: "interquery",
				DestPortName:   "icqhost",
				Order:          ibc.Unordered,
				Version:        "icq-1",
			},
		})

	require.NoError(t, ic.Build(ctx, eRep, interchaintest.InterchainBuildOptions{
		TestName:  t.Name(),
		Client:    client,
		NetworkID: network,

		SkipPathCreation: false,
	}))
}
