package rest

func (c ThingsRestConnection) ListFeatureProperties(id, featureId string) (map[string]interface{}, error) {
	// GET /things/{thingId}/features/{featureId}/properties
	return nil, ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) DeleteFeatureProperties(id, featureId string) error {
	// DELETE /things/{thingId}/features/{featureId}/properties
	return ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) UpsertFeatureProperties(id, featureId string) error {
	// PUT /things/{thingId}/features/{featureId}/properties
	return ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) DeleteFeatureProperty(id, featureId, path string) error {
	// DELETE /things/{thingId}/features/{featureId}/properties/{propertyPath}
	return ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) GetFeatureProperty(id, featureId, path string) (interface{}, error) {
	// GET /things/{thingId}/features/{featureId}/properties/{propertyPath}
	return nil, ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) UpsertFeatureProperty(id, featureId, path string) error {
	// PUT /things/{thingId}/features/{featureId}/properties/{propertyPath}
	return ERR_NOT_IMPLEMENTED
}
