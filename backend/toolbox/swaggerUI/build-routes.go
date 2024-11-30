package swaggerui

import (
	"fmt"

	"github.com/Tom-Mendy/SentryLink/schemas"
)

func buildRouteEntry(route schemas.Route) map[string]interface{} {
	return map[string]interface{}{
		route.Method: map[string]interface{}{
			"tags":        route.Tags,
			"description": route.Description,
			"produces":    route.Product,
			"parameters":  buildParameters(route.Params, route.ParamQueryType),
			"responses":   buildResponses(route.Responses),
		},
	}
}

func buildParameters(params map[string]string, paramType string) []map[string]interface{} {
	var parameters []map[string]interface{}
	for name, typ := range params {
		parameters = append(parameters, map[string]interface{}{
			"type":        typ,
			"description": typ,
			"name":        name,
			"in":          paramType,
			"required":    true,
		})
	}
	return parameters
}

func buildResponses(responses map[int][]string) map[string]interface{} {
	responseMap := make(map[string]interface{})
	for code, descAndSchema := range responses {
		description := descAndSchema[0]
		schemaRef := descAndSchema[1]

		responseMap[fmt.Sprintf("%d", code)] = map[string]interface{}{
			"description": description,
			"schema": map[string]interface{}{
				"$ref": fmt.Sprintf("#/definitions/%s", schemaRef),
			},
		}
	}
	return responseMap
}
