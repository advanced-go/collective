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

// Append
//

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

// Retrieval

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

type QueryRequest struct {
	NSID string
	NSS  string
}

// IRetrieval -
// TODO : add a way to query/search list of aspects or frames
type IRetrieval struct {
	Things func(h core.ErrorHandler, urns []ThingRequest) ([]ThingResponse, *core.Status)
	Frame  func(h core.ErrorHandler, name Urn, version int) ([]Urn, *core.Status)
	Query  func(h core.ErrorHandler, req QueryRequest) ([]Urn, *core.Status)
}

var Retrieve = func() *IRetrieval {
	return &IRetrieval{
		Things: func(h core.ErrorHandler, urns []ThingRequest) ([]ThingResponse, *core.Status) {
			return nil, core.StatusOK()
		},
		Frame: func(h core.ErrorHandler, name Urn, version int) ([]Urn, *core.Status) {
			return nil, core.StatusOK()
		},
		Query: func(h core.ErrorHandler, req QueryRequest) ([]Urn, *core.Status) {
			return nil, core.StatusOK()
		},
	}
}()

// Relations
//

const (
	RelationNone        = "none"
	RelationDirect      = "direct"
	RelationResemblance = "resemblance"

	VersionName = "ver"
	CompareName = "cmp"
	TopName     = "top"

	ExistValue    = "exist"
	NotExistValue = "!exist"
	LikeValue     = "like"
	NotLikeValue  = "!like"
)

type ThingConstraints interface {
	Urn | Resource | Tags
}

type RelateFilter struct {
	Urns      []Urn
	Resources []Resource
	Tags      Tags
	From      Timestamp
	To        Timestamp
}

func Relate[T ThingConstraints](h core.ErrorHandler, thing1, thing2 T, filter1, filter2 RelateFilter, values url.Values) (string, *core.Status) {
	return RelationNone, core.StatusOK()
}
