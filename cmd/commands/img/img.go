package img

import (
	"mdocker/cmd/commands"
	myLogger "mdocker/logger"
	"mdocker/docker"
	"strings"
)

var CmdRun = &commands.Command{
	UsageLine: "img show\n" +
		"          img clear",
	Short:     "clear temporary images",
	Long: `
clear temporary images
`,
	Run:    Img,
}

var image = docker.NewImages()

func init() {
	commands.AvailableCommands = append(commands.AvailableCommands, CmdRun)
}

func Img(cmd *commands.Command, args []string) int {
	if len(args) == 0 {
		return showImg()
	}
	switch strings.ToLower(args[0]) {
	case "show":
		return showImg()
	case "clear":
		return clearImg()
	}
	return 0
}

func showImg() int {
	if err := image.Show(); err != nil {
		myLogger.Log.Fatal(err.Error())
	}
	return 0
}

func clearImg() int {
	if err := image.ClearTemp(); err != nil {
		myLogger.Log.Fatal(err.Error())
	}
	return 0
}