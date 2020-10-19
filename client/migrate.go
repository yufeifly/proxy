package client

import (
	"encoding/json"
	"github.com/levigross/grequests"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/model"
)

func (cli *Client) SendMigrate(opts model.MigrateReqOpts) error {
	mOpts := model.MigrateOpts{
		Address: opts.Dst,
		//Container:     opts.Container,
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

	//url := "http://127.0.0.1:6789/container/migrate"
	url := "http://" + opts.Src.IP + ":" + opts.Src.Port + "/container/migrate"
	_, err = grequests.Post(url, ro)
	logrus.Warn("SendMigrate finished")
	if err != nil {
		return err
	}
	return nil
}
