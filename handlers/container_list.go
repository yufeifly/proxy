package handlers

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/api/server/httputils"
	"github.com/yufeifly/proxy/container"
	"github.com/yufeifly/proxy/utils"
	"net/http"
	"strconv"
)

// ListContainer handler for redirecting request of listing container(s)
func ListContainer(c *gin.Context) {
	header := "container.List"
	// get list options
	filter, err := filters.FromParam(c.Query("filters"))
	if err != nil {
		logrus.Panic(err)
	}

	listOpts := types.ContainerListOptions{
		All:     httputils.BoolValue(c.Query("all")),
		Size:    httputils.BoolValue(c.Query("size")),
		Since:   c.Query("since"),
		Before:  c.Query("before"),
		Filters: filter,
	}

	if tmpLimit := c.Query("limit"); tmpLimit != "" {
		limit, err := strconv.Atoi(tmpLimit)
		if err != nil {
			logrus.Panic(err)
		}
		listOpts.Limit = limit
	}
	// get target address
	targetAddr, err := utils.ParseAddress(c.Query("Address"))

	listReqOpts := container.ListReqOpts{
		Address:              targetAddr,
		ContainerListOptions: listOpts,
	}

	containers, err := container.ListContainers(listReqOpts)
	if err != nil {
		utils.ReportErr(c, http.StatusInternalServerError, err)
		logrus.Panic(err)
	}

	// return the result
	list := make(gin.H)
	for _, Container := range containers {
		logrus.WithFields(logrus.Fields{
			"ContainerID": Container.ID[:10],
			"Image":       Container.Image,
		}).Infof("%s, List infos", header)
		list[Container.ID[:10]] = Container.Image
	}
	c.JSON(http.StatusOK, list)
}
