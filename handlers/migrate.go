/*
Package handlers proxy container migrate --src=<ip:port> --container=<> [--checkpoint=<> --checkpoint-dir==<>] --dst=<ip:port>
*/
package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/migration"
	"github.com/yufeifly/proxy/utils"
	"net/http"
)

// MigrateContainer handler for redirecting request of migrating container
func Migrate(c *gin.Context) {

	service := c.PostForm("Service")
	containerName := c.PostForm("ContainerName")
	srcAddr := c.PostForm("Source")
	dstAddr := c.PostForm("Destination")
	checkpointID := c.PostForm("CheckpointID")
	checkpointDir := c.PostForm("CheckpointDir")

	src, err := utils.ParseAddress(srcAddr)
	if err != nil {
		utils.ReportErr(c, http.StatusBadRequest, err)
		logrus.Panic(err)
	}
	dst, err := utils.ParseAddress(dstAddr)
	if err != nil {
		utils.ReportErr(c, http.StatusBadRequest, err)
		logrus.Panic(err)
	}

	opts := migration.MigrateReqOpts{
		SID:           service,
		CName:         containerName,
		Src:           src,
		Dst:           dst,
		CheckpointID:  checkpointID,
		CheckpointDir: checkpointDir,
	}

	if err = migration.MigrateWithLogging(opts); err != nil {
		utils.ReportErr(c, http.StatusInternalServerError, err)
		logrus.Panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"result": "success"})
}
