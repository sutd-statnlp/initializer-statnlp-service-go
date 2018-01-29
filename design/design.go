package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("intital", func() {
	Title("The initial service")
	Description("The initial service for StatNLP")
	Scheme("http")
	Host("localhost:8080")
})

var _ = Resource("message", func() {
	BasePath("/api")
	DefaultMedia(MessageMedia)

	Action("list", func() {
		Description("get all messages")
		Routing(GET("/messages"))
		Response(OK, CollectionOf(MessageMedia))
		Response(NotFound)
	})
})

var MessageMedia = MediaType("application/json", func() {
	Description("An message")
	ContentType("application/json")
	TypeName("Message")
	Attributes(func() {
		Attribute("title", String, "message title")
		Attribute("content", String, "message content")
	})
	View("default", func() {
		Attribute("title")
		Attribute("content")
	})
})

var _ = Resource("swagger", func() {
	Description("API docs")
	Files("/api-docs", "swagger/swagger.json")
	Files("/*filepath", "ui/")
})
