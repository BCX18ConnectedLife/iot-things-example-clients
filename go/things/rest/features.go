package rest

import (
	"encoding/json"
	"io/ioutil"

	"bytes"
)

func (c ThingsRestConnection) UpsertFeature(id, name string, feature *Feature) (err error) {
	b, err := json.Marshal(feature)
	if err != nil {
		return
	}
	_, err = c.doRequest("PUT", c.createUrl("%s/things/%s/features/%s", c.Endpoint, id, name), bytes.NewBuffer(b))
	return
}

func (c ThingsRestConnection) GetFeature(id, feature string) (f map[string]*Feature, err error) {
	resp, err := c.doRequest("GET", c.createUrl("%s/things/%s/features/%s", c.Endpoint, id, feature), nil)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &f)
	return
}

func (c ThingsRestConnection) ListFeatures(id string) ([]*Feature, error) {
	// GET /things/{thingId}/features
	return nil, ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) DeleteFeature(id, featureId string) error {
	// DELETE /things/{thingId}/features/{featureId}
	return ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) DeleteFeatures(id string) error {
	// DELETE /things/{thingId}/features
	return ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) UpsertFeatures(id string, f *Feature) (*Feature, error) {
	// PUT /things/{thingId}/features
	return nil, ERR_NOT_IMPLEMENTED
}
