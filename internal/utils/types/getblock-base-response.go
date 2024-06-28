package types

type GetBlockBaseResponse[T any] struct {
	Id      string `json:"id"`
	JsonRPC string `json:"jsonrpc"`
	Result  T      `json:"result"`
}
