package redis

import (
	"fmt"
	"github.com/yufeifly/proxy/scheduler"
	"github.com/yufeifly/proxy/ticket"
	"strconv"
	"testing"
)

func TestSet(t *testing.T) {
	proxyService := ""
	service, _ := scheduler.Default().GetService(proxyService)
	for i := 0; i < 50; i++ {
		if i == 5 {
			service.Ticket().Set(ticket.Logging)
		}
		s := strconv.Itoa(i)
		err := Set("", s, s+"#")
		if err != nil {
			fmt.Printf("err: %v\n", err)
		} else {
			fmt.Println("pass")
		}
	}
}
