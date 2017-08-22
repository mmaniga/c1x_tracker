package c1xcore

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type MetaData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type Response struct {
	Meta MetaData    `json:"meta"`
	Data interface{} `json:"data,omitempty"`
}

func newRouter(subroute string) *mux.Router {

	muxRouter := mux.NewRouter().StrictSlash(true)
	subRouter := muxRouter.PathPrefix(subroute).Subrouter()
	for _, route := range routes {
		subRouter.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return muxRouter
}

type Routes []route

var routes = make(Routes, 0)

func useMiddleware(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}
	return h
}

func contextWrapper(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func AddRoutes(methodName string, methodType string, mRoute string, handlerFunc http.HandlerFunc) {
	r := route{
		Name:        methodName,
		Method:      methodType,
		Pattern:     mRoute,
		HandlerFunc: useMiddleware(handlerFunc, contextWrapper),
	}
	routes = append(routes, r)
}

func Start(port, subroute string) {
	allowedOrigins := handlers.AllowedOrigins([]string{"*"}) // Allowing all origin as of now

	allowedMethods := handlers.AllowedMethods([]string{
		"POST",
		"GET",
		"OPTIONS"})

	http.ListenAndServe(":"+port, handlers.CORS(allowedMethods, allowedOrigins)(newRouter(subroute)))
}

func ConstructResponse(statusCode int, msg string, data interface{}) Response {
	res := Response{}
	res.Meta.Code = statusCode
	res.Meta.Msg = msg
	res.Data = data
	if data != nil {
		fmt.Println("response  : ", res)
	}

	return res
}
