package drive

import (
	"errors"
	"time"
)

// Origins represents the response structure returned when requesting a list of origins from 511.org
type Origins struct {
	CacheStamp time.Time `xml:"cachestamp,attr"`
	Origins    []Origin  `xml:"origin"`
	client     *Client
}

// Get returns a specific origin by ID, or an error
func (o Origins) Get(ID int) (*Origin, error) {
	if o.Origins == nil {
		return nil, errors.New("no origins found")
	}
	if len(o.Origins) == 0 {
		return nil, errors.New("no origins found")
	}
	for _, v := range o.Origins {
		if v.Node == ID {
			return &v, nil
		}
	}
	return nil, errors.New("nrigin ID not found")
}

// Origin represents the structure for a single origin returned from 511.org
type Origin struct {
	City      string `xml:"city"`
	MainRoad  string `xml:"mainRoad"`
	CrossRoad string `xml:"crossRoad"`
	Node      int    `xml:"node"`
	client    *Client
}

// Destinations returns the list of valid destinations for this specific origin
func (o Origin) Destinations() (*Destinations, error) {
	return o.client.Destinations(o.Node)
}
