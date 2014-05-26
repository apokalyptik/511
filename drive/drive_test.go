package drive

import (
	"flag"
	"testing"
)

var testkey string
var origins *Origins
var client *Client
var err error

func init() {
	flag.StringVar(&testkey, "apikey", "", "An API Key to test with")
}

func TestOrigins(t *testing.T) {
	client = New(testkey)
	origins, err = client.Origins()
	if err != nil {
		t.Errorf("Expected no error, got \"%s\"", err.Error())
	}
	if origins == nil {
		t.Errorf("Expected non nil origins, got nil")
	}
	found := false
	for _, v := range origins.Origins {
		if v.Node != 1265 {
			continue
		}
		found = true
		break
	}
	if false == found {
		t.Errorf("Expected to find origin node 1265, did not")
	}
}

func TestDestinations(t *testing.T) {
	dests, err := client.Destinations(1265)
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
	if dests == nil {
		t.Errorf("Expected destinations not to be nil")
	}
	if len(dests.Destinations) < 1 {
		t.Errorf("Expected more than 0 destintions")
	}
}

func TestTraffic(t *testing.T) {
	origins, _ := client.Origins()
	origin, _ := origins.Get(1162)
	dests, _ := origin.Destinations()
	dest, _ := dests.Get(1265)
	traffic, err := dest.Traffic()
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
	if traffic == nil {
		t.Errorf("Expected traffic routes not to be nil")
	}
	if len(traffic.Routes) < 1 {
		t.Errorf("Expected more than zero routes")
	}
	if traffic.Routes[0].TravelTime == 0 {
		t.Errorf("Expected a non zero drive time on first route")
	}
}
