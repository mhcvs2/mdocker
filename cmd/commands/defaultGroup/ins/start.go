package start
import (
	"mdocker/cmd/commands"
	"mdocker/models/v1"
)

var CmdRun = &commands.Command{
	UsageLine: "ins",
	Short:     "get ips",
	Long: `
get ips
`,
	Run:    ins,
}

var dins = docker.NewInspect()

func init() {
	commands.AddGroup("default", CmdRun)
}

func ins(cmd *commands.Command, args []string) int {
	dins.AllIps()
	return 0
}