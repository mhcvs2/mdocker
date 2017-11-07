package start
import (
	"mdocker/cmd/commands"
	"mdocker/models/v1"
)

var CmdRun = &commands.Command{
	UsageLine: "rm",
	Short:     "rm container",
	Long: `
rm container
`,
	Run:    rm,
}

func init() {
	commands.AddGroup("default", CmdRun)
}

func rm(cmd *commands.Command, args []string) int {
	docker.Run("Rm", args)
	return 0
}