package stop
import (
	"mdocker/cmd/commands"
	"mdocker/models/v1"
	"fmt"
)

var CmdRun = &commands.Command{
	UsageLine: "stop",
	Short:     "stop container",
	Long: `
stop container
`,
	Run:    stop,
}

func init() {
	commands.AddGroup("default", CmdRun)
}

func stop(cmd *commands.Command, args []string) int {
	res := docker.Run("Stop", args)
	for _, r := range res {
		fmt.Println(r)
	}
	return 0
}