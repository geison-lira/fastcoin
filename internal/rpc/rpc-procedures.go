package rpc

import (
	"encoding/json"
	"errors"
	"fmt"
)

func (r *RPCConnection) procedures(req RPCRequest) RPCResponse {
	var res RPCResponse
	switch req.Method {
	case "set_tx":
		var tx Transaction
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
		//err := SubmitTransaction(tx)
		//testing
		fmt.Println("SubTx")
		//end
		if err != nil {
			res.NewResponse(false, nil, err.Error())
			return res
		}
		res.NewResponse(true, nil, "")
	case "get_st":
		var stReq StatementRequest
		var stRes StatementResponse
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
		//bl, txs, err := FetchStatement(stReq.WalletAddress)
		//testing
		fmt.Println("GetBl")
		var bl int64
		var txs []Transaction
		bl = 10000
		txs = append(txs, Transaction{"sys1", "syswall", "3341", 10000, 3, "9900"})
		//end
		if err != nil {
			res.NewResponse(false, nil, err.Error())
			return res
		}
		stRes.NewResponse(stReq.WalletAddress, bl, txs)
		res.NewResponse(true, stRes, "")
	case "get_bl":
		var blReq BalanceRequest
		var blRes BalanceResponse
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
		//bl, err := FetchBalance(blReq.WalletAddress)
		//testing
		var bl int64
		bl = 10000
		//end
		if err != nil {
			res.NewResponse(false, nil, err.Error())
			return res
		}
		blRes.NewResponse(blReq.WalletAddress, bl)
		res.NewResponse(true, blRes, "")
	default:
		err := errors.New("unknown procedure")
		res.NewResponse(false, nil, err.Error())
	}
	return res
}
