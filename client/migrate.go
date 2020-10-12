package client

import (
	"encoding/json"
	"github.com/levigross/grequests"
	"github.com/yufeifly/proxy/model"
)

func (cli *Client) SendMigrate(opts model.MigrateReqOpts) error {
	mopts := model.MigrateOpts{
		Address:       opts.Dst,
		Container:     opts.Container,
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

	url := "http://127.0.0.1:6789/container/migrate"
	_, err = grequests.Post(url, ro)
	if err != nil {
		return err
	}
	return nil
}
