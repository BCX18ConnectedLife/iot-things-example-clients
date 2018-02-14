package rest

import (
	"encoding/json"
	"io/ioutil"
)

func (c ThingsRestConnection) Search(ns string, q Query) (sr *SearchResults, err error) {
	var qURL = "%s/search/things%s"
	qURL = c.createUrl(qURL, c.Endpoint, q.ToString())
	
	resp, err := c.doRequest("GET", qURL, nil)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &sr)
	return
}
