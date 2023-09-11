package mongo

import (
	"fmt"
	"os/exec"
)

type Mongo struct {
	Container     string `yaml:"container"`
	MongoUser     string `yaml:"mongouser"`
	MongoPassword string `yaml:"mongopassword"`
}

func (t Mongo) Validate() error {
	if t.Container == "" {
		return fmt.Errorf(`"container" value cannot be empty`)
	}
	if t.MongoUser == "" {
		return fmt.Errorf(`"mongouser" value cannot be empty`)
	}
	if t.MongoPassword == "" {
		return fmt.Errorf(`"mongopassword" value cannot be empty`)
	}

	return nil
}

func (t Mongo) Run(file string) error {

	err := exec.Command("sh", "-c", fmt.Sprintf(`docker exec -i %v /usr/bin/mongodump --username "%v" --password "%v" --authenticationDatabase admin --gzip --archive > %v`,
		t.Container, t.MongoUser, t.MongoPassword, file)).Run()
	if err != nil {
		return err
	}

	return nil
}
