package schemas

import "github.com/gin-gonic/gin"


type Route struct {
	Path string
	Method string
	Handler gin.HandlerFunc
	Description string
	Product []string
	Tags []string
	ParamQueryType string
	Params map[string]string
	Responses map[int][]string
}
