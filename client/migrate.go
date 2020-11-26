package client

import (
	"encoding/json"
	"github.com/levigross/grequests"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/model"
)

func (cli *Client) SendMigrate(opts model.MigrateReqOpts) error {
	mOpts := model.MigrateOpts{
		Address:       opts.Dst,
		ServiceID:     opts.ServiceID,
		ProxyService:  opts.ProxyService,
		CheckpointID:  opts.CheckpointID,
		CheckpointDir: opts.CheckpointDir,
	}

	mOptsJson, err := json.Marshal(mOpts)
	if err != nil {
		return err
	}

	ro := &grequests.RequestOptions{
		JSON: mOptsJson,
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
