package core

const (
	PkgPath = "github/advanced-go/collective/core"
)

// Refactoring is the evolutionary process applied to programming
// We need a way to get the last component of a Urn, as the first n-1 components
// are essentially locators, like the path + resource in a Url.
// So, given Urn how do we get the resource name?

type Urn string
type Uri string
type Timestamp string // Comparison of timestamps must support a temporal ordering

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

// Applications can create as many domains as needed
// "agent" is the reserved domain for the agent collective supporting agent development

const (
	AgentNID = "agent" // Restricted NID/Domain

	ThingNSS    = "thing"    // urn:agent:thing.{module-package}:{type}
	AspectNSS   = "aspect"   // urn:agent:aspect.testing-aspect
	FrameNSS    = "frame"    // urn:agent:frame:testing-frame
	RuleNSS     = "rule"     // urn:agent:rule:testing-rule
	GuidanceNSS = "guidance" // urn:agent:guidance:testing-rule
)

// Identifiers - name for things plus a timestamp supporting a temporal ordering
type Identifiers struct {
	Name    Urn       `json:"name"`
	Created Timestamp `json:"created"`
}

// Thing - something named in the real world, versioned as things can be stateful
type Thing struct {
	Id       Identifiers `json:"id"`
	Resource string      `json:"resource"`
	Version  int         `json:"version"`
}

// Aspect - the way in which something is viewed by the mind, not to be conflated with what the something has
type Aspect struct {
	Id Identifiers `json:"id"`
}

// Relation - relate an aspect to a thing
type Relation struct {
	Created Timestamp `json:"created"`
	Aspect  Urn       `json:"aspect"`
	Thing   Urn       `json:"thing"` // No Urn versioning supported
}

// Frame - list of aspects, versioned as frames are stateful
type Frame struct {
	Id      Identifiers `json:"id"`
	Aspects []Urn       `json:"aspects"`
	Version int         `json:"version"`
}

// Guidance - created from sending text through a NLP API
type Guidance struct {
	Id             Identifiers `json:"id"`
	References     []Urn       `json:"references"` // Maybe this is a relation + a thing??
	PartOfSpeech   string      `json:"part"`
	Lemma          string      `json:"lemma"`
	DependencyEdge string      `json:"dependency"`
}

// This is based on syntax analysis

// Rule - s are part of core, and are a prescribed guide for conduct or action
type Rule struct {
	Id             Identifiers `json:"id"`
	References     []Urn       `json:"references"` // Things that is the target of this rule
	PartOfSpeech   string      `json:"part"`
	Lemma          string      `json:"lemma"`
	DependencyEdge string      `json:"dependency"`
}
