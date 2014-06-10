package infect

import (
    "os"
    "path"
)

var home string      = os.Getenv("HOME")
// var vimhome string   = path.Join(home, ".vim")
// var vimrc string     = path.Join(home, ".vimrc")
var vimhome string   = path.Join(home, "tmp", ".vim")
var vimrc string     = path.Join(home, "tmp", ".vimrc")
var startDir, err = os.Getwd()

var bundleDir string = path.Join(vimhome, "bundle")
