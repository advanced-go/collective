package access

import "github.com/advanced-go/collective/core"

const (
	PkgPath = "github/advanced-go/collective/access"
)

const (
	AgentNID = "agent" // Restricted NID/Domain

	ThingNSS    = "thing"    // urn:agent:thing.{module-package}:{type}
	AspectNSS   = "aspect"   // urn:agent:aspect.testing-aspect
	FrameNSS    = "frame"    // urn:agent:frame:testing-frame
	RuleNSS     = "rule"     // urn:agent:rule:testing-rule
	GuidanceNSS = "guidance" // urn:agent:guidance:testing-rule
)

// https://www.rfc-editor.org/rfc/rfc8141#section-2
//   NSS           = pchar *(pchar / "/")

// https://www.rfc-editor.org/rfc/rfc3986
//  pchar         = unreserved / pct-encoded / sub-delims / ":" / "@"
//  unreserved  = ALPHA / DIGIT / "-" / "." / "_" / "~"
//  pct-encoded = "%" HEXDIG HEXDIG
//  sub-delims  = "!" / "$" / "&" / "'" / "(" / ")"
//                  / "*" / "+" / "," / ";" / "="
//  assigned-name = "urn" ":" NID ":" NSS
// Urn syntax : "urn" : NID : NSS
// NID == Domain

// Applications can create as many domains/NISD as needed
// "agent" is the reserved domain for the agent collective supporting agent development

type Urn string

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

// IRetrieval1 -
// TODO : Support type names??
//
//		Can we add name -> type agreement
//		Can we doe some sort of heuristic check to see if the name resembles any existing type Urn??
//	 Modified Urn for things to contain a path and resource, like a Url, and the resource name
//	 of a Thing can now be verified against other aspects.
type IRetrieval1 struct {
	Things func(h core.ErrorHandler, urns []Urn) ([]ThingResponse, *core.Status)
}

var Retrieve1 = func() *IRetrieval1 {
	return &IRetrieval1{
		Things: func(h core.ErrorHandler, urns []Urn) ([]ThingResponse, *core.Status) {
			return nil, core.StatusOK()
		},
	}
}()
