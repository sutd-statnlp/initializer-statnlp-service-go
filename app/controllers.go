// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "intital": Application Controllers
//
// Command:
// $ goagen
// --design=statnlp-initializer-service-go/design
// --out=$(GOPATH)/src/statnlp-initializer-service-go
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// MessageController is the controller interface for the Message actions.
type MessageController interface {
	goa.Muxer
	List(*ListMessageContext) error
}

// MountMessageController "mounts" a Message resource controller on the given service.
func MountMessageController(service *goa.Service, ctrl MessageController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListMessageContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/api/messages", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Message", "action", "List", "route", "GET /api/messages")
}

// SwaggerController is the controller interface for the Swagger actions.
type SwaggerController interface {
	goa.Muxer
	goa.FileServer
}

// MountSwaggerController "mounts" a Swagger resource controller on the given service.
func MountSwaggerController(service *goa.Service, ctrl SwaggerController) {
	initService(service)
	var h goa.Handler

	h = ctrl.FileHandler("/api-docs", "swagger/swagger.json")
	service.Mux.Handle("GET", "/api-docs", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger/swagger.json", "route", "GET /api-docs")

	h = ctrl.FileHandler("/*filepath", "ui/")
	service.Mux.Handle("GET", "/*filepath", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "ui/", "route", "GET /*filepath")

	h = ctrl.FileHandler("/", "ui/index.html")
	service.Mux.Handle("GET", "/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "ui/index.html", "route", "GET /")
}