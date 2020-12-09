package client

import (
	"encoding/json"
	"github.com/levigross/grequests"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/api/types"
	"github.com/yufeifly/proxy/model"
)

// SendMigrate ...
func (cli *Client) SendMigrate(options types.MigrateOpts) error {
	mOpts := model.MigrateOpts{
		Address:       options.Address,
		ServiceID:     options.ServiceID,
		ProxyService:  options.ProxyService,
		CheckpointID:  options.CheckpointID,
		CheckpointDir: options.CheckpointDir,
	}

	mOptsJSON, err := json.Marshal(mOpts)
	if err != nil {
		return err
	}

	ro := &grequests.RequestOptions{
		JSON: mOptsJSON,
	}

	//example url := "http://127.0.0.1:6789/container/migrate"
	url := cli.getAPIPath("/container/migrate")
	_, err = grequests.Post(url, ro)
	logrus.Debug("SendMigrate finished")
	if err != nil {
		return err
	}
	return nil
}
