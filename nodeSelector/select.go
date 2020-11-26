package nodeSelector

import (
	"github.com/yufeifly/proxy/config"
	"github.com/yufeifly/proxy/model"
)

// BestTarget select the best target for migration
func BestTarget() model.Node {
	return model.Node{
		Address: model.Address{
			IP:   "192.168.227.147",
			Port: config.DefaultMigratorListeningPort,
		},
	}
}
