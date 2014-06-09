package infect

import (
    "os"
    "path"
)

var home string      = os.Getenv("HOME")
var vimhome string   = path.Join(home, ".vim")
var vimrc string     = path.Join(home, ".vimrc")

var bundleDir string = path.Join(vimhome, "bundle")
