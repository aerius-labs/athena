package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/golang/protobuf/proto"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgMsgSendQueryAllBalances{}, "qcrescent/MsgSendQueryAllBalances", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMsgSendQueryAllBalances{},
	)
	// this line is used by starport scaffolding # 3

	// Just for using in proto.Print
	registry.RegisterImplementations((*proto.Message)(nil),
		&banktypes.QueryAllBalancesRequest{},
		&banktypes.QueryAllBalancesResponse{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
