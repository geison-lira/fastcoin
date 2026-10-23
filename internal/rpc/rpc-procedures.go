package rpc

import (
	"encoding/json"
	"errors"
)

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
		var stRes StatementResponse
		bytes, _ := json.Marshal(req.Params)
		json.Unmarshal(bytes, &stReq)
		bl, txs, err := FetchStatement(stReq.WalletAddress)
		if err != nil {
			res.NewResponse(false, nil, err.Error())
		} else {
			stRes.NewResponse(stReq.WalletAddress, bl, txs)
			res.NewResponse(true, stRes, "")
		}
	case "get_bl":
		var blReq BalanceRequest
		var blRes BalanceResponse
		bytes, _ := json.Marshal(req.Params)
		json.Unmarshal(bytes, &blReq)
		bl, err := FetchBalance(blReq.WalletAddress)
		if err != nil {
			res.NewResponse(false, nil, err.Error())
		} else {
			blRes.NewResponse(blReq.WalletAddress, bl)
			res.NewResponse(true, blRes, "")
		}
	default:
		err := errors.New("unknown procedure")
		res.NewResponse(false, nil, err.Error())
	}
	return res
}
