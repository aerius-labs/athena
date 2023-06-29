package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
)

const TypeMsgSendQueryAllBalances = "send_query_all_balances"

var _ sdk.Msg = &MsgMsgSendQueryAllBalances{}

func NewMsgSendQueryAllBalances(creator, channelId string, addr string, page *query.PageRequest) *MsgMsgSendQueryAllBalances {
	return &MsgMsgSendQueryAllBalances{
		Creator:    creator,
		ChannelId:  channelId,
		Address:    addr,
		Pagination: page,
	}
}

func (msg *MsgMsgSendQueryAllBalances) Route() string {
	return RouterKey
}

func (msg *MsgMsgSendQueryAllBalances) Type() string {
	return TypeMsgSendQueryAllBalances
}

func (msg *MsgMsgSendQueryAllBalances) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMsgSendQueryAllBalances) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMsgSendQueryAllBalances) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
