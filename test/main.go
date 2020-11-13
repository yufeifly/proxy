package main

import (
	"encoding/json"
	"github.com/levigross/grequests"
	"github.com/sirupsen/logrus"
	"strconv"
	"sync"
	"time"
)

type MigOpts struct {
	Service       string
	CheckpointID  string
	CheckpointDir string
	Src           string
	Dst           string
}

func AccessRedis(wg *sync.WaitGroup) {
	for i := 0; i < 500; i++ {
		data := make(map[string]string, 3)
		data["key"] = "key" + strconv.Itoa(i)
		data["value"] = "value" + strconv.Itoa(i)
		data["service"] = "service1"
		ro := grequests.RequestOptions{
			Data: data,
		}
		url := "http://127.0.0.1:6788/redis/set"
		resp, err := grequests.Post(url, &ro)
		if err != nil {
			logrus.Errorf("AccessRedis.Post err: %v", err)
			continue
		}
		var respStr string
		err = json.NewDecoder(resp.RawResponse.Body).Decode(&respStr)
		if err != nil {
			logrus.Errorf("AccessRedis.Decode err: %v", err)
			continue
		}
		logrus.Infof("resp: %v", respStr)
		time.Sleep(100 * time.Millisecond)
	}
	wg.Done()
}

func TriggerMigration(opts MigOpts) {
	data := make(map[string]string, 5)
	data["Service"] = opts.Service
	data["CheckpointID"] = opts.CheckpointID
	data["CheckpointDir"] = opts.CheckpointDir
	data["Src"] = opts.Src
	data["Dst"] = opts.Dst
	ro := grequests.RequestOptions{
		Data: data,
	}
	url := "http://127.0.0.1:6788/container/migrate"
	resp, err := grequests.Post(url, &ro)
	if err != nil {
		logrus.Errorf("TriggerMigration err: %v", err)
	}
	logrus.Infof("resp: %v", resp.RawResponse.Body)
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go AccessRedis(&wg)

	time.Sleep(2 * time.Second)

	opts := MigOpts{
		Service:       "service1",
		CheckpointID:  "cp-redis",
		CheckpointDir: "/tmp",
		Src:           "127.0.0.1:6789",
		Dst:           "127.0.0.1:6789",
	}
	TriggerMigration(opts)
	wg.Wait()
}
