/*
Package weight implements predicate to give more weight to a route.
*/
package weight

import (
	"net/http"

	"github.com/zalando/skipper/predicates"
	"github.com/zalando/skipper/routing"
)

// The predicate can be referenced in eskip by the name "Weight".
const Name = "Weight"

type (
	spec struct{}

	Predicate struct {
		weight int
	}
)

// New creates a predicate specification, whose instances can be used to give more weight to a route.
//
// Eskip example:
//
// 	Weight(100) -> "https://www.example.org";
//
func New() routing.PredicateSpec { return &spec{} }

func (s *spec) Name() string { return Name }

func (s *spec) Create(args []interface{}) (routing.Predicate, error) {
	if len(args) != 1 {
		return nil, predicates.ErrInvalidPredicateParameters
	}

	if weight, ok := args[0].(float64); ok {
		return &Predicate{int(weight)}, nil
	}

	if weight, ok := args[0].(int); ok {
		return &Predicate{weight}, nil
	}

	return nil, predicates.ErrInvalidPredicateParameters
}

func (p *Predicate) Match(r *http.Request) bool {
	return true
}
