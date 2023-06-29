package types

// IBC events
const (
	EventTypeTimeout     = "timeout"
	EventTypeQueryResult = "query_result"
	// this line is used by starport scaffolding # ibc/packet/event

	AttributeKeyAckSuccess = "success"
	AttributeKeyAck        = "acknowledgement"
	AttributeKeyAckError   = "error"
	AttributeKeySequence   = "sequence"
)
