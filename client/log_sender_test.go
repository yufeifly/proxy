package client

import (
	"encoding/json"
	"fmt"
	"github.com/yufeifly/proxy/api/logger"
	"github.com/yufeifly/proxy/api/types"
	"strconv"
	"testing"
)

func TestClient_SendLog(t *testing.T) {
	proxyService := "service1"
	dataLogger := logger.NewLogger(proxyService)
	dataLogger.Log.Last = true
	var data []string
	for i := 0; i < 50; i++ {
		s := strconv.Itoa(i)
		tmp := []string{s, s + "#"}
		tmpJSON, _ := json.Marshal(tmp)
		data = append(data, string(tmpJSON))
	}
	dataLogger.Log.LogQueue = data
	logWithID := logger.LogWithServiceID{
		Log:            dataLogger.Log,
		ProxyServiceID: proxyService,
	}
	cli := NewClient(types.Address{
		IP:   "127.0.0.1",
		Port: "6789",
	})
	err := cli.SendLog(logWithID)
	if err != nil {
		t.Errorf("cli.SendLog failed, err : %v", err)
	} else {
		fmt.Println("pass")
	}
}
