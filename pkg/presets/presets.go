package presets

import (
	"github.com/mducoli/cronupper/pkg/presets/dockervolume"
	"github.com/mducoli/cronupper/pkg/presets/mongo"
	"github.com/mducoli/cronupper/pkg/presets/postgres"
	"github.com/mducoli/cronupper/pkg/types"
)

var Presets = map[string]types.Preset{
	"docker-volume": dockervolume.DockerVolume{},
	"postgres":      postgres.Postgres{},
	"mongo":         mongo.Mongo{},
}
