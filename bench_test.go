package bench

import (
	"testing"
)

var defaultMsg = []byte("send message")

func BenchmarkClientReceive(b *testing.B) {
	b.Logf("Log a simple message without any contexual fields")

	for _, v := range ipcs {
		b.Run(v.name(), func(b *testing.B) {

			out := &blackhole{}
			bm := v.new("client_receive")

			b.ResetTimer()

			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {

					bm.writeServer(defaultMsg)
					//ipc.Sleep()
					msg := bm.readClient()
					if string(msg) == string(defaultMsg) {
						out.Read()
					}
				}
			})

			if out.ReadCount() != uint64(b.N) {
				b.Fatalf(
					"Mismatch in read count. Expected: %d, Actual: %d",
					b.N,
					out.ReadCount(),
				)
			} else {
				b.Logf(
					"Correct read count. Expected: %d, Actual: %d",
					b.N,
					out.ReadCount(),
				)
			}
		})
	}
}

func BenchmarkServerReceive(b *testing.B) {
	b.Logf("Log a simple message without any contexual fields")

	for _, v := range ipcs {
		b.Run(v.name(), func(b *testing.B) {
			out := &blackhole{}
			bm := v.new("server_receive")

			b.ResetTimer()

			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					bm.writeClient(defaultMsg)
					msg := bm.readServer()
					if string(msg) == string(defaultMsg) {
						out.Read()
					}
				}
			})

			if out.ReadCount() != uint64(b.N) {
				b.Fatalf(
					"Mismatch in read count. Expected: %d, Actual: %d",
					b.N,
					out.ReadCount(),
				)
			} else {
				b.Logf(
					"Correct read count. Expected: %d, Actual: %d",
					b.N,
					out.ReadCount(),
				)
			}
		})
	}
}
