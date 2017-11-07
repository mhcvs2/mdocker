package img

import (
	"mdocker/cmd/commands"
	"mdocker/docker"
)

var CmdRun = &commands.Command{
	UsageLine: "start name1 name2\n",
	Short:     "start containers",
	Long: `
start containers
`,
	Run:    Stop,
}

var dm = docker.NewDockerManage()

func init() {
	commands.AvailableCommands = append(commands.AvailableCommands, CmdRun)
}

func Stop(cmd *commands.Command, args []string) int {
	dm.Start(args, true)
	return 0
}