package client

import (
	"path"

	"github.com/VictorLowther/crowbar-api/datatypes"
)

// Hammer wraps datatypes.Hammer to provide client API functionality
type Hammer struct {
	datatypes.Hammer
	Timestamps
	apiHelper
}

// Hammerer is anything that a Hammer can be bound to.
type Hammerer interface {
	Crudder
	hammers()
}

// Hammers returns all of the Hammers.
func Hammers(scope ...Hammerer) (res []*Hammer, err error) {
	res = make([]*Hammer, 0)
	paths := make([]string, len(scope))
	for i := range scope {
		paths[i] = urlFor(scope[i])
	}
	paths = append(paths, "hammers")
	return res, List(path.Join(paths...), &res)
}
