package things

type Client interface {
	Dial(ep, user, pass, token string) (Connection, error)
}

type Connection interface {
	Add(*Thing) (*Thing, error)
	List() ([]*Thing, error)
	Update(*Thing) error
	Delete(string) error

	// Search
	Search(ns string, q Query) (*SearchResults, error)

	// Messaging
	Claim(id string, payload interface{}, payloadType byte) (interface{}, error)
	SendMessageTo(id, subject string, payload interface{}, payloadType byte) error
}

type RestConnection interface {
	Connection

	Get(string) ([]*Thing, error)

	// ACL
	ListACL(id string) ([]*ACLEntry, error)
	GetACL(id string, subject string) (*ACLEntry, error)
	UpsertACL(id string, acl *ACLEntry) error
	DeleteACL(id string, acl string) error
	UpdateACL(id string, acl map[string]*ACLEntry) error

	// Attributes
	GetAttributes(id string) (map[string]interface{}, error)
	UpsertAttributes(id string, attr map[string]interface{}) error
	DeleteAttributes(id string) error
	GetAttribute(id, path string) (interface{}, error)
	UpsertAttribute(id, path string, attrs interface{}) error
	DeleteAttribute(id, path string) error

	// Features
	GetFeature(id, feature string) (map[string]*Feature, error)
	UpsertFeature(id, name string, feature *Feature) error
	DeleteFeature(id, featureId string) error
	ListFeatures(id string) ([]*Feature, error)
	DeleteFeatures(id string) error
	UpsertFeatures(id string, f *Feature) (*Feature, error)

	// Feature Properties
	ListFeatureProperties(id, featureId string) (map[string]interface{}, error)
	DeleteFeatureProperties(id, featureId string) error
	UpsertFeatureProperties(id, featureId string) error
	DeleteFeatureProperty(id, featureId, path string) error
	GetFeatureProperty(id, featureId, path string) (interface{}, error)
	UpsertFeatureProperty(id, featureId, path string) error

	// Relations
	ListRelations() (r []*Relation, err error)
	AddRelation(o *Relation) (r *Relation, err error)
	DeleteRelation(id string) error
	GetRelation(id string) (*Relation, error)
	UpsertRelation(id string) (*Relation, error)
	GetRelationACL(id string) (map[string]*ACLEntry, error)
	UpsertRelationACL(id string) error
	DeleteRelationACLForSubject(id, subject string) error
	GetRelationACLForSubject(id, subject string) (*ACLEntry, error)
	UpsertRelationACLForSubject(id, subject string) (*ACLEntry, error)
	GetRelationSourceId(id string) (string, error)
	UpdateRelationSource(id string) error
	GetRelationTargetId(id string) (string, error)
	UpdateRelationTarget(id string) error
	DeleteRelationAttributes(id string) error
	ListRelationAttributes(id string) (map[string]interface{}, error)
	UpsertRelationAttributes(id string) error
	DeleteRelationAttribute(id, path string) error
	GetRelationAttribute(id, path string) (map[string]interface{}, error)
	UpsertRelationAttribute(id, path string) error

	SendsMessageFrom(id, subject string, payload interface{}, payloadType byte) error
	SendMessageToFeature(id, featureId, subject string, payload interface{}, payloadType byte) error
	SendMessageFromFeature(id, featureId, subject string, payload interface{}, payloadType byte) error
}

type WebSocketConnection interface {
	Connection

	Get(string) (*Thing, error)

	ObserveEvents(chan *WSMessage) error
}
