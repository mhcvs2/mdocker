package img

import (
	"mdocker/cmd/commands"
	"mdocker/docker"
)

var CmdRun = &commands.Command{
	UsageLine: "stop name1 name2\n",
	Short:     "stop containers",
	Long: `
stop containers
`,
	Run:    Stop,
}

var dm = docker.NewDockerManage()

func init() {
	commands.AvailableCommands = append(commands.AvailableCommands, CmdRun)
}

func Stop(cmd *commands.Command, args []string) int {
	dm.Stop(args, true)
	return 0
}