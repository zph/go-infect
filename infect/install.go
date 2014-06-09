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
    fmt.Println("Results %#v", bundleReqs)
    for _, line := range bundleReqs {
        // check each directory
        // either update or new clone
        // look for extra processing params
        repo := strings.SplitN(line, " ", 2)
        gitPull(repo[0])
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
