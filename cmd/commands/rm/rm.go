package img

import (
	"mdocker/cmd/commands"
	"mdocker/docker"
)

var CmdRun = &commands.Command{
	UsageLine: "rm name1 name2\n",
	Short:     "remove containers",
	Long: `
remove containers
`,
	Run:    Stop,
}

var dm = docker.NewDockerManage()

func init() {
	commands.AvailableCommands = append(commands.AvailableCommands, CmdRun)
}

func Stop(cmd *commands.Command, args []string) int {
	dm.Stop(args, false)
	dm.Remove(args, true)
	return 0
}