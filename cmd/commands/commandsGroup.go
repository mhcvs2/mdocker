package commands

import "strings"

type CommandGroup struct {

	//Group name
	Name string

	// PreRun performs an operation before running the command group
	PreRun func(cmd *Command, args []string)

	// Short is the short description shown in the 'go help' output.
	Short string

	//Long is the long message shown in the 'go help <this-command>' output.
	Long string

	//if true, command group is not showed in help and can't be use.
	Disabled bool

	Commands []*Command

}

func (cgp *CommandGroup) AddCommand(command *Command) {
	cgp.Commands = append(cgp.Commands, command)
}

func (cgp *CommandGroup) Show() bool {
	return !cgp.Disabled
}


var AvailableCommandGroups = []*CommandGroup{}

func GetCommandGroup(GroupName string) (*CommandGroup, bool) {
	for _, cgp := range AvailableCommandGroups {
		if strings.Contains(cgp.Name, GroupName) {
			return cgp, true
		}
	}
	return nil, false
}

func GetDefaultCommandGroup() (*CommandGroup, bool) {
	defaultGroup, ok := GetCommandGroup("default")
	return defaultGroup, ok
}

func AddGroup(groupName string, command *Command) {
	cgp,ok := GetCommandGroup(groupName)
	if !ok {
		cgp = &CommandGroup{
			Name: groupName,
			Short: "no short description",
			Long:
			`no long description
			`,
			Commands: []*Command{},
		}
		AvailableCommandGroups = append(AvailableCommandGroups, cgp)
	}
	cgp.AddCommand(command)
}
