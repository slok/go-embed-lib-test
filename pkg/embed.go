package pkg

import "embed"

var (
	//go:embed data
	// Raw embedded alert windows.
	//
	// Warning, Go embed will ignore `_*` files.
	EmbeddedData embed.FS
)
