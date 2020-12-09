package nodeSelector

import (
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/api/types"
	"github.com/yufeifly/proxy/cluster"
	"github.com/yufeifly/proxy/config"
)

// BestTarget select the best target for migration
func BestTarget() cluster.Node {
	logrus.Infof("cluster proxy: %v", cluster.DefaultCluster().GetProxy())
	return cluster.Node{
		Address: types.Address{
			IP:   "192.168.227.147",
			Port: config.DefaultMigratorListeningPort,
		},
	}
}
