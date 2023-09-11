package types

type Uploader interface {
	Config() any
	Upload(filepath string, filename string, conf any) error
	Validate() error
	ValidateConfig(conf any) error
}
