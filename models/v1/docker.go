package docker

import (
	"mdocker/utils"
	myLogger "mdocker/logger"
	"mdocker/pipline"
	"strings"
	"mdocker/config"
)

type DockerManage struct {
	Cmd string
}

func NewDockerManage() *DockerManage {
	_, err := utils.Exec_shell("which docker")
	if err != nil {
		myLogger.Log.Fatal("docker not found in this host")
	}
	return &DockerManage{"docker"}
}

func (docker *DockerManage) getCmd(s string) string {
	return docker.Cmd + " " +s
}

func (docker *DockerManage) baseManage(action, name string) string {

	action_cmd := docker.getCmd(action + " " + name)
	if _, err := utils.Exec_shell(action_cmd); err != nil {
		return err.Error()
	} else {
		return "'"+action_cmd+"' success."
	}
}

func (docker *DockerManage) Stop(name string) string {
	return docker.baseManage("stop", name)
}

func (docker *DockerManage) Start(name string) string {
	return docker.baseManage("start", name)
}

func (docker *DockerManage) Rm(name string) string {
	return docker.baseManage("rm -f", name)
}

func (docker *DockerManage) GetAllId(ignore bool) []string {
	cmd := docker.getCmd("ps -qa")
	out,_ := utils.Exec_shell(cmd)
	if out != "" {
		if ignore {
			return utils.DeleteFrom(strings.Split(out, "\n"), config.Conf.Ignore)
		} else {
			return strings.Split(out, "\n")
		}
	} else {
		return []string{}
	}
}

func (docker *DockerManage) GetAllStartId(ignore bool) []string {
	cmd := docker.getCmd("ps -q")
	out,_ := utils.Exec_shell(cmd)
	if out != "" {
		if ignore {
			return utils.DeleteFrom(strings.Split(out, "\n"), config.Conf.Ignore)
		} else {
			return strings.Split(out, "\n")
		}
	} else {
		return []string{}
	}
}

func (docker *DockerManage) GetAllStopId(ignore bool) (ids []string) {
	all := docker.GetAllId(ignore)
	start := docker.GetAllStartId(ignore)
	return utils.DeleteFrom(all, start)
}

func Run(action string, names []string) []string {
	docker := NewDockerManage()
	if len(names) == 0 {
		names = docker.GetAllId(true)
	}
	done := make(chan struct{})
	defer close(done)
	res := pipline.RunPipline(docker, action, names, done)
	return res
}


