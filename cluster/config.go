package cluster

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/model"
	"io/ioutil"
	"os"
	"path/filepath"
)

var DefaultCluster model.Cluster

func init() {
	err := LoadClusterConfig()
	if err != nil {
		logrus.Panicf("cluster.init LoadClusterConfig failed, err: %v", err)
	}
}

func LoadClusterConfig() error {
	// fixme using GetWd function is not elegant
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	configFilePath := filepath.Join(dir, "cluster/cluster.json")
	logrus.Infof("configFilePath: %v", configFilePath)
	jsonFile, err := os.Open(configFilePath)
	if err != nil {
		logrus.Errorf("cluster.LoadClusterConfig open file failed, err: %v", err)
		return err
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &DefaultCluster)
	if err != nil {
		logrus.Errorf("cluster.LoadClusterConfig Unmarshal failed, err: %v", err)
		return err
	}
	//fmt.Printf("the cluster: %v", Cluster)
	return nil
}

func Cluster() *model.Cluster {
	return &DefaultCluster
}
