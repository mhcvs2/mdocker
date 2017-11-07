package main

import (
	"flag"
	//flag "github.com/spf13/pflag"
	"log"
	"os"
	"mdocker/cmd"
	"mdocker/cmd/commands"
	"mdocker/utils"
	MyLogger "mdocker/logger"
	"fmt"
	"mdocker/config"
)

func main() {

	MyLogger.LoadLogOption()
	flag.Usage = cmd.Usage
	flag.Parse()
	log.SetFlags(0)
	if ok, _ := utils.PathExists("/etc/mdocker.conf"); ok{
		config.LoadConfig("/etc/mdocker.conf")
	}

	args := flag.Args()

	if len(args) < 1 {
		cmd.SampleUsage()
		os.Exit(2)
		return
	}

	if args[0] == "help" {
		cmd.Usage()
		os.Exit(2)
		return
	}

	defaultGroup, _ := commands.GetDefaultCommandGroup()
	for _, c := range defaultGroup.Commands {
		if utils.InSlice(args[0], c.SliceName()) && c.Run != nil {
			c.UsageLine = c.UsageLine
			if len(args) > 1  && args[1] == "help"{
				cmd.SubCommandHelp(c)
				os.Exit(2)
				return
			}
			utils.PreRun()
			c.Flag.Usage = func() { c.Usage() }
			if c.CustomFlags {
				args = args[1:]
			} else {
				c.Flag.Parse(args[1:])
				args = c.Flag.Args()
			}

			if c.PreRun != nil {
				c.PreRun(c, args)
			}

			os.Exit(c.Run(c, args))
			return
		}
	}

	commandGroup, ok := commands.GetCommandGroup(args[0])
	if ok && commandGroup.Disabled { ok = false }
	if ok {
		if len(args)== 1 || args[1] == "help"{
			cmd.CommandHelp(commandGroup)
			os.Exit(2)
			return
		}
		for _, c := range commandGroup.Commands {
			if utils.InSlice(args[1], c.SliceName()) && c.Run != nil {
				c.UsageLine = commandGroup.Name + " " + c.UsageLine
				if len(args) > 2  && args[2] == "help"{
					cmd.SubCommandHelp(c)
					os.Exit(2)
					return
				}
				utils.PreRun()
				c.Flag.Usage = func() { c.Usage() }
				if c.CustomFlags {
					args = args[2:]
				} else {
					c.Flag.Parse(args[2:])
					args = c.Flag.Args()
				}
				if commandGroup.PreRun != nil {
					commandGroup.PreRun(c, args)
				}
				if c.PreRun != nil {
					c.PreRun(c, args)
				}

				os.Exit(c.Run(c, args))
				return
			}
		}
		utils.Tmpl(fmt.Sprintf(cmd.ErrorTemplate, "Unknown subcommand"), nil)
		os.Exit(2)
	} else {
		utils.Tmpl(fmt.Sprintf(cmd.ErrorTemplate, "Unknown command"), nil)
		os.Exit(2)
	}
}