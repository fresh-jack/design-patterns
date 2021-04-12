package function

import (
	"fmt"
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	s, err := NewServer("localhost", 1024, Protocol("udp"), Timeout(300*time.Second), MaxConns(100))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}
