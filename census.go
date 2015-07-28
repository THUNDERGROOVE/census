// The census package is used to query data from the census API.
//
// It's centered more so around data from Planetside 2
package census

import (
	"strings"
)

var BaseURL = "http://census.daybreakgames.com/"
var BaseURLOld = "http://census.soe.com/"

func init() {
	//BaseURL = BaseURLOld
}

// CensusData is a struct that contains various metadata that a Census request can have.
type CensusData struct {
	error string `json:"error"`
}

func (c *CensusData) Error() string {
	return c.error
}

// NewCensus returns a new census object given your service ID
func NewCensus(ServiceID string, Namespace string) *Census {
	c := new(Census)
	c.serviceID = ServiceID
	c.namespace = Namespace
	return c
}

// Census is the main object you use to query data
type Census struct {
	serviceID string
	namespace string
}

func (c *Census) CleanNamespace() string {
	if strings.Contains(c.namespace, ":") {
		return strings.Split(c.namespace, ":")[0]
	}
	return c.namespace
}

func (c *Census) IsEU() bool {
	if strings.Contains(c.namespace, "eu") {
		return true
	}
	return false
}
