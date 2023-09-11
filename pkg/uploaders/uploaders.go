package uploaders

import (
	"github.com/mducoli/cronupper/pkg/types"
	"github.com/mducoli/cronupper/pkg/uploaders/s3"
)

var Uploaders = map[string]types.Uploader{
	"s3": s3.S3{},
}
