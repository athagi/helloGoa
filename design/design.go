package design

import (
	. "goa.design/goa/dsl"
)

// API describes the global properties of the API server.
var _ = API("calc", func() {
	Title("Calculator Service")
	Description("HTTP service for adding numbers, a goa teaser")
	Server("calc", func() {
		Host("localhost", func() { URI("http://localhost:8088") })
	})
})

// Service describes a service
var _ = Service("calc", func() {
	Description("The calc service performs operations on numbers")
	// Method describes a service method (endpoint)
	Method("add", func() {
		// Payload describes the method payload
		// Here the payload is an object that consists of two fields
		Payload(func() {
			// Attribute describes an object field
			Attribute("a", Int, "Left operand")
			Attribute("b", Int, "Right operand")
			// Both attributes must be provided when invoking "add"
			Required("a", "b")
		})
		// Result describes the method result
		// Here the result is a simple integer value
		Result(Int)
		// HTTP describes the HTTP transport mapping
		HTTP(func() {
			// Requests to the service consist of HTTP GET requests
			// The payload fields are encoded as path parameters
			GET("/add/{a}/{b}")
			// Responses use a "200 OK" HTTP status
			// The result is encoded in the response body
			Response(StatusOK)
		})
	})
})

// var Target = Type("Target", func() {
// 	Attribute("URI", String, "Name of uri.")
// 	Attribute("duration", Int, "duration.")
// 	Attribute("statuscode", Int, "statuscode of uri")
// })

// var _ = Service("ping", func() {
// 	Description("ping service")
// 	// Method describes a service method (endpoint)
// 	Method("create", func() {
// 		// Payload describes the method payload
// 		// Here the payload is an object that consists of two fields
// 		Payload(func() {
// 			Attribute("uri", String)
// 			Attribute("duration", Int)
// 			// Both attributes must be provided when invoking "add"
// 			Required("uri")
// 		})
// 		// HTTP describes the HTTP transport mapping
// 		HTTP(func() {
// 			POST("/")
// 			Body("uri", "duration")
// 		})
// 	})

// 	Method("index", func() {
// 		Description("Index all URIs")
// 		// Payload(ListTargets)
// 		// Payload(ArrayOf(Target))
// 		Result(func() {
// 			Attribute("marker", String, "Pagination marker")
// 			Attribute("Targets", ArrayOf(Target), "list of targets")
// 		})
// 		HTTP(func() {
// 			GET("")
// 			Response(StatusOK, func() {
// 				Header("marker")
// 				Body("Targets")
// 			})
// 		})
// 	})

// 	Method("show", func() {
// 		Payload(func() {
// 			Attribute("uri", String)
// 		})
// 		HTTP(func() {
// 			GET("/")
// 			Body("uri")
// 		})
// 	})

// 	Method("delete", func() {
// 		Payload(func() {
// 			Attribute("uri", String)
// 		})
// 		HTTP(func() {
// 			DELETE("/")
// 			Body("uri")
// 		})
// 	})

// })

var _ = Service("openapi", func() {
	// Serve the file with relative path ../../gen/http/openapi.json for
	// requests sent to /swagger.json.
	Files("/swagger.json", "../../gen/http/openapi.json")
})
