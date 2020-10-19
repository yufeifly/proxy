/*
proxy container migrate --src=<ip:port> --container=<> [--checkpoint=<> --checkpoint-dir==<>] --dst=<ip:port>
*/
package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/migration"
	"github.com/yufeifly/proxy/model"
	"github.com/yufeifly/proxy/utils"
)

// MigrateContainer handler for redirecting request of migrating container
func MigrateContainer(c *gin.Context) {

	//Container := c.Query("Container")
	ProxyService := c.PostForm("Service") // of proxy
	CheckpointID := c.PostForm("CheckpointID")
	CheckpointDir := c.PostForm("CheckpointDir")
	SrcIP := c.PostForm("SrcIP")
	SrcPort := c.PostForm("srcPort")
	DestIP := c.PostForm("DestIP")
	DestPort := c.PostForm("DestPort")

	opts := model.MigrateReqOpts{
		Src: model.Address{
			IP:   SrcIP,
			Port: SrcPort,
		},
		Dst: model.Address{
			IP:   DestIP,
			Port: DestPort,
		},
		// ServiceID
		ProxyService:  ProxyService,
		CheckpointID:  CheckpointID,
		CheckpointDir: CheckpointDir,
	}

	err := migration.TrySendMigrate(opts)
	if err != nil {
		utils.ReportErr(c, err)
		logrus.Panic(err)
	}
	logrus.Warn("migration.TrySendMigrate finished")
	//
	c.JSON(200, gin.H{"result": "success"})
}
