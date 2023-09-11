package executer

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/mducoli/cronupper/pkg/types"
)

func bashEval(s string) (string, error) {

	s = strings.ReplaceAll(s, `"`, `\"`)

	out, err := exec.Command("sh", "-c", fmt.Sprintf(`echo "%v"`, s)).Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}

func Execute(job *types.Job) error {

	dir, err := os.MkdirTemp("", "cronupper")
	if err != nil {
		return err
	}

	filename, err := bashEval(job.Upload.Filename)
	if err != nil {
		return err
	}
	filepath := dir + "/" + filename

	log.Printf(`%v: Saving to file: "%v"`, job.Id, filename)

	err = job.Preset.Run(filepath)
	if err != nil {
		return err
	}

	log.Printf(`%v: Saving done`, job.Id)

	err = job.Upload.To.Upload(filepath, filename, job.Upload.Config)
	if err != nil {
		return err
	}

	log.Printf(`%v: Uploading done`, job.Id)

	err = os.RemoveAll(dir)
	if err != nil {
		return err
	}

	return nil
}
