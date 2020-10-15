package utils

import (
	"fmt"
	"testing"
)

func TestParseAddress(t *testing.T) {
	addr := "127.0.0.1"
	res, err := ParseAddress(addr)
	if err != nil {
		t.Errorf("ParseAddress failed, err: %v", err)
		return
	}
	fmt.Printf("ip: %v, port: %v\n", res.IP, res.Port)
}
