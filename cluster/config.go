package cluster

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
)

var defaultCluster Cluster

func init() {
	err := LoadClusterConfig()
	if err != nil {
		logrus.Panicf("cluster.init LoadClusterConfig failed, err: %v", err)
	}
}

// LoadClusterConfig ...
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
	err = json.Unmarshal(byteValue, &defaultCluster)
	if err != nil {
		logrus.Errorf("cluster.LoadClusterConfig Unmarshal failed, err: %v", err)
		return err
	}
	//fmt.Printf("the cluster: %v", Cluster)
	return nil
}

// Cluster return default cluster
func DefaultCluster() *Cluster {
	return &defaultCluster
}
