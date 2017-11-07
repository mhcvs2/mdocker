package test
import (
	"mdocker/cmd/commands"
	"fmt"
)

var CmdRun = &commands.Command{
	UsageLine: "test [-test-option=***]",
	Short:     "test short",
	Long: `
test long
`,
	Run:    test,
	PreRun: preTest,
}
var (
	testOption string
)

func init() {
	CmdRun.Flag.StringVar(&testOption, "test-option", "", "test option description")
	commands.AddGroup("default", CmdRun)
}

func preTest(cmd *commands.Command, args []string) {
	fmt.Println("test----------------------------------pre run")
}

func test(cmd *commands.Command, args []string) int {
	fmt.Println("test----------------------------------run")
	return 0
}