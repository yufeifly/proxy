package utils

import (
	"fmt"
	"testing"
)

func TestNameServiceByProxyService(t *testing.T) {
	name := NameServiceByProxyService("service1")
	fmt.Println(name)
}
