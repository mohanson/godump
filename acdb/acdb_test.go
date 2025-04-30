package acdb

import (
	"testing"

	"github.com/mohanson/godump/doa"
)

func TestClient(t *testing.T) {
	for _, client := range []*Client{Mem(), Doc(t.TempDir()), Lru(4), Map(t.TempDir())} {
		client.Log(0)
		client.SetEncode("n", 0)
		doa.Doa(doa.Try(client.GetInt("n")) == 0)
		client.SetEncode("s", "Hello World!")
		doa.Doa(doa.Try(client.GetString("s")) == "Hello World!")
	}
}
