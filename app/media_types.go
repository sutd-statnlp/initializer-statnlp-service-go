// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "intital": Application Media Types
//
// Command:
// $ goagen
// --design=statnlp-initializer-service-go/design
// --out=$(GOPATH)/src/statnlp-initializer-service-go
// --version=v1.3.1

package app

// An message (default view)
//
// Identifier: application/json; view=default
type Message struct {
	// message content
	Content *string `form:"content,omitempty" json:"content,omitempty" xml:"content,omitempty"`
	// message title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
}

// MessageCollection is the media type for an array of Message (default view)
//
// Identifier: application/json; type=collection; view=default
type MessageCollection []*Message