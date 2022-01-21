package types

import "github.com/cosmos/cosmos-sdk/types/query"

// ========================= Address Payload =========================
type AddressPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            Address                `json:"input"`
}

type Address struct {
	Address string `json:"address"`
}

// ========================= Account Balance Payload =========================

type AccountBalancePayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            AccountBalanceArgs     `json:"input"`
}

type AccountBalanceArgs struct {
	Address string `json:"address"`
	Height  int64  `json:"height"`
}

// ========================= Delegation Payload =========================

type DelegationPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            DelegationArgs         `json:"input"`
}

type DelegationArgs struct {
	Address    string             `json:"address"`
	Pagination *query.PageRequest `json:"pagination"`
}