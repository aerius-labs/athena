package wasmbinding

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	qcrescent "github.com/placeholder-dapps/athena/x/qcrescent/keeper"
)

func RegisterCustomPlugins(qcrescent *qcrescent.Keeper) []wasmkeeper.Option {
	messengerDecoratorOpt := wasmkeeper.WithMessageHandlerDecorator(
		CustomMessageDecorator(qcrescent),
	)
	return []wasm.Option{
		messengerDecoratorOpt,
	}

}
