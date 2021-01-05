package app

import "github.com/Narven/launchpad-manager/src/controllers/tickets"

func mapUrls() {

	v1 := router.Group("/v1")
	{
		v1.POST("/tickets", tickets.Create)
		v1.GET("/tickets", tickets.GetAll)
	}
}
