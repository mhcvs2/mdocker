package docker

import (
	"mdocker/utils"
	"fmt"
	"errors"
)

type Images struct {
	Manage *DockerManage
}

func NewImages() *Images {
	manage := NewDockerManage()
	return &Images{manage}
}

func (images *Images) ClearTemp() error {

	get_tmp_image_cmd := images.Manage.getCmd("images -q -f dangling=true")
	rm_cmd := images.Manage.getCmd(fmt.Sprintf("rmi $(%s)", get_tmp_image_cmd))

	if result := utils.OutIsNil(get_tmp_image_cmd); result == false {
		return errors.New("no images to clear")
	}
	_, err := utils.Exec_shell(rm_cmd)
	if err != nil {
		return err
	}
	return nil
}

func (Images *Images) Show() error {
	show_cmd := Images.Manage.getCmd("images")
	if res := utils.OutIsNil(show_cmd); res == false {
		return errors.New("no images")
	}
	out, err := utils.Exec_shell(show_cmd)
	if err != nil {
		return err
	}
	fmt.Println(out)
	return nil
}