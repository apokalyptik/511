// Package drive impliments an APi client for the 511.org Drive Time API (http://511.org/developer-resources_driving-times-api.asp)
package drive

import (
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Client is the base interface for accessing the API
type Client struct {
	key          string
	http         *http.Client
	origins      *Origins
	destinations map[int]*Destinations
}

// Traffic fetches traffic conditions (getpathlist API call) given an Origin ID and a Destination ID
func (c Client) Traffic(originID, destID int) (*Traffic, error) {
	list := new(Traffic)
	resp, err := c.http.Get(fmt.Sprintf("https://services.my511.org/traffic/getpathlist.aspx?token=%s&o=%d&d=%d", c.key, originID, destID))
	if err != nil {
		return list, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return list, err
	}
	err = xml.Unmarshal(data, &list)
	if err != nil {
		return list, err
	}
	return list, err
}

// Origins returns a list of all allowed Origins.  This data is cached in the client for 1 hour
// since tit should not change regularly
func (c Client) Origins() (*Origins, error) {
	if c.origins != nil {
		if time.Now().Sub(c.origins.CacheStamp) < time.Hour {
			return c.origins, nil
		}
	}
	list := new(Origins)
	list.client = &c
	resp, err := c.http.Get(fmt.Sprintf("https://services.my511.org/traffic/getoriginlist.aspx?token=%s", c.key))
	if err != nil {
		return c.origins, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.origins, err
	}
	err = xml.Unmarshal(data, &list)
	if err != nil {
		return c.origins, err
	}
	if len(list.Origins) > 0 {
		for k := range list.Origins {
			list.Origins[k].client = &c
		}
	}
	c.origins = list
	return list, nil
}

// Destinations return a list of all allowed Destinations given an Origin Node.  This data is cached in
// the client for 1 hour. Since the data should not change regularly
func (c Client) Destinations(fromOrigin int) (*Destinations, error) {
	var list *Destinations
	if d, ok := c.destinations[fromOrigin]; ok {
		if time.Now().Sub(d.Timestamp) < time.Hour {
			return d, nil
		}
		list = d
	} else {
		list = new(Destinations)
		list.client = &c
	}
	resp, err := c.http.Get(fmt.Sprintf("http://services.my511.org/traffic/getdestinationlist.aspx?token=%s&o=%d", c.key, fromOrigin))
	if err != nil {
		return list, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return list, err
	}
	err = xml.Unmarshal(data, &list)
	if err != nil {
		return list, err
	}
	if len(list.Destinations) > 0 {
		for k := range list.Destinations {
			list.Destinations[k].client = &c
			list.Destinations[k].origin = fromOrigin
		}
	}
	c.destinations[fromOrigin] = list
	return list, nil
}

// New returns a new API client configured to use the supplied APIKey.  The client contacts
// 511.org over HTTPS, which is "safer" than HTTP, but since 511.org uses a certificate
// which is valid, just not for the hostname we're accessing we have to skip checks
// on the certificate itself.
func New(APIKey string) *Client {
	return &Client{
		key:          APIKey,
		destinations: make(map[int]*Destinations),
		http: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
	}
}
