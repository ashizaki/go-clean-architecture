package controller

import "net/http"

type Context interface {
	Query(string) string
	Param(string) string
	Bind(interface{}) error
	Status(int)
	JSON(int, interface{})
	Request() *http.Request
}
