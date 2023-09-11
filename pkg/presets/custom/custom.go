package custom

import (
	"fmt"
	"os/exec"
)

type Custom struct {
	Command string `yaml:"command"`
}

func (t Custom) Validate() error {
	if t.Command == "" {
		return fmt.Errorf(`"command" value cannot be empty`)
	}

	return nil
}

func (t Custom) Run(file string) error {
	err := exec.Command("sh", "-c", fmt.Sprintf("%v > %v", t.Command, file)).Run()
	if err != nil {
		return err
	}

	return nil
}
