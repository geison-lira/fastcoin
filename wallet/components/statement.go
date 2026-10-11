package components

type StatementRequest struct {
	Address string `json:"address"`
}

type StatementResponse struct {
	Address      string        `json:"address"`
	Balance      int64         `json:"balance"`
	Transactions []Transaction `json:"transactions"`
}
