package pklconf

import "embed"

//go:embed *.pkl **/*.pkl
var PklFiles embed.FS
