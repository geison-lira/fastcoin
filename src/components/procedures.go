package components

import "encoding/json"

func (r *RPCConnection) procedures(req RPCRequest) RPCResponse {
	var res RPCResponse
	switch req.Method {
	case "set_tx":
		var tx Transaction
		bytes, _ := json.Marshal(req.Params)
		json.Unmarshal(bytes, &tx)
		err := SubmitTransaction(tx)
		if err != nil {
			res.NewResponse(false, nil, err.Error())
		} else {
			res.NewResponse(true, nil, "")
		}
	case "get_st":
		var stReq StatementRequest
		bytes, _ := json.Marshal(req.Params)
		json.Unmarshal(bytes, &stReq)
		bl, txs, err := FetchStatement(stReq.WalletAddress)
		stRes := StatementResponse{WalletAddress: stReq.WalletAddress, Balance: bl, Transactions: txs}
		if err != nil {
			res.NewResponse(false, nil, err.Error())
		} else {
			res.NewResponse(true, stRes, "")
		}
	case "get_bl":
		var blReq BalanceRequest
		bytes, _ := json.Marshal(req.Params)
		json.Unmarshal(bytes, &blReq)
		bl, err := FetchBalance(blReq.WalletAddress)
		blRes := BalanceResponse{WalletAddress: blReq.WalletAddress, Balance: bl}
		if err != nil {
			res.NewResponse(false, nil, err.Error())
		} else {
			res.NewResponse(true, blRes, "")
		}
	default:
		res = RPCResponse{Success: false, Error: "unknown procedure"}
	}
	return res
}
