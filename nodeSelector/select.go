package nodeSelector

import (
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/api/types"
	"github.com/yufeifly/proxy/cluster"
	"github.com/yufeifly/proxy/config"
)

// BestTarget todo select the best target for migration
func BestTarget() cluster.Node {
	proxyIP := cluster.Default().GetProxy().IP
	logrus.Infof("cluster proxy: %v", proxyIP)
	return cluster.Node{
		Address: types.Address{
			IP:   proxyIP,
			Port: config.DefaultMigratorListeningPort,
		},
	}
}
