package message

import (
	"his6/base/message"
	"testing"
)

func BenchmarkMqsend(b *testing.B) {
	for n := 0; n < b.N; n++ {
		message.Send("slowsql", "content here")
	}
}

func TestMqsend(t *testing.T) {
	message.Send("slowreq", "json here")
	message.Send("slowsql", "json there")
}
