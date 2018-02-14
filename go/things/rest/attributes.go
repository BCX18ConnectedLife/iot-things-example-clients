package rest

import (
	"encoding/json"
	"io/ioutil"
)

func (c ThingsRestConnection) GetAttribute(id, path string) (attr interface{}, err error) {
	resp, err := c.doRequest("GET", c.createUrl("%s/things/%s/attributes/%s", c.Endpoint, id, path), nil)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &attr)
	return
}

func (c ThingsRestConnection) UpsertAttribute(id, path string, attrs interface{}) error {
	// PUT /things/{thingId}/attributes/{attributePath}
	return ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) DeleteAttribute(id, path string) error {
	// DELETE /things/{thingId}/attributes
	return ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) GetAttributes(id string) (attrs map[string]interface{}, err error) {
	resp, err := c.doRequest("GET", c.createUrl("%s/things/%s/attributes", c.Endpoint, id), nil)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &attrs)
	return
}

func (c ThingsRestConnection) DeleteAttributes(id string) error {
	// DELETE /things/{thingId}/attributes/{attributePath}
	return ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) UpsertAttributes(id string, attr map[string]interface{}) error {
	// PUT /things/{thingId}/attributes/{attributePath}
	return ERR_NOT_IMPLEMENTED
}
