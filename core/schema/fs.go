package schema

import _ "embed"

//go:embed git.graphqls
var Git string

//go:embed directory.graphqls
var Directory string

//go:embed file.graphqls
var File string

//go:embed container.graphqls
var Container string
