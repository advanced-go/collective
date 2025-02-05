package access

import "github.com/advanced-go/collective/core"

const (
	PkgPath = "github/advanced-go/collective/access"
)

// IAppend - append to the collective
// TODO: add tags, descriptions
type IAppend struct {
	Thing    func(h core.ErrorHandler, nid, path, resource string, content []byte) *core.Status
	Aspect   func(h core.ErrorHandler, nid, resource string) *core.Status
	Frame    func(h core.ErrorHandler, nid, resource string, aspects []core.Urn) *core.Status
	Relation func(h core.ErrorHandler, aspect, thing core.Urn) *core.Status
}

var Append = func() *IAppend {
	return &IAppend{
		Thing: func(h core.ErrorHandler, nid, path, resource string, content []byte) *core.Status {
			return core.StatusOK()
		},
		Aspect: func(h core.ErrorHandler, nid, resource string) *core.Status {
			return core.StatusOK()
		},
		Frame: func(h core.ErrorHandler, nid, resource string, aspects []core.Urn) *core.Status {
			return core.StatusOK()
		},
		Relation: func(h core.ErrorHandler, aspect, thing core.Urn) *core.Status {
			return core.StatusOK()
		},
	}
}()

// ThingRequest -
type ThingRequest struct {
	Name    core.Urn
	Version int
}

// ThingResponse -
type ThingResponse struct {
	Name    core.Urn
	Content []byte
	Status  *core.Status
}

// IRetrieval -
// TODO : may be add retrievals for aspects and frames.
type IRetrieval struct {
	Things func(h core.ErrorHandler, urns []ThingRequest) ([]ThingResponse, *core.Status)
}

var Retrieve = func() *IRetrieval {
	return &IRetrieval{
		Things: func(h core.ErrorHandler, urns []core.Urn) ([]ThingResponse, *core.Status) {
			return nil, core.StatusOK()
		},
	}
}()

// IRetrieval1 -
// TODO : Support type names??
//
//		Can we add name -> type agreement
//		Can we doe some sort of heuristic check to see if the name resembles any existing type Urn??
//	 Modified Urn for things to contain a path and resource, like a Url, and the resource name
//	 of a Thing can now be verified against other aspects.
type IRetrieval1 struct {
	Things func(h core.ErrorHandler, urns []core.Urn) ([]ThingResponse, *core.Status)
}

var Retrieve1 = func() *IRetrieval1 {
	return &IRetrieval1{
		Things: func(h core.ErrorHandler, urns []core.Urn) ([]ThingResponse, *core.Status) {
			return nil, core.StatusOK()
		},
	}
}()
