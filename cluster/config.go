package cluster

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
)

// LoadConfig ...
func LoadConfig() error {
	// fixme using GetWd function is not elegant
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	configFilePath := filepath.Join(dir, "cluster/cluster.json")
	logrus.Debugf("configFilePath: %v", configFilePath)
	jsonFile, err := os.Open(configFilePath)
	if err != nil {
		logrus.Errorf("cluster.LoadConfig open file failed, err: %v", err)
		return err
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &defaultCluster)
	if err != nil {
		logrus.Errorf("cluster.LoadConfig Unmarshal failed, err: %v", err)
		return err
	}
	logrus.Debugf("the cluster: %v", defaultCluster)
	return nil
}
