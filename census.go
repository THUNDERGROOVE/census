// The census package is used to query data from the census API.
//
// It's centered more so around data from Planetside 2
package census

var BaseURL = "http://census.daybreakgames.com/"
var BaseURLOld = "http://census.soe.com/"

func init() {
	BaseURL = BaseURLOld
}

type CensusData struct {
	Error string `json:"error"`
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
