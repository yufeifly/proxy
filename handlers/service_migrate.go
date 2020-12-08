package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/migration"
	"github.com/yufeifly/proxy/utils"
	"net/http"
)

// MigrateService ...
func MigrateService(c *gin.Context) {
	ProxyService := c.PostForm("Service") // of proxy
	CheckpointID := c.PostForm("CheckpointID")
	CheckpointDir := c.PostForm("CheckpointDir")
	SrcAddr := c.PostForm("Src")
	DstAddr := c.PostForm("Dst")

	src, err := utils.ParseAddress(SrcAddr)
	if err != nil {
		utils.ReportErr(c, http.StatusBadRequest, err)
		logrus.Panic(err)
	}
	dst, err := utils.ParseAddress(DstAddr)
	if err != nil {
		utils.ReportErr(c, http.StatusBadRequest, err)
		logrus.Panic(err)
	}

	opts := migration.MigrateReqOpts{
		Src:           src,
		Dst:           dst,
		ProxyService:  ProxyService,
		CheckpointID:  CheckpointID,
		CheckpointDir: CheckpointDir,
	}

	err = migration.TryMigrateWithLogging(opts)
	if err != nil {
		utils.ReportErr(c, http.StatusInternalServerError, err)
		logrus.Panic(err)
	}
	logrus.Warn("migration.TryMigrateWithLogging finished")
	//
	c.JSON(http.StatusOK, gin.H{"result": "success"})
}
