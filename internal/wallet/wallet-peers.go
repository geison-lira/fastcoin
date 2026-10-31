package wallet

import (
	"errors"
	"fastcoin/internal/rpc"
	"time"
)

func (w *WalletNode) activeCore() (string, rpc.RPCConnection, error) {
	var connection rpc.RPCConnection
	var req rpc.RPCRequest
	for _, coreAddress := range w.Peers {
		connection.NewConnection("tcp", coreAddress, 3*time.Second)
		err := req.NewRequest("ping", nil)
		if err != nil {
			continue
		}
		res, err := connection.Send(req)
		if err != nil {
			continue
		}
		if res.Success {
			return coreAddress, connection, nil
		}
	}
	return "", rpc.RPCConnection{}, errors.New("network is down")
}
