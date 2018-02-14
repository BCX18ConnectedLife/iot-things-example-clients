package ws

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
	"crypto/tls"
	"log"
	"strings"
	"encoding/json"
	"errors"
	"time"
	"net/url"
	"github.com/BCX18ConnectedLife/iot-things-example-clients/go/things/client"
)

func Dial(ep, user, pass, token string, cfg *client.Configuration) (WebSocketConnection, error) {
	conn := &ThingsWebSocketConnection{
		requests: make(map[string]chan *WSMessage),
		ThingsConnection: client.ThingsConnection{
			Endpoint: ep,
			Username: user,
			Password: pass,
			Token:    token,
			Configuration:      cfg,
		},
	}

	err := conn.dial()

	return conn, err
}

type ThingsWebSocketConnection struct {
	client.ThingsConnection
	wsConn          *websocket.Conn
	requests        map[string]chan *WSMessage
	obsEvents       []chan *WSMessage
}

func (c *ThingsWebSocketConnection) Add(t *Thing) (nt *Thing, err error) {
	tid := ":" + uuid.Must(uuid.NewV4()).String()
	if t.ThingId != "" {
		tid = t.ThingId
	}

	ch, cid, payload, err := c.doAsyncRequest(tid, "create", "/", t)
	defer c.deleteAsyncRequest(cid, ch)
	if err != nil {
		return
	}

	err = c.wsConn.WriteJSON(payload)
	if err != nil {
		return
	}

	msg := <-ch

	if msg.Status != 201 {
		err = errors.New("Expecting 201, returned " + msg.StatusAsString())
		return
	}

	// TODO: Improve this. Currently converting MAP > JSON String > Object
	nt = msg.ValueAsThing()
	return
}

func (c *ThingsWebSocketConnection) Update(t *Thing) (err error) {
	ch, cid, payload, err := c.doAsyncRequest(t.ThingId, "modify", "/", t)
	defer c.deleteAsyncRequest(cid, ch)
	if err != nil {
		return
	}

	err = c.wsConn.WriteJSON(payload)
	if err != nil {
		return
	}

	msg := <-ch

	if msg.Status != 204 {
		err = errors.New("Expecting 204, returned " + msg.StatusAsString())
	}
	return
}

func (c *ThingsWebSocketConnection) Get(s string) (nt *Thing, err error) {
	ch, cid, payload, err := c.doAsyncRequest(s, "retrieve", "/", nil)
	defer c.deleteAsyncRequest(cid, ch)
	err = c.wsConn.WriteJSON(payload)
	if err != nil {
		panic(err.Error())
	}

	msg := <-ch

	if msg.Status == 404 {
		err = errors.New("Thing was not found")
		return
	}

	// TODO: Improve this. Currently converting MAP > JSON String > Object
	pre := msg.Value
	jsonString, err := json.Marshal(pre)
	if err != nil {
		return
	}

	json.Unmarshal(jsonString, &nt)
	return
}

func (c *ThingsWebSocketConnection) Delete(tid string) (err error) {
	ch, cid, payload, err := c.doAsyncRequest(tid, "delete", "/", nil)
	defer c.deleteAsyncRequest(cid, ch)
	err = c.wsConn.WriteJSON(payload)
	if err != nil {
		return
	}

	msg := <-ch

	if msg.Status != 204 {
		err = errors.New("Expecting 204, returned " + msg.StatusAsString())
	}
	return
}

func (c *ThingsWebSocketConnection) Search(ns string, q Query) (sr *SearchResults, err error) {
	ch, cid, payload, err := c.doAsyncRequest(ns + ":_", "search", "/", nil)
	defer c.deleteAsyncRequest(cid, ch)

	qsMap := make(map[string]interface{})
	json.Unmarshal([]byte(q.ToString()), &qsMap)

	if q.ToString() != "?" {
		log.Println("qString", q.ToString())
		log.Println("qsMap", qsMap)

		payload.Value = qsMap
	}

	z, _ := json.Marshal(payload)
	log.Println(string(z))
	err = c.wsConn.WriteJSON(payload)
	if err != nil {
		return
	}

	msg := <-ch

	pre := msg.Value
	jsonString, err := json.Marshal(pre)
	if err != nil {
		return
	}

	err = json.Unmarshal(jsonString, &sr)

	return

	/*
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
	 */
	return nil, ERR_NOT_IMPLEMENTED
}

func (c *ThingsWebSocketConnection) List() ([]*Thing, error) {
	/*
	t, err = c.Get("")

	return
	 */
	return nil, ERR_NOT_IMPLEMENTED
}

func (c *ThingsWebSocketConnection) Claim(id string, payload interface{}, payloadType byte) (interface{}, error) {
	return nil, ERR_NOT_IMPLEMENTED
}

func (c *ThingsWebSocketConnection) SendMessageTo(id, subject string, payload interface{}, payloadType byte) error {
	return ERR_NOT_IMPLEMENTED
}

func (c *ThingsWebSocketConnection) startSendEvents() (err error) {
	err = c.wsConn.WriteMessage(websocket.TextMessage, []byte("START-SEND-EVENTS"))

	return
}

func (c *ThingsWebSocketConnection) ObserveEvents(ch chan *WSMessage) (err error) {
	if err = c.startSendEvents(); err != nil {
		return
	}

	c.obsEvents = append(c.obsEvents, ch)
	return nil
}

func (c *ThingsWebSocketConnection) createTopic(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

func (c *ThingsWebSocketConnection) doAsyncRequest(tid, command, path string, value interface{}) (ch chan *WSMessage, cid string, payload *WSMessage, err error) {
	fmt.Println("doAsyncRequest", tid, command, path, value)
	ch = make(chan *WSMessage)
	cid = uuid.Must(uuid.NewV4()).String()
	c.requests[cid] = ch

	payload = &WSMessage{
		Headers: &WSMessageHeader{},
	}

	parts := strings.Split(tid, ":")
	ns := parts[0]
	id := parts[1]

	if ns == "" || id == "" {
		err = errors.New("Namespace or Thing ID cannot be empty!")
		return
	}

	if command == "search" {
		payload.Topic = c.createTopic("%s/%s/things/twin/search", ns, id)
	} else {
		payload.Topic = c.createTopic("%s/%s/things/twin/commands/%s", ns, id, command)
	}

	payload.Headers.CorrelationId = cid
	payload.Path = path
	payload.Value = value

	log.Println(payload)

	return
}

func (c *ThingsWebSocketConnection) deleteAsyncRequest(id string, ch chan *WSMessage) {
	delete(c.requests, id)
	close(ch)
}

func (c *ThingsWebSocketConnection) startPing() {
	ticker := time.NewTicker(60 * time.Second)
	for {
		select {
		case <-ticker.C:
			if err := c.wsConn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func (c *ThingsWebSocketConnection) start() {
	go c.startPing()

	for {
		messageType, message, err := c.wsConn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		c.handleMessage(messageType, message)
	}
}

func (c *ThingsWebSocketConnection) dispatchEvent(chs []chan *WSMessage, msg *WSMessage) {
	for _, ch := range chs {
		ch <- msg
	}
}

func (c *ThingsWebSocketConnection) messageOfEventType(msg *WSMessage, t int) bool {
	return false
}

func (c *ThingsWebSocketConnection) handleMessage(mt int, b []byte) {
	var msg *WSMessage

	var msg_str = string(b)
	if strings.HasPrefix(msg_str, "{") { // only JSON strings
		json.Unmarshal(b, &msg)

		if msg.Topic != "" { // Incoming Event
			c.dispatchEvent(c.obsEvents, msg)

			switch {
			case c.messageOfEventType(msg, EventTypeCreate):
			case c.messageOfEventType(msg, EventTypeDelete):
			case c.messageOfEventType(msg, EventTypeUpdate):
			}
		} else { // Incoming Request
			cid := msg.Headers.CorrelationId
			if cid != "" {
				ch := c.requests[cid]
				if ch != nil {
					ch <- msg
				} else {
					log.Println("Unknown correlation id encountered")
				}
			}
		}
	} else {
		log.Println("Received: " + msg_str)
	}
}

func (c ThingsWebSocketConnection) createDialer() *websocket.Dialer {
	dialer := &websocket.Dialer{
		TLSClientConfig: &tls.Config{},
	}

	cfg := c.Configuration
	if cfg != nil {
		if cfg.SkipSslVerify {
			dialer.TLSClientConfig.InsecureSkipVerify = cfg.SkipSslVerify
		}

		if cfg.Proxy != "" {
			dialer.Proxy = func(*http.Request) (*url.URL, error) {
				return url.Parse(cfg.Proxy)
			}
		}
	}
	return dialer
}

func (c *ThingsWebSocketConnection) dial() (err error) {
	headers := http.Header{}
	headers["Authorization"] = []string{fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(c.Username+":"+c.Password)))}
	headers["x-cr-api-token"] = []string{c.Token}

	d := c.createDialer()
	c.wsConn, _, err = d.Dial(fmt.Sprintf("%s/ws/1", c.Endpoint), headers)
	if err != nil {
		return
	}

	go c.start()

	return
}
