package client

import (
	"encoding/json"
	"github.com/levigross/grequests"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/model"
)

func (cli *Client) SendMigrate(opts model.MigrateReqOpts) error {
	mopts := model.MigrateOpts{
		Address:       opts.Dst,
		Container:     opts.Container,
		ServiceID:     opts.ServiceID,
		CheckpointID:  opts.CheckpointID,
		CheckpointDir: opts.CheckpointDir,
	}

	moptsJson, err := json.Marshal(mopts)
	if err != nil {
		return err
	}

	ro := &grequests.RequestOptions{
		JSON: moptsJson,
	}

	//url := "http://127.0.0.1:6789/container/migrate"
	url := "http://" + opts.Src.IP + ":" + opts.Src.Port + "/container/migrate"
	_, err = grequests.Post(url, ro)
	logrus.Warn("SendMigrate finished")
	if err != nil {
		return err
	}
	return nil
}
