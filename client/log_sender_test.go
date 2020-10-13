package client

import (
	"encoding/json"
	"fmt"
	"github.com/yufeifly/proxy/model"
	"strconv"
	"testing"
)

func TestClient_SendLog(t *testing.T) {
	logger := model.NewLogger()
	logger.Log.Last = true
	var data []string
	for i := 0; i < 50; i++ {
		s := strconv.Itoa(i)
		tmp := []string{s, s + "#"}
		tmpJson, _ := json.Marshal(tmp)
		data = append(data, string(tmpJson))
	}
	logger.Log.LogQueue = data
	cli := Client{}
	err := cli.SendLog(logger.Log)
	if err != nil {
		t.Errorf("cli.SendLog failed, err : %v", err)
	} else {
		fmt.Println("pass")
	}
}
