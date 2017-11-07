package testCmd
import (
	"mdocker/cmd/commands"
	"fmt"
)

var CmdRun = &commands.Command{
	UsageLine: "test",
	Short:     "short",
	Long: `
long
`,
	Run:    test,
}

func init() {
	//commands.AvailableCommands = append(commands.AvailableCommands, CmdRun)
	commands.AddGroup("test", CmdRun)
}

func test(cmd *commands.Command, args []string) int {
	fmt.Println("i am test in test group")
	return 0
}