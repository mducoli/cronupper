package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/mducoli/cronupper/pkg/presets"
	"github.com/mducoli/cronupper/pkg/types"
	"github.com/mducoli/cronupper/pkg/uploaders"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
)

type configParse struct {
	Jobs map[string]struct {
		Preset string `yaml:"preset"`
		Cron   string `yaml:"cron"`
		Config map[string]interface{}

		Upload struct {
			To       string `yaml:"to"`
			Filename string `yaml:"filename"`
			Config   map[string]interface{}
		} `yaml:"upload"`
	} `yaml:"jobs"`

	Uploaders map[string]struct {
		Provider string `yaml:"provider"`
		Config   map[string]interface{}
	} `yaml:"uploaders"`
}

func Parse(configLocation string) (*types.Config, error) {

	// get file
	file, err := os.ReadFile(configLocation)
	if err != nil {
		return nil, err
	}

	// parse config file
	var raw configParse

	err = yaml.Unmarshal(file, &raw)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("raw: %v\n", raw)

	// create final config

	var config = types.Config{
		Jobs:      map[string]types.Job{},
		Uploaders: map[string]types.Uploader{},
	}

	for k, v := range raw.Uploaders {

		uploader, has := uploaders.Uploaders[v.Provider]
		if !has {
			return nil, fmt.Errorf(`Invalid upload provider in "%v": "%v"`, k, v.Provider)
		}

		err := replaceEnv(v.Config)
		if err != nil {
			return nil, err
		}
		mapstructure.Decode(v.Config, &uploader)
		err = uploader.Validate()
		if err != nil {
			return nil, fmt.Errorf(`Invalid uploader environment confuguration in "%v": %v`, k, err)
		}

		config.Uploaders[k] = uploader
	}

	for k, v := range raw.Jobs {

		if v.Cron == "" {
			return nil, fmt.Errorf(`Empty cron syntax in "%v": "%v"`, k, v.Cron)
		}

		preset, has := presets.Presets[v.Preset]
		if !has {
			return nil, fmt.Errorf("Invalid preset: %v", v.Preset)
		}

		err := replaceEnv(v.Config)
		if err != nil {
			return nil, err
		}
		mapstructure.Decode(v.Config, &preset)
		err = preset.Validate()
		if err != nil {
			return nil, fmt.Errorf(`Invalid preset configuration in "%v": %v`, k, err)
		}

		to, has := config.Uploaders[v.Upload.To]
		if !has {
			return nil, fmt.Errorf(`Uploader not defined in "%v": "%v"`, k, v.Upload.To)
		}

		upload_config := to.Config()
		err = replaceEnv(v.Upload.Config)
		if err != nil {
      return nil, fmt.Errorf("Unexpected error")
		}
		mapstructure.Decode(v.Upload.Config, &upload_config)
		err = to.ValidateConfig(upload_config)
		if err != nil {
			return nil, fmt.Errorf(`Invalid uploader configuration in "%v": %v`, k, err)
		}

		config.Jobs[k] = types.Job{
			Id:     k,
			Cron:   v.Cron,
			Preset: preset,
			Upload: types.ConfigJobUpload{
				To:       to,
				Filename: v.Upload.Filename,
				Config:   upload_config,
			},
		}
	}

	// fmt.Printf("config: %v\n", config)

	return &config, nil
}

func replaceEnv(conf any) error {
	config, ok := conf.(map[string]interface{})
	if !ok {
		return fmt.Errorf("Value is not of type map[string]interface{}")
	}

	for k, v := range config {
		trystring, ok := v.(string)
		if ok {
			config[k] = replaceTemplatedEnvString(trystring)
			continue
		}

		trymap, ok := v.(map[string]interface{})
		if ok {
			err := replaceEnv(trymap)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func replaceTemplatedEnvString(s string) string {
	res := ""
	pts1 := strings.Split(s, "${")

	for _, v := range pts1 {
		pts2 := strings.SplitN(v, "}", 2)

		if len(pts2) < 2 {
			res += pts2[0]
			continue
		}

		pts2[0] = os.Getenv(pts2[0])
		res += pts2[0] + pts2[1]
	}

	return res
}
