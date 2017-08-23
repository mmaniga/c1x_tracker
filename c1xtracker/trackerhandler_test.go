package c1xtracker

import (
	//"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTracC(t *testing.T) {
	res := httptest.NewRecorder()
	a := assert.New(t)
	req, err := http.NewRequest("GET", "/v2/c",nil)
	trackC(res,req)
	a.NoError(err)
	a.Equal(res.Code, 200)
}