package access

import (
	"github.com/advanced-go/collective/core"
	"net/url"
)

const (
	PkgPath = "github/advanced-go/collective/access"
)

type Urn string
type Tags string
type Resource string
type Timestamp string

// IAppend - append to the collective
// TODO: add tags for context and resemblance query support
//
//	add descriptions for context
type IAppend struct {
	Thing    func(h core.ErrorHandler, nid, path, resource string, content []byte) *core.Status
	Aspect   func(h core.ErrorHandler, nid, resource string) *core.Status
	Frame    func(h core.ErrorHandler, nid, resource string, aspects []Urn) *core.Status
	Relation func(h core.ErrorHandler, aspect, thing Urn) *core.Status
}

var Append = func() *IAppend {
	return &IAppend{
		Thing: func(h core.ErrorHandler, nid, path, resource string, content []byte) *core.Status {
			return core.StatusOK()
		},
		Aspect: func(h core.ErrorHandler, nid, resource string) *core.Status {
			return core.StatusOK()
		},
		Frame: func(h core.ErrorHandler, nid, resource string, aspects []Urn) *core.Status {
			return core.StatusOK()
		},
		Relation: func(h core.ErrorHandler, aspect, thing Urn) *core.Status {
			return core.StatusOK()
		},
	}
}()

// ThingRequest -
type ThingRequest struct {
	Name    Urn
	Version int
}

// ThingResponse -
type ThingResponse struct {
	Name    Urn
	Content []byte
	Status  *core.Status
}

// IRetrieval -
// TODO : add a way to query/search list of aspects
//
//	add a way to query/search a list of frames
type IRetrieval struct {
	Things func(h core.ErrorHandler, urns []ThingRequest) ([]ThingResponse, *core.Status)
	Frame  func(h core.ErrorHandler, name Urn, version int) ([]Urn, *core.Status)
}

var Retrieve = func() *IRetrieval {
	return &IRetrieval{
		Things: func(h core.ErrorHandler, urns []ThingRequest) ([]ThingResponse, *core.Status) {
			return nil, core.StatusOK()
		},
		Frame: func(h core.ErrorHandler, name Urn, version int) ([]Urn, *core.Status) {
			return nil, core.StatusOK()
		},
	}
}()

const (
	RelationNone        = "none"
	RelationDirect      = "direct"
	RelationResemblance = "resemblance"
)

type ThingConstraints interface {
	Urn | Resource | Tags | Timestamp
}

type FilterConstraints interface {
	[]Urn | Timestamp
}

func Relate[T ThingConstraints, U FilterConstraints](h core.ErrorHandler, thing1, thing2 T, filter1, filter2 U, values url.Values) (string, Urn, *core.Status) {
	return RelationNone, "", core.StatusOK()
}

/*
// IRelate -
// TODO : Support type names??
//
//		Can we add name -> type agreement
//		Can we doe some sort of heuristic check to see if the name resembles any existing type Urn??
//	 Modified Urn for things to contain a path and resource, like a Url, and the resource name
//	 of a Thing can now be verified against other aspects.
//
// The thing of type supports
// Resource - resource name
// Urn    - thing or aspect name
// Tags   - list of tag names
type IRelate struct {
	Relate func(h core.ErrorHandler, thing any, aspect Urn, values url.Values) (string, *core.Status)
}

var Retrieve1 = func() *IRelate {
	return &IRelate{
		Relate: func(h core.ErrorHandler, thing any, aspect Urn, values url.Values) (string, *core.Status) {
			return RelationNone, core.StatusOK()
		},
	}
}()


*/
