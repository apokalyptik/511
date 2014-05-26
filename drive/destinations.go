package drive

import (
	"errors"
	"time"
)

// Destinations represents the data returned from a 511.org request for destinations
// given a specific origin ID
type Destinations struct {
	OriginNode   int           `xml:"originNode,attr"`
	Timestamp    time.Time     `xml:"timestamp,attr"`
	Destinations []Destination `xml:"destination"`
	client       *Client
}

// Get returns a specific Destination by its Node ID, or an error
func (d Destinations) Get(ID int) (*Destination, error) {
	if d.Destinations == nil {
		return nil, errors.New("no destinations found")
	}
	if len(d.Destinations) == 0 {
		return nil, errors.New("no destinations found")
	}
	for _, v := range d.Destinations {
		if v.Node == ID {
			return &v, nil
		}
	}
	return nil, errors.New("destination ID not found")
}

// Destination represents the data structure of a single destination.  It carries with
// it the Node ID for the origin from whence the list was requested for the purposes of
// the structures Traffic() method.
type Destination struct {
	City      string `xml:"city"`
	MainRoad  string `xml:"mainRoad"`
	CrossRoad string `xml:"crossRoad"`
	Node      int    `xml:"node"`
	client    *Client
	origin    int
}

// Traffic returns the traffic conditions between this Destination Node and its oprigin Node
func (d Destination) Traffic() (*Traffic, error) {
	return d.client.Traffic(d.origin, d.Node)
}
