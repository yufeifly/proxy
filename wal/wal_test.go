/*
error can't load package: import cycle not allowed
*/
package wal

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
)

// TestSendLastLog pass
func TestSendLastLog(t *testing.T) {

	var data []string
	for i := 0; i < 50; i++ {
		s := strconv.Itoa(i)
		tmp := []string{s, s + "#"}
		tmpJson, _ := json.Marshal(tmp)
		data = append(data, string(tmpJson))
	}
	logger.Log.LogQueue = data
	logger.Log.Last = true

	err := SendLastLog()
	if err != nil {
		t.Errorf("SendLastLog failed: %v\n", err)
	} else {
		fmt.Println("pass")
	}
}
