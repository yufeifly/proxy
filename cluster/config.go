package cluster

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
)

var defaultCluster Cluster

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
	logrus.Infof("the cluster: %v", defaultCluster)
	return nil
}

// DefaultCluster return default cluster
func DefaultCluster() *Cluster {
	return &defaultCluster
}
