package rest

func (c ThingsRestConnection) Claim(id string, payload interface{}, payloadType byte) (interface{}, error) {
	// POST /things/{thingId}/inbox/claim
	return nil, ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) SendMessageTo(id, subject string, payload interface{}, payloadType byte) error {
	// POST /things/{thingId}/inbox/messages/{messageSubject}
	return ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) SendsMessageFrom(id, subject string, payload interface{}, payloadType byte) error {
	// POST /things/{thingId}/outbox/messages/{messageSubject}
	return ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) SendMessageToFeature(id, featureId, subject string, payload interface{}, payloadType byte) error {
	// POST /things/{thingId}/features/{featureId}/inbox/messages/{messageSubject}
	return ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) SendMessageFromFeature(id, featureId, subject string, payload interface{}, payloadType byte) error {
	// POST /things/{thingId}/features/{featureId}/outbox/messages/{messageSubject}
	return ERR_NOT_IMPLEMENTED
}
