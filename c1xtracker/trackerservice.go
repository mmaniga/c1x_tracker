package c1xtracker

import (
	"github.com/mani_clx/c1x_tracker/c1xcore"
)

func Start() {
	c1xcore.AddRoutes(
		"trackC",
		"GET",
		"/c",
		trackC,
	)
	c1xcore.Start("8090", "/v1")
}
