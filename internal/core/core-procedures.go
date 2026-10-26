package core

import (
	"encoding/json"
	"errors"
	"fastcoin/internal/rpc"
)

func (c *CoreNode) procedures(req rpc.RPCRequest) rpc.RPCResponse {
	var res rpc.RPCResponse
	switch req.Method {
	case "set_tx":
		var tx rpc.Transaction
		bytes, err := json.Marshal(req.Params)
		if err != nil {
			res.NewResponse(false, nil, err.Error())
			return res
		}
		err = json.Unmarshal(bytes, &tx)
		if err != nil {
			res.NewResponse(false, nil, err.Error())
			return res
		}
		err = c.Blockchain.SubmitTransaction(tx)
		if err != nil {
			res.NewResponse(false, nil, err.Error())
			return res
		}
		res.NewResponse(true, nil, "")
	case "get_st":
		var stReq rpc.StatementRequest
		var stRes rpc.StatementResponse
		bytes, err := json.Marshal(req.Params)
		if err != nil {
			res.NewResponse(false, nil, err.Error())
			return res
		}
		err = json.Unmarshal(bytes, &stReq)
		if err != nil {
			res.NewResponse(false, nil, err.Error())
			return res
		}
		bl, txs := c.Blockchain.FetchStatement(stReq.WalletAddress)
		stRes.NewResponse(stReq.WalletAddress, bl, txs)
		res.NewResponse(true, stRes, "")
	case "get_bl":
		var blReq rpc.BalanceRequest
		var blRes rpc.BalanceResponse
		bytes, err := json.Marshal(req.Params)
		if err != nil {
			res.NewResponse(false, nil, err.Error())
			return res
		}
		err = json.Unmarshal(bytes, &blReq)
		if err != nil {
			res.NewResponse(false, nil, err.Error())
			return res
		}
		bl := c.Blockchain.FetchBalance(blReq.WalletAddress)
		blRes.NewResponse(blReq.WalletAddress, bl)
		res.NewResponse(true, blRes, "")
	default:
		err := errors.New("unknown procedure")
		res.NewResponse(false, nil, err.Error())
	}
	return res
}
