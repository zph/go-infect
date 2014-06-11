package infect

import (
	"os"
	"path"
)

func vimhome() string {
	var vimhome string
	if os.Getenv("INFECT_DEBUG") != "" {
		vimhome = path.Join(home, "tmp", ".vim")
	} else {
		vimhome = path.Join(home, ".vim")
	}
	return vimhome
}

func vimrc() string {
	var vimrc string
	if os.Getenv("INFECT_DEBUG") != "" {
		vimrc = path.Join(home, "tmp", ".vimrc")
	} else {
		vimrc = path.Join(home, ".vimrc")
	}
	return vimrc
}

var home string = os.Getenv("HOME")

var startDir, err = os.Getwd()

var bundleDir string = path.Join(vimhome(), "bundle")
