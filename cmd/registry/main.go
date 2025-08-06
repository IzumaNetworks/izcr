package main

import (
	_ "net/http/pprof"

	"go.izuma.io/izcr/registry"
	_ "go.izuma.io/izcr/registry/auth/htpasswd"
	_ "go.izuma.io/izcr/registry/auth/silly"
	_ "go.izuma.io/izcr/registry/auth/token"
	_ "go.izuma.io/izcr/registry/proxy"
	_ "go.izuma.io/izcr/registry/storage/driver/azure"
	_ "go.izuma.io/izcr/registry/storage/driver/filesystem"
	_ "go.izuma.io/izcr/registry/storage/driver/gcs"
	_ "go.izuma.io/izcr/registry/storage/driver/inmemory"
	_ "go.izuma.io/izcr/registry/storage/driver/middleware/cloudfront"
	_ "go.izuma.io/izcr/registry/storage/driver/middleware/redirect"
	_ "go.izuma.io/izcr/registry/storage/driver/middleware/rewrite"
	_ "go.izuma.io/izcr/registry/storage/driver/s3-aws"
)

func main() {
	// NOTE(milosgajdos): if the only two commands registered
	// with registry.RootCmd fail they will halt the program
	// execution and  exit the program with non-zero exit code.
	// nolint:errcheck
	registry.RootCmd.Execute()
}
