package ins

import (
	"mdocker/cmd/commands"
	"mdocker/docker"
	"strings"
)

var CmdRun = &commands.Command{
	UsageLine: "ins name \n" +
		"          ins ip name1 name2...",
	Short:     "get ip by container id",
	Long: `
get ip by container id
`,
	Run:    getIp,
}

var inspect = docker.NewInspect()

func init() {
	commands.AvailableCommands = append(commands.AvailableCommands, CmdRun)
}

func getIp(cmd *commands.Command, args []string) int {
	//started <- true

	if len(args) == 0 {
		inspect.AllIps()
	} else {
		switch strings.ToLower(args[0]) {
		case "ip":
			if len(args) == 1 {
				inspect.AllIps()
			} else {
				inspect.Ips(args[1:])
			}
		default:
			inspect.AllInfo(args[0])
		}
	}
	return 0

}
