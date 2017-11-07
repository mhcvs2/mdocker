package start
import (
	"mdocker/cmd/commands"
	"mdocker/models/v1"
)

var CmdRun = &commands.Command{
	UsageLine: "clear",
	Short:     "clear img",
	Long: `
clear img
`,
	Run:    clear,
}

var dins = docker.NewImages()

func init() {
	commands.AddGroup("default", CmdRun)
}

func clear(cmd *commands.Command, args []string) int {
	dins.ClearTemp()
	return 0
}