package kube_elk

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		TodoIndex,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		TodoCreate,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		TodoShow,
	},
	Route{
		"TodoDelete",
		"DELETE",
		"/todos/{todoId}",
		TodoDelete,
	},
	Route{
		"ELKDelete",
		"DELETE",
		"/elk/{OrgId}",
		ElkDelete,
	},
	Route{
		"ELKCreate",
		"POST",
		"/elk",
		ElkCreate,
	},
	Route{
		"ELKGet",
		"GET",
		"/elk/{OrgId}",
		ElkGet,
	}, Route{
		"ElkServiceList",
		"GET",
		"/elk/{OrgId}/service",
		ElkServiceList,
	},Route{
		"ElkDeployList",
		"GET",
		"/elk/{OrgId}/deploy",
		ElkDeployList,
	},}