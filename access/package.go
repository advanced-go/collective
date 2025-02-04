package access

import "github.com/advanced-go/collective/core"

const (
	PkgPath = "github/advanced-go/collective/access"
)

type IAppend struct {
	Thing    func(h core.ErrorHandler, t *core.Thing, content []byte) *core.Status
	Aspect   func(h core.ErrorHandler, t *core.Aspect) *core.Status
	Relation func(h core.ErrorHandler, t *core.Relation) *core.Status
	Frame    func(h core.ErrorHandler, t *core.Frame) *core.Status
}

var Append = func() *IAppend {
	return &IAppend{
		Thing: func(h core.ErrorHandler, t *core.Thing, content []byte) *core.Status {
			return core.StatusOK()
		},
		Aspect: func(h core.ErrorHandler, t *core.Aspect) *core.Status {
			return core.StatusOK()
		},
		Relation: func(h core.ErrorHandler, t *core.Relation) *core.Status {
			return core.StatusOK()
		},
		Frame: func(h core.ErrorHandler, t *core.Frame) *core.Status {
			return core.StatusOK()
		},
	}
}()

type ThingResponse struct {
	Name    core.Urn
	Content []byte
	Status  *core.Status
}

type IRetrieval struct {
	Things func(h core.ErrorHandler, urns []core.Urn) ([]ThingResponse, *core.Status)
}

var Retrieval = func() *IRetrieval {
	return &IRetrieval{
		Things: func(h core.ErrorHandler, urns []core.Urn) ([]ThingResponse, *core.Status) {
			return nil, core.StatusOK()
		},
	}
}()
