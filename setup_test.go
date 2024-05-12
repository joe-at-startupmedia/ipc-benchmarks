package bench

import "sync/atomic"

type ipcBenchmark interface {
	new(name string) ipcBenchmark
	name() string
	writeServer(bytes []byte)
	readServer() []byte
	writeClient(bytes []byte)
	readClient() []byte
}

var ipcs = []ipcBenchmark{
	&golangipcBench{},
	&golangipcUnencryptedBench{},
}

type blackhole struct {
	count uint64
}

func (s *blackhole) ReadCount() uint64 {
	return atomic.LoadUint64(&s.count)
}

func (s *blackhole) Read() {
	atomic.AddUint64(&s.count, 1)
}
