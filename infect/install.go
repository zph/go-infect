package infect

import (
	"github.com/codegangsta/cli"
    "strings"
    "fmt"
)

func install(_ *cli.Context){

    s := content(vimrc)
    lines := split(s)
    mLines := magicLines(lines)
    bundleReqs := bundleRequests(mLines)
    for _, line := range bundleReqs {
        // add go-routines and cap worker pool at ... 20?
        processRepo(line)
    }
}

func processRepo(line string) {
        // check each directory
        // either update or new clone
        // look for extra processing params
        arr := strings.SplitN(line, " ", 2)
        repo := arr[0]
        res := dirExists(outputDir(repo))
        // fmt.Printf("Dir: %#v, %s\n", res, repo)
        if res {
            // git pull
            gitPull(repo)
        } else {
            gitClone(repo)
        }
        if len(arr) > 1 {
            cmds := arr[1]
            if cmds != "" {
                fmt.Printf("Cmds: %s\n", cmds)
                // TODO: execute cmds
            }
        }
}

func bundleRequests(m []string) []string {
    bundles := make([]string, 0)
    for _, line := range m {
        arr := strings.SplitN(line, " ", 2)
        switch arr[0] {
        case "bundle":
            bundles = append(bundles, arr[1])
        default:
            // noop
            // what about other commands like shell??
        }
    }
    return bundles
}
