package bench

import (
	"fmt"
	ipc "github.com/joe-at-startupmedia/golang-ipc"
)

type golangipcBench struct {
	server *ipc.Server
	client *ipc.Client
}

func (b *golangipcBench) new(name string) ipcBenchmark {
	ipcName := fmt.Sprintf("golangipc_%s", name)
	server, err := ipc.StartServer(&ipc.ServerConfig{Name: ipcName, Encryption: true})
	if err != nil {
		panic(err)
	}

	client, err2 := ipc.StartClient(&ipc.ClientConfig{Name: ipcName, Encryption: true})
	if err2 != nil {
		panic(err)
	}

	return &golangipcBench{
		server: server,
		client: client,
	}
}

func (b *golangipcBench) name() string {
	return "golang-ipc"
}

func (b *golangipcBench) writeServer(bytes []byte) {
	err := b.server.Write(2, bytes)
	if err != nil {
		//log.Printf("writeServer err: %s", err)
	}
}

func (b *golangipcBench) readServer() []byte {
	msg, err := b.server.Read()
	if err != nil {
		//log.Printf("readServer err: %s", err)
	} else if msg.MsgType != 2 {
		//log.Printf("readServer using recursion for msg %d: %s", msg.MsgType, msg.Status)
		return b.readServer()
	}
	return msg.Data
}

func (b *golangipcBench) writeClient(bytes []byte) {
	err := b.client.Write(2, bytes)
	if err != nil {
		//log.Printf("writeClient err: %s", err)
	}
}

func (b *golangipcBench) readClient() []byte {
	msg, err := b.client.Read()
	if err != nil {
		//log.Printf("readClient err: %s", err)
	} else if msg.MsgType != 2 {
		//log.Printf("readClient using recursion for msg %d: %s", msg.MsgType, msg.Status)
		return b.readClient()
	}
	return msg.Data
}
