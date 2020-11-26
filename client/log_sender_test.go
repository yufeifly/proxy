package client

import (
	"encoding/json"
	"fmt"
	"github.com/yufeifly/proxy/model"
	"strconv"
	"testing"
)

func TestClient_SendLog(t *testing.T) {
	proxyService := "service1"
	logger := model.NewLogger(proxyService)
	logger.Log.Last = true
	var data []string
	for i := 0; i < 50; i++ {
		s := strconv.Itoa(i)
		tmp := []string{s, s + "#"}
		tmpJson, _ := json.Marshal(tmp)
		data = append(data, string(tmpJson))
	}
	logger.Log.LogQueue = data
	logWithID := model.LogWithServiceID{
		Log:            logger.Log,
		ProxyServiceID: proxyService,
	}
	cli := Client{
		Target: model.Address{
			IP:   "127.0.0.1",
			Port: "6789",
		},
	}
	err := cli.SendLog(logWithID)
	if err != nil {
		t.Errorf("cli.SendLog failed, err : %v", err)
	} else {
		fmt.Println("pass")
	}
}
