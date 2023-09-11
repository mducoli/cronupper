package postgres

import (
	"fmt"
	"os/exec"
)

type Postgres struct {
	Container    string `yaml:"container"`
	PostgresUser string `yaml:"postgresuser"`
}

func (t Postgres) Validate() error {
	if t.Container == "" {
		return fmt.Errorf(`"container" value cannot be empty`)
	}
	if t.PostgresUser == "" {
		return fmt.Errorf(`"postgresuser" value cannot be empty`)
	}

	return nil
}

func (t Postgres) Run(file string) error {

	err := exec.Command("sh", "-c", fmt.Sprintf("docker exec -i %v /usr/bin/pg_dumpall -U %v > %v", t.Container, t.PostgresUser, file)).Run()
  if err != nil {
    return err
  }

	return nil
}
