package rest

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (c ThingsRestConnection) ListRelations() (r []*Relation, err error) {
	resp, err := c.doRequest("GET", c.createUrl("%s/relations", c.Endpoint), nil)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &r)

	return
}

func (c ThingsRestConnection) AddRelation(o *Relation) (r *Relation, err error) {
	b, err := json.Marshal(o)
	if err != nil {
		return
	}

	var resp *http.Response
	if o.RelationId != "" {
		resp, err = c.doRequest("PUT", c.createUrl("%s/relations/%s", c.Endpoint, o.RelationId), bytes.NewBuffer(b))
	} else {
		resp, err = c.doRequest("POST", c.createUrl("%s/relations", c.Endpoint), bytes.NewBuffer(b))
	}

	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	json.Unmarshal(body, &r)
	return
}

func (c ThingsRestConnection) DeleteRelation(id string) error {
	// DELETE /relations/{relationId}
	return ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) GetRelation(id string) (*Relation, error) {
	// GET /relations/{relationId}
	return nil, ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) UpsertRelation(id string) (*Relation, error) {
	// PUT /relations/{relationId}
	return nil, ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) GetRelationACL(id string) (map[string]*ACLEntry, error) {
	// GET /relations/{relationId}/acl
	return nil, ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) UpsertRelationACL(id string) error {
	// PUT /relations/{relationId}/acl
	return ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) DeleteRelationACLForSubject(id, subject string) error {
	// DELETE /relations/{relationId}/acl/{authorizationSubject}
	return ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) GetRelationACLForSubject(id, subject string) (*ACLEntry, error) {
	// GET /relations/{relationId}/acl/{authorizationSubject}
	return nil, ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) UpsertRelationACLForSubject(id, subject string) (*ACLEntry, error) {
	// PUT /relations/{relationId}/acl/{authorizationSubject}
	return nil, ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) GetRelationSourceId(id string) (string, error) {
	// GET /relations/{relationId}/source
	return "", ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) UpdateRelationSource(id string) error {
	// PUT /relations/{relationId}/source
	return ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) GetRelationTargetId(id string) (string, error) {
	// GET /relations/{relationId}/target
	return "", ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) UpdateRelationTarget(id string) error {
	// PUT /relations/{relationId}/target
	return ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) DeleteRelationAttributes(id string) error {
	// DELETE /relations/{relationId}/attributes
	return ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) ListRelationAttributes(id string) (map[string]interface{}, error) {
	// GET /relations/{relationId}/attributes
	return nil, ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) UpsertRelationAttributes(id string) error {
	// PUT /relations/{relationId}/attributes
	return ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) DeleteRelationAttribute(id, path string) error {
	// DELETE /relations/{relationId}/attributes/{attributePath}
	return ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) GetRelationAttribute(id, path string) (map[string]interface{}, error) {
	// GET /relations/{relationId}/attributes/{attributePath}
	return nil, ERR_NOT_IMPLEMENTED
}

func (c ThingsRestConnection) UpsertRelationAttribute(id, path string) error {
	// PUT /relations/{relationId}/attributes/{attributePath}
	return ERR_NOT_IMPLEMENTED
}
