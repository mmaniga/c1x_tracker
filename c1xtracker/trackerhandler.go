package c1xtracker

import (
	"github.com/mani_clx/c1x_tracker/c1xcore"
	"github.com/unrolled/render"
	"net/http"
)

func trackC(w http.ResponseWriter, r *http.Request) {
	render := render.New()
	//vars := mux.Vars(r)
	res := map[string]string{
		"exchange": "c1x",
		"x":        "X",
		"y":        "Y",
		"z":        "Z",
	}
	successResponse := c1xcore.ConstructResponse(http.StatusOK, "Response : ", res)
	render.JSON(w, http.StatusOK, successResponse)
	return
}
