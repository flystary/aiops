package receiver

import "github.com/flystary/aiops/modules/transfer/net/receiver/rpc"

func Start() {
	go rpc.StartRpc()
}
