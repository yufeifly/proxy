package client

import (
	"encoding/json"
	"github.com/levigross/grequests"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/api/types"
)

// SendMigrate send migrate request to dst node
func (cli *Client) SendMigrate(options types.MigrateOpts) error {
	optsJSON, err := json.Marshal(options)
	if err != nil {
		return err
	}

	ro := &grequests.RequestOptions{
		JSON: optsJSON,
	}

	//example url := "http://127.0.0.1:6789/container/migrate"
	url := cli.getAPIPath("/container/migrate")
	_, err = grequests.Post(url, ro)
	logrus.Debug("client.SendMigrate finished")
	if err != nil {
		return err
	}
	return nil
}
