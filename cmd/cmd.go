package cmd

import (
	"mdocker/cmd/commands"
	"mdocker/utils"
)

import (
	_ "mdocker/cmd/commands/v1/testGroup/testCmd"
	_ "mdocker/cmd/commands/defaultGroup/start"
	_ "mdocker/cmd/commands/defaultGroup/stop"
	_ "mdocker/cmd/commands/defaultGroup/rm"
	_ "mdocker/cmd/commands/defaultGroup/ins"
	_ "mdocker/cmd/commands/defaultGroup/clear"
)

import (
	_ "mdocker/cmd/commands/defaultGroup/version"
	//_ "mdocker/cmd/commands/test"
)

var sampleGroupUsageTemplate = `{{description}}

{{"USAGE" | headline}}
    {{app |bold}} {{"[command] [subcommand] [arguments]" | bold}}

{{$length := len .}}{{if gt $length 1}}{{"AVAILABLE COMMAND GROUPS" | headline}}
{{range $i,$group := .}} {{ if ne $group.Name "default" }} {{ if $group.Show }}
  o {{$group.Name | printf "%-20s" | bold}} {{$group.Short}}
{{end}}{{end}}{{end}}{{end}}
{{"AVAILABLE COMMANDS" | headline}} {{range $i,$group := .}} {{ if eq $group.Name "default" }}
{{range $i,$c := $group.Commands}}{{if $c.Runnable}}
  o {{$c.Name | printf "%-20s" | bold}} {{$c.Short}}{{end}}{{end}}
{{end}}{{end}}
Use {{app |bold}} {{"help" | bold}} for more information about all commands.

Use {{app |bold}} {{"[command] help" | bold}} for more information about a command.

Use {{app |bold}} {{"[command] [subcommand] help" | bold}} for more information about a subcommand.
`

var groupsUsageTemplate = `{{description}}

{{"USAGE" | headline}}
    {{app |bold}} {{"[command] [subcommand] [arguments]" | bold}}

{{$length := len .}}{{if gt $length 2}}{{"AVAILABLE COMMAND GROUPS" | headline}}
{{range $i,$group := .}} {{ if ne $group.Name "default" }} {{ if $group.Show }}
  o {{$group.Name | printf "%-31s" | bold}} {{$group.Short | blue}}
{{range $i,$c := $group.Commands}}{{if $c.Runnable}}
        o {{$c.Name | printf "%-25s" | bold}} {{$c.Short}}{{end}}{{end}}
{{end}}{{end}}{{end}}{{end}}
{{"AVAILABLE COMMANDS" | headline}} {{range $i,$group := .}} {{ if eq $group.Name "default" }}
{{range $i,$c := $group.Commands}}{{if $c.Runnable}}
  o {{$c.Name | printf "%-31s" | bold}} {{$c.Short}}{{end}}{{end}}
{{end}}{{end}}
Use {{app |bold}} {{"[command] help" | bold}} for more information about a command.

Use {{app |bold}} {{"[command] [subcommand] help" | bold}} for more information about a subcommand.
`

var groupUsageTemplate = `{{"DESCRIPTION" | headline}}
  {{tmpltostr .Long . | trim}}

{{"USAGE" | headline}}
    {{app |bold}} {{.Name | bold}} {{"[subcommand] [arguments]" | bold}}

{{"AVAILABLE SUBCOMMANDS: " | headline}}
{{range $i,$c := .Commands}}{{if $c.Runnable}}
     o {{$c.Name | printf "%-25s" | bold}} {{$c.Short}}{{end}}{{end}}

Use {{app |bold}} {{.Name | bold}} {{" [subcommand] help" | bold}} for more information about a subcommand.

`

var helpTemplate = `{{"DESCRIPTION" | headline}}
  {{tmpltostr .Long . | trim}}

{{"USAGE" | headline}}
  {{app | bold}} {{.UsageLine | printf "%s" | bold}}
{{if .Options}}{{endline}}{{"OPTIONS" | headline}}{{range $k,$v := .Options}}
  {{$k | printf "-%s" | bold}}
      {{$v}}
  {{end}}{{end}}
`

var ErrorTemplate = `{{app}}: %s.

Use  {{app |bold}} {{"help" | bold}} for more information about {{app}}.

Use  {{app |bold}} {{"[command] help" | bold}} for more information about a command.

Use  {{app |bold}} {{"[command] [subcommand] help" | bold}} for more information about a subcommand.
`

func Usage() {
	utils.Tmpl(groupsUsageTemplate, commandGroupsFilter())
}

func commandGroupsFilter() []*commands.CommandGroup {
	var tmpCommandGroups = []*commands.CommandGroup{}
	for _, cgs := range commands.AvailableCommandGroups {
		if len(cgs.Commands) != 0 || cgs.Name == "default" {
			tmpCommandGroups = append(tmpCommandGroups, cgs)
		}
	}
	return tmpCommandGroups
}

func SampleUsage() {
	utils.Tmpl(sampleGroupUsageTemplate, commandGroupsFilter())
}

func CommandHelp(group *commands.CommandGroup) {
	utils.Tmpl(groupUsageTemplate, group)
}

func SubCommandHelp(cmd *commands.Command) {
	utils.Tmpl(helpTemplate, cmd)
}