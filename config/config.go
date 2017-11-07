package config

import (
	"encoding/json"
	"io/ioutil"
	beeLogger "mdocker/logger"
)

var Conf = struct {
	AppName            string
	Version            string
	Description        string
	Ignore             []string
	Wn                 int
}{
	AppName: "mdocker",
	Version: "0.1.0",
	Description: "mdocker description",
	Ignore:[]string{},
	Wn:4,
}

// dirStruct describes the application's directory structure
type dirStruct struct {
	WatchAll    bool `json:"watch_all" yaml:"watch_all"`
	Controllers string
	Models      string
	Others      []string // Other directories
}

// bale
type bale struct {
	Import string
	Dirs   []string
	IngExt []string `json:"ignore_ext" yaml:"ignore_ext"`
}

// database holds the database connection information
type database struct {
	Driver string
	Conn   string
}

// LoadConfig loads the bee tool configuration.
// It looks for Beefile or bee.json in the current path,
// and falls back to default configuration in case not found.
func LoadConfig(file string) {

	err := parseJSON(file, &Conf)
	if err != nil {
		beeLogger.Log.Errorf("Failed to parse JSON file: %s", err)
	}
}

func parseJSON(path string, v interface{}) error {
	var (
		data []byte
		err  error
	)
	data, err = ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, v)
	return err
}