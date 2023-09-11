package types

type Preset interface {
  Validate() error
	Run(file string) error
}
