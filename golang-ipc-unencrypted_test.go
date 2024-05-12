package bench

import (
	"fmt"
	ipc "github.com/joe-at-startupmedia/golang-ipc"
)

type golangipcUnencryptedBench struct {
	golangipcBench
}

func (b *golangipcUnencryptedBench) new(name string) ipcBenchmark {
	ipcName := fmt.Sprintf("golangipc_unecrypted_%s", name)
	server, err := ipc.StartServer(&ipc.ServerConfig{Name: ipcName, Encryption: false})
	if err != nil {
		panic(err)
	}

	client, err2 := ipc.StartClient(&ipc.ClientConfig{Name: ipcName, Encryption: false})
	if err2 != nil {
		panic(err)
	}

	return &golangipcBench{
		server: server,
		client: client,
	}
}

func (b *golangipcUnencryptedBench) name() string {
	return "golang-ipc-unencrypted"
}
