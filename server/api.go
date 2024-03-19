package server

import (
	"fmt"
	"net/http"

	echo "github.com/labstack/echo/v4"

	"realtimemap-go/backend/grains"

	"github.com/asynkron/protoactor-go/cluster"
)

func serveApi(e *echo.Echo, cluster *cluster.Cluster) {

	e.GET("/api/opportunities/:id", func(c echo.Context) error {
		var id string
		if err := echo.PathParamsBinder(c).String("id", &id).BindError(); err != nil {
			c.String(http.StatusBadRequest, "Unable to get id from the request")
		}
		opportunityClient := grains.GetOpportunityGrainClient(cluster, id)
		if grainResponse, err := opportunityClient.GetOpportunity(&grains.GetOpportunityRequest{}); err == nil {
			return c.JSON(http.StatusOK, grainResponse)
		} else {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("Unable to call grain for opportunity %s: %v", id, err))
		}

	})

}
