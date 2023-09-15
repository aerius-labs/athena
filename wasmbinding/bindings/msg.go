package wasmbinding

import (
	"github.com/cosmos/cosmos-sdk/types/query"
)

type QcrescentMsg struct {
	SendQueryAllBalance *SendQueryAllBalance `json:"qcrescent_msg"`
}

type SendQueryAllBalance struct {
	Creator    string             `json:"creator"`
	ChannelId  string             `json:"channel_id"`
	Address    string             `json:"address"`
	Pagination *query.PageRequest `json:"pagination"`
}
