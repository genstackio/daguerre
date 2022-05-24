package assets

import (
	"embed"
)

//go:embed layer-types/*.json
var LayerTypes embed.FS

//go:embed images/*.png
var Images embed.FS
