package migration

import (
	"fmt"
	"github.com/yufeifly/proxy/model"
	"testing"
)

func TestTryMigrateWithLogging(t *testing.T) {
	reqOpts := model.MigrateReqOpts{
		Src: model.Address{
			IP:   "127.0.0.1",
			Port: "6789",
		},
		Dst: model.Address{
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
