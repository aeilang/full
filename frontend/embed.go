package frontend

import (
	"embed"
	"io/fs"
)

//go:embed dist/*
var asserts embed.FS

func Assets() (fs.FS, error) {
	return fs.Sub(asserts, "dist")
}
