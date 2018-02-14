package things

import (
	"bytes"
	"errors"
	"net/url"
	"strings"
	"encoding/json"
)

var ERR_NOT_IMPLEMENTED = errors.New("This method has not been implemented yet.")

const (
	EventTypeCreate = 0
	EventTypeUpdate = 1
	EventTypeDelete = 2
)

func NewThing() *Thing {
	return &Thing{
		Acl:        make(map[string]*ACLEntry),
		Attributes: make(map[string]interface{}),
		Features:   make(map[string]*Feature),
	}
}

func NewThingFromContent(b[] byte) (t *Thing) {
	json.Unmarshal(b, &t)

	return t
}

type WSMessage struct {
	Topic   string           `json:"topic"`
	Headers *WSMessageHeader `json:"headers"`
	Path    string           `json:"path"`
	Value   interface{}      `json:"value,omitempty"`
	Status  int             `json:"status,omitempty"`
}

func (m *WSMessage) StatusAsString() string {
	return string(m.Status)
}

func (m *WSMessage) ValueAsThing() (t *Thing) {
	pre := m.Value
	jsonString, err := json.Marshal(pre)
	if err != nil {
		return
	}

	json.Unmarshal(jsonString, &t)
	return
}

type WSMessageHeader struct {
	CorrelationId string `json:"correlation-id"`
}

type Thing struct {
	ThingId    string                 `json:"thingId,omitempty"`
	Acl        map[string]*ACLEntry   `json:"acl,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Features   map[string]*Feature    `json:"features,omitempty"`
}

func (t *Thing) Bytes() (b []byte, err error) {
	b, err = json.Marshal(t)

	return
}

func (t *Thing) String() (c string, err error) {
	var b []byte
	if b, err = t.Bytes(); err != nil {
		return
	}

	c = string(b)
	return
}

func (t *Thing) GetRelation() *RelationFeature {
	rel := t.Features["relation"]
	if rel != nil {
		if rel.Properties["source"] != nil && rel.Properties["target"] != nil {
			return &RelationFeature{
				Source: rel.Properties["source"].(string),
				Target: rel.Properties["target"].(string),
			}
		}
	}
	return nil
}

type ACLEntry struct {
	Read         bool `json:"READ"`
	Write        bool `json:"WRITE"`
	Administrate bool `json:"ADMINISTRATE"`
}

func NewFeature() *Feature {
	return &Feature{
		Properties: make(map[string]interface{}),
	}
}

type Feature struct {
	Properties map[string]interface{} `json:"properties"`
}

type ErrorResponse struct {
	Status      int    `json:"status"`
	Error       string `json:"error"`
	Message     string `json:"message"`
	Description string `json:"-"`
	Href        string `json:"-"`
}

type Query interface {
	ToString() string
}

type DefaultQuery struct {
	Fields  []string
	Filters []*Filter
}

func NewStringQuery(filter, opts, fields string) Query {
	return &StringQuery{
		Filter:  filter,
		Options: opts,
		Fields:  fields,
	}
}

type StringQuery struct {
	Filter  string
	Options string
	Fields  string
}

func (s *StringQuery) ToString() string {
	var buffer bytes.Buffer

	buffer.WriteString("?")

	if s.Filter != "" {
		c := url.QueryEscape(s.Filter)
		c = strings.Replace(c, "%28", "(", -1)
		c = strings.Replace(c, "%29", ")", -1)

		buffer.WriteString("filter=")
		buffer.WriteString(c)
		buffer.WriteString("&")
	}

	if s.Options != "" {
		c := url.QueryEscape(s.Options)
		c = strings.Replace(c, "%28", "(", -1)
		c = strings.Replace(c, "%29", ")", -1)

		buffer.WriteString("option=")
		buffer.WriteString(c)
		buffer.WriteString("&")
	}

	if s.Fields != "" {
		c := url.QueryEscape(s.Fields)
		c = strings.Replace(c, "%28", "(", -1)
		c = strings.Replace(c, "%29", ")", -1)

		buffer.WriteString("fields=")
		buffer.WriteString(c)
	}

	return buffer.String()
}

type Filter struct {
}

type Option struct {
}

type SearchResults struct {
	Items          []*Thing `json:"items"`
	NextPageOffset int      `json:"nextPageOffset"`
}

func NewRelation() *Relation {
	return &Relation{
		Acl:        make(map[string]*ACLEntry),
		Attributes: make(map[string]interface{}),
	}
}

type RelationFeature struct {
	Source string
	Target string
}

type Relation struct {
	RelationId string                 `json:"relationId,omitempty"`
	Acl        map[string]*ACLEntry   `json:"acl"`
	Attributes map[string]interface{} `json:"attributes"`
	Source     string                 `json:"source"`
	Target     string                 `json:"target"`
}
