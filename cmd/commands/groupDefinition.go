package commands

var defaultGroup = &CommandGroup{
	Name: "default",
	Short: "run subcommand without command",
	Long:
	`run subcommand without command
	`,
	Commands: []*Command{},
}

var test = &CommandGroup{
	Name: "test",
	PreRun: PreRun,
	Disabled:true,
	Short: "test group",
	Long:
	`test group
	`,
	Commands: []*Command{},
}

func init() {
	AvailableCommandGroups = append(AvailableCommandGroups, defaultGroup)
	AvailableCommandGroups = append(AvailableCommandGroups, test)
}

func PreRun(cmd *Command, args []string) {
}