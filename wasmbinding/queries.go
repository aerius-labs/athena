package wasmbinding

import qcrescentKeeper "github.com/placeholder-dapps/athena/x/qcrescent/keeper"

type QueryPlugin struct {
	qcrescentKeeper *qcrescentKeeper.Keeper
}

// NewQueryPlugin returns a reference to a new QueryPlugin
func NewQueryPlugin(qk *qcrescentKeeper.Keeper) *QueryPlugin {
	return &QueryPlugin{
		qcrescentKeeper: qk,
	}
}
