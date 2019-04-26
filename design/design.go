package design

import (
	. "goa.design/goa/dsl"
)

// API describes the global properties of the API server.
var _ = API("ping", func() {
	Title("Ping Service")
	Description("HTTP service for pinging")
	Server("ping", func() {
		Host("localhost", func() { URI("http://localhost:8088") })
	})
})

var Target = Type("Target", func() {
	Attribute("URI", String, "Name of uri.")
	Attribute("duration", Int, "duration.")
	Attribute("statuscode", Int, "statuscode of uri")
})

var _ = Service("ping", func() {
	Description("ping service")
	// Method describes a service method (endpoint)
	// Method("create", func() {
	// 	// Payload describes the method payload
	// 	// Here the payload is an object that consists of two fields
	// 	Payload(func() {
	// 		Attribute("uri", String)
	// 		Attribute("duration", Int)
	// 		// Both attributes must be provided when invoking "add"
	// 		Required("uri")
	// 	})
	// 	// HTTP describes the HTTP transport mapping
	// 	HTTP(func() {
	// 		POST("/")
	// 		Body("uri", "duration")
	// 		Response(StatusCreated)
	// 	})
	// })

	Method("index", func() {
		Description("Index all URIs")
		// Payload(ListTargets)
		// Payload(ArrayOf(Target))
		Result(func() {
			Attribute("marker", String, "Pagination marker")
			Attribute("Targets", ArrayOf(Target), "list of targets")
		})
		HTTP(func() {
			GET("/")
			Response(StatusOK, func() {
				Body("Targets")
			})
		})
	})

	// Method("show", func() {
	// 	Payload(func() {
	// 		Attribute("uri", String)
	// 		Required("uri")
	// 	})
	// 	Result(Target)
	// 	HTTP(func() {
	// 		GET("/")
	// 		Response(StatusOK)
	// 	})
	// })

	// Method("delete", func() {
	// 	Payload(func() {
	// 		Attribute("uri", String)
	// 	})
	// 	HTTP(func() {
	// 		DELETE("/")
	// 		Response(StatusNoContent)
	// 	})
	// })

})
