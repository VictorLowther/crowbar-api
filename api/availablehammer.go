package api

import "github.com/digitalrebar/rebar-api/datatypes"

// AvailableHammer helps track what Hammers are available to be bound to a Node.
type AvailableHammer struct {
	datatypes.AvailableHammer
	Timestamps
	apiHelper
}

func (o *AvailableHammer) hammers() {}

// AvailableHammers returns all of the available Hammers that can be
// bound to a node.
func (c *Client) AvailableHammers() (res []*AvailableHammer, err error) {
	res = make([]*AvailableHammer, 0)
	return res, c.List("available_hammers", &res)
}
