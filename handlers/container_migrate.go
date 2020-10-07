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

// MigrateContainer handler for migrating container
func MigrateContainer(c *gin.Context) {
	Container := c.Query("container")
	CheckpointID := c.Query("checkpointID")
	SrcIP := c.Query("destIP")
	SrcPort := c.Query("srcPort")
	DestIP := c.Query("destIP")
	DestPort := c.Query("destPort")
	CheckpointDir := c.Query("checkpointDir")

	opts := model.MigrateOpts{
		Src: model.Address{
			IP:   SrcIP,
			Port: SrcPort,
		},
		Dst: model.Address{
			IP:   DestIP,
			Port: DestPort,
		},
		Container:     Container,
		CheckpointID:  CheckpointID,
		CheckpointDir: CheckpointDir,
	}
	err := migration.TryMigrate(opts)
	if err != nil {
		utils.ReportErr(c, err)
		logrus.Panic(err)
	}
	//
	c.JSON(200, gin.H{
		"result": "success",
	})
}
