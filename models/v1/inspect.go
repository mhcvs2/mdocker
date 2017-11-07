package docker

import (
	"mdocker/utils"
	"fmt"
	"strings"
	myLogger "mdocker/logger"
)

type Inspect struct {
	Manage *DockerManage
	cmd string
}

func NewInspect() *Inspect {
	manage := NewDockerManage()
	cmd := manage.getCmd("inspect")
	return &Inspect{manage, cmd}
}

func (inspect *Inspect) getInspectCmd(s string) string {
	return inspect.cmd + " " +s
}

func (inspect *Inspect) getips(ids []string) (ips []string) {
	ips_cmd := inspect.getInspectCmd("--format '{{ .NetworkSettings.IPAddress }}' ")
	for _, name := range ids {
		out, err := utils.Exec_shell(ips_cmd + name)
		if err != nil {
			myLogger.Log.Error("Get ip of " + name + "fail: " + err.Error())
			continue
		}
		ips = append(ips, out)
	}
	return ips
}

func (inspect *Inspect) idsNames() (ids, names []string) {
	ps_cmd := inspect.Manage.getCmd("ps")
	ids_cmd := inspect.Manage.getCmd("ps | awk '{print $1}'")
	names_cmd := inspect.Manage.getCmd("ps | awk '{print $NF}'")
	if res := utils.OutIsNil(ps_cmd); res == false {
		myLogger.Log.Fatal("no alive container")
	}

	out, err := utils.Exec_shell(ids_cmd)
	if err != nil {
		myLogger.Log.Fatal("docker ps err: " + err.Error())
	}
	ids = strings.Split(out, "\n")[1:]

	out, err = utils.Exec_shell(names_cmd)
	if err != nil {
		myLogger.Log.Fatal("docker ps err: " + err.Error())
	}
	names = strings.Split(out, "\n")[1:]
	return ids, names
}

func (inspect *Inspect) ipsOut(ips [] string, namess ...[]string) {
	printMessage := ""
	for i :=0; i < len(ips); i++ {
		printMessage += ips[i]+"  "
		if len(namess) > 0 {
			for j := 0; j < len(namess); j++ {
				printMessage += namess[j][i] +"  "
			}
			printMessage += "\n"
		}
	}
	fmt.Print(printMessage)
}

func (inspect *Inspect) AllInfo(name string) {
	if out, err := utils.Exec_shell(inspect.cmd + " " + name); err != nil {
		myLogger.Log.Fatal(err.Error())
	} else {
		fmt.Println(out)
	}
}

func (inspect *Inspect) AllIps() {
	ids, names := inspect.idsNames()
	ips := inspect.getips(ids)
	inspect.ipsOut(ips, ids, names)
}

func (inspect *Inspect) Ips(names []string) {
	ips := inspect.getips(names)
	inspect.ipsOut(ips, names)
}
