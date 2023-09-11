package types

type Config struct {
	Jobs      map[string]Job
	Uploaders map[string]Uploader
}

type Job struct {
	Id     string
	Cron   string
	Preset Preset
	Upload ConfigJobUpload
}

type ConfigJobUpload struct {
	To       Uploader
	Filename string
	Config   any
}
