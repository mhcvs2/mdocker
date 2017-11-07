package docker

import (
	"mdocker/utils"
	myLogger "mdocker/logger"
	"fmt"
)

type DockerManage struct {
	Cmd string
}

func NewDockerManage() *DockerManage {
	out, err := utils.Exec_shell("which docker")
	if err != nil {
		myLogger.Log.Fatal("docker not found in this host")
	}
	return &DockerManage{out}
}

func (docker *DockerManage) getCmd(s string) string {
	return docker.Cmd + " " +s
}

func (docker *DockerManage) baseManage(action, sub_cmd string, names []string, check bool) {
	res := utils.OutIsNil(sub_cmd)
	if check {
		if res == false {
			myLogger.Log.Fatal("no container to act.")
		}
	} else if res == false {
		return
	}
	action_cmd := docker.getCmd(action)
	if len(names) == 0 {
		action_cmd = fmt.Sprintf(action_cmd + " $(%s)", sub_cmd)
	} else {
		for _, name := range names {
			action_cmd += " " + name
		}
	}
	if _, err := utils.Exec_shell(action_cmd); err != nil {
		myLogger.Log.Fatal(err.Error())
	} else {
		myLogger.Log.Info("'"+action_cmd+"' success.")
	}
}

func (docker *DockerManage) Stop(names []string, check bool) {
	sub_cmd := docker.getCmd("ps -q")
	docker.baseManage("stop", sub_cmd, names, check)
}

func (docker *DockerManage) Start(names []string, check bool) {
	sub_cmd := docker.getCmd("ps -qa")
	docker.baseManage("start", sub_cmd, names, check)
}

func (docker *DockerManage) Remove(names []string, check bool) {
	sub_cmd := docker.getCmd("ps -qa")
	docker.baseManage("rm -f", sub_cmd, names, check)
}




