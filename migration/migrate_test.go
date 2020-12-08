package migration

import (
	"fmt"
	"github.com/yufeifly/proxy/api/types"
	"testing"
)

func TestTryMigrateWithLogging(t *testing.T) {
	reqOpts := MigrateReqOpts{
		Src: types.Address{
			IP:   "127.0.0.1",
			Port: "6789",
		},
		Dst: types.Address{
			IP:   "127.0.0.1",
			Port: "6789",
		},
		ServiceID:     "service.A1",
		ProxyService:  "service1",
		CheckpointID:  "cp-redis",
		CheckpointDir: "/tmp",
	}
	err := TryMigrateWithLogging(reqOpts)
	if err != nil {
		t.Errorf("TryMigrateWithLogging failed, err : %v", err)
	} else {
		fmt.Println("pass")
	}
}
