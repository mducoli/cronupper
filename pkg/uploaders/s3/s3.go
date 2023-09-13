package s3

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type S3Config struct {
	Bucket string `yaml:"bucket"`
	Prefix string `yaml:"prefix"`
}

type S3 struct {
	Endpoint   string `yaml:"endpoint"`
	Access_Key string `yaml:"access_key"`
	Secret_Key string `yaml:"secret_key"`
}

func (S3) ValidateConfig(conf any) error {
	var config, ok = conf.(S3Config)
	if !ok { //shoud probably never happen
		return fmt.Errorf("Unexpected error")
	}

	if config.Bucket == "" {
		return fmt.Errorf(`"bucket" value shoud not be empty`)
	}

	return nil
}

func (t S3) Validate() error {
	if t.Endpoint == "" {
		return fmt.Errorf(`"endpoint" value shoud not be empty`)
	}
	if t.Access_Key == "" {
		return fmt.Errorf(`"access_key" value shoud not be empty`)
	}
	if t.Secret_Key == "" {
		return fmt.Errorf(`"Secret_key" value shoud not be empty`)
	}

	return nil
}

func (S3) Config() any {
	var conf = S3Config{}
	return conf
}

func (t S3) Upload(path string, name string, c any) error {

	var conf, ok = c.(S3Config)
	if !ok {
		return fmt.Errorf("Unexpected error")
	}

	log.Printf(`Uploading to s3: "%v"`, conf.Prefix+name)

	ctx := context.Background() // TODO: understand wtf is this

	minioClient, err := minio.New(t.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(t.Access_Key, t.Secret_Key, ""),
		Secure: true,
	})
	if err != nil {
		return err
	}

  _, err = minioClient.FPutObject(ctx, conf.Bucket, conf.Prefix+name, path, minio.PutObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}
