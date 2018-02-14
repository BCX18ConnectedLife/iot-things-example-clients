package rest

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"crypto/tls"
	"github.com/BCX18ConnectedLife/iot-things-example-clients/go/things/client"
)

func Dial(ep, user, pass, token string, cfg *client.Configuration) (RestConnection, error) {
	return &ThingsRestConnection{
		ThingsConnection: client.ThingsConnection{
			Endpoint:  ep + "/api/1",
			Username:  user,
			Password:  pass,
			Token:     token,
			Configuration: cfg,
		},
	}, nil
}

type ThingsRestConnection struct {
	client.ThingsConnection
}

func (c ThingsRestConnection) doRequest(method, url string, body io.Reader) (resp *http.Response, err error) {
	req, err := c.createHttpRequest(method, url, body)
	if err != nil {
		return
	}

	httpClient := c.createHttpClientInstance()
	resp, err = httpClient.Do(req)

	return
}

func (c ThingsRestConnection) createHttpRequest(method, url string, body io.Reader) (req *http.Request, err error) {
	req, err = http.NewRequest(method, url, body)

	if err == nil {
		req.Header.Set("x-cr-api-token", c.Token)
		req.SetBasicAuth(c.Username, c.Password)
		req.Header.Set("Content-Type", "application/json")
	}

	return
}

func (c ThingsRestConnection) createHttpClientInstance() *http.Client {
	cli := &http.Client{

	}
	tr := &http.Transport{}

	cfg := c.Configuration
	if cfg != nil {
		if cfg.Proxy != "" {
			u, _ := url.Parse(cfg.Proxy)
			tr.Proxy = http.ProxyURL(u)
		}

		if (cfg.SkipSslVerify) {
			tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		}
	}

	cli.Transport = tr

	return cli
}

func (c ThingsRestConnection) createUrl(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}
