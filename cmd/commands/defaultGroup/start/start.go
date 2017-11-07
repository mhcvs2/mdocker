package start
import (
	"mdocker/cmd/commands"
	"mdocker/models/v1"
)

var CmdRun = &commands.Command{
	UsageLine: "start",
	Short:     "start container",
	Long: `
start container
`,
	Run:    start,
}

func init() {
	commands.AddGroup("default", CmdRun)
}

func start(cmd *commands.Command, args []string) int {
	docker.Run("Start", args)
	return 0
}