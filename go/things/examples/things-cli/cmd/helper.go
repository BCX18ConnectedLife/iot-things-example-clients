package cmd

import (
	"os/user"
	"os"
	"errors"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"github.com/BCX18ConnectedLife/iot-things-example-clients/go/things"
	"github.com/BCX18ConnectedLife/iot-things-example-clients/go/things/ws"
	"github.com/BCX18ConnectedLife/iot-things-example-clients/go/things/client"
)

var username, password, token, proxy, endpoint, thingId, file, content string
var searchFilter, searchOpts, searchFields, searchNS string

type Config struct {
	Username string	`json:"username"`
	Password string	`json:"password"`
	ApiToken string	`json:"token"`
	Proxy string	`json:"proxy"`
	Endpoint string	`json:"endpoint"`
}

func readConfigFile() (cfg *Config, err error) {
	if !configFileExists() {
		return nil, errors.New("Configuration File does not exist")
	}

	usr, _ := user.Current()
	dir := usr.HomeDir
	f := dir + "/.things-cli.json"
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &cfg)
	if err != nil {
		return
	}
	return
}

func configFileExists() bool {
	usr, _ := user.Current()
	dir := usr.HomeDir
	f := dir + "/.things-cli.json"
	if _, err := os.Stat(f); os.IsNotExist(err) {
		return false
	}
	return true
}

func createConfigFile(cfg *Config) (err error) {
	b, err := json.MarshalIndent(cfg, "", "\t")
	if err != nil {
		return
	}

	usr, _ := user.Current()
	dir := usr.HomeDir
	f := dir + "/.things-cli.json"
	err = ioutil.WriteFile(f, b, 0644)

	return
}

func createConn() (things.WebSocketConnection, error) {
	// Read from .things-cli file
	config, err := readConfigFile()
	if err != nil {
		os.Exit(0)
	}

	ep := config.Endpoint
	user := config.Username
	pw := config.Password
	tok := config.ApiToken
	proxy := config.Proxy

	cfg := &client.Configuration{
		//SkipSslVerify: true,
	}

	if proxy != "" {
		cfg.Proxy = proxy
	}

	conn, err := ws.Dial(ep, user, pw, tok, cfg)
	if err != nil {
		panic(err.Error())
	}

	return conn, err
}

func loadThingsJsonFile(f string) ([]*things.Thing, error) {
	return nil, nil
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(-1)
}
