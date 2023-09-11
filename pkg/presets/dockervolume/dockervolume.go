package dockervolume

import (
	"fmt"
	"os/exec"
)

type DockerVolume struct {
	Volume string `yaml:"volume"`
}

func (t DockerVolume) Validate() error {
	if t.Volume == "" {
		return fmt.Errorf(`"volume" value cannot be empty`)
	}

	return nil
}

func (t DockerVolume) Run(file string) error {
  err := exec.Command("sh", "-c", fmt.Sprintf("docker run --rm -v %v:/volume busybox sh -c 'tar -cOzf - /volume' > %v", t.Volume, file)).Run()
	if err != nil {
		return err
	}

	return nil
}
