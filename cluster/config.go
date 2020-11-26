package cluster

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/model"
	"io/ioutil"
	"os"
)

var DefaultCluster model.Cluster

func init() {
	LoadClusterConfig()
}

func LoadClusterConfig() {
	jsonFile, err := os.Open("cluster.json")
	if err != nil {
		logrus.Panicf("cluster.LoadClusterConfig open file failed, err: %v", err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &DefaultCluster)
	if err != nil {
		logrus.Panicf("cluster.LoadClusterConfig Unmarshal failed, err: %v", err)
	}
	//fmt.Printf("the cluster: %v", Cluster)
}

func Cluster() model.Cluster {
	return DefaultCluster
}

func GetProxy() model.Node {
	return Cluster().Master
}

func GetWorkers() []model.Node {
	return Cluster().Worker
}
