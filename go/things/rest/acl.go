package rest

func (c ThingsRestConnection) ListACL(id string) ([]*ACLEntry, error) {
	// GET /things/{thingId}/acl
	return nil, ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) GetACL(id string, subject string) (*ACLEntry, error) {
	// GET /things/{thingId}/acl/{authorizationSubject}
	return nil, ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) UpsertACL(id string, acl *ACLEntry) error {
	// PUT /things/{thingId}/acl/{authorizationSubject}
	return ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) DeleteACL(id string, acl string) error {
	// DELETE /things/{thingId}/acl/{authorizationSubject}
	return ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) UpdateACL(id string, acl map[string]*ACLEntry) error {
	// PUT /things/{thingId}/acl
	return ERR_NOT_IMPLEMENTED
}
