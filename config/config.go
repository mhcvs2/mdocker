package config

import (
	"encoding/json"
	"io/ioutil"
	beeLogger "mdocker/logger"
)

var Conf = struct {
	Version            int
	GoInstall          bool      `json:"go_install" yaml:"go_install"` // Indicates whether execute "go install" before "go build".
	DirStruct          dirStruct `json:"dir_structure" yaml:"dir_structure"`
	CmdArgs            []string  `json:"cmd_args" yaml:"cmd_args"`
	Envs               []string
	Bale               bale
	Database           database
	EnableReload       bool              `json:"enable_reload" yaml:"enable_reload"`
	EnableNotification bool              `json:"enable_notification" yaml:"enable_notification"`
	Scripts            map[string]string `json:"scripts" yaml:"scripts"`
}{
	GoInstall: true,
	DirStruct: dirStruct{
		Others: []string{},
	},
	CmdArgs: []string{},
	Envs:    []string{},
	Bale: bale{
		Dirs:   []string{},
		IngExt: []string{},
	},
	Database: database{
		Driver: "mysql",
	},
	EnableNotification: true,
	Scripts:            map[string]string{},
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

	// Set variables
	if len(Conf.DirStruct.Controllers) == 0 {
		Conf.DirStruct.Controllers = "controllers"
	}

	if len(Conf.DirStruct.Models) == 0 {
		Conf.DirStruct.Models = "models"
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
