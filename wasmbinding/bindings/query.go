package wasmbinding

import types "github.com/cosmos/cosmos-sdk/codec/types"

type QcrescentQuery struct {
	QueryQueryStateRequest `json:"query_query_state_request"`
}

type QueryQueryStateRequest struct {
	Sequence uint64 `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
}
type QueryQueryStateResponse struct {
	Request  types.Any  `protobuf:"bytes,1,opt,name=request,proto3" json:"request"`
	Response *types.Any `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}
