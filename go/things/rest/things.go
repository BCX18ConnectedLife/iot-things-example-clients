package rest

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"net/http"
	"net/url"

	"strings"
	"log"
)

func (c ThingsRestConnection) Get(ids string) (ts []*Thing, err error) {
	var qUrl string

	multi := true
	idsList := strings.Split(ids, ",")
	if ids == "" {
		qUrl = c.createUrl("%s/things", c.Endpoint)
	} else {
		if len(idsList) > 1 {
			qUrl = c.createUrl("%s/things?ids=%s", c.Endpoint, url.QueryEscape(ids))
		} else {
			multi = false
			qUrl = c.createUrl("%s/things/%s", c.Endpoint, ids)
		}
	}

	resp, err := c.doRequest("GET", qUrl, nil)
	if err != nil {
		return
	}

	if resp.StatusCode == 404 {
		return nil, nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	log.Println("Body", string(body))

	if multi {
		err = json.Unmarshal(body, &ts)
	} else {
		var t *Thing
		err = json.Unmarshal(body, &t)

		ts = []*Thing{t}
	}

	return
}

func (c ThingsRestConnection) List() (t []*Thing, err error) {
	t, err = c.Get("")

	return
}

func (c ThingsRestConnection) Add(t *Thing) (nt *Thing, err error) {
	b, err := json.Marshal(t)
	if err != nil {
		return
	}

	var resp *http.Response
	if t.ThingId != "" {
		resp, err = c.doRequest("PUT", c.createUrl("%s/things/%s", c.Endpoint, t.ThingId), bytes.NewBuffer(b))
	} else {
		resp, err = c.doRequest("POST", c.createUrl("%s/things", c.Endpoint), bytes.NewBuffer(b))
	}
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	json.Unmarshal(body, &nt)
	return
}

func (c ThingsRestConnection) Update(t *Thing) (err error) {
	b, err := json.Marshal(t)
	if err != nil {
		return
	}

	_, err = c.doRequest("PUT", c.createUrl("%s/things/%s", c.Endpoint, t.ThingId), bytes.NewBuffer(b))
	return
}

func (c ThingsRestConnection) Delete(id string) (err error) {
	_, err = c.doRequest("DELETE", c.createUrl("%s/things/%s", c.Endpoint, id), nil)

	return
}
