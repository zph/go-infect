package infect

import (
	"bufio"
	"fmt"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

const unverifiedDir bool = false
const verifiedDir bool = true

func install(_ *cli.Context) {

	s := content(vimrc())
	lines := split(s)
	mLines := magicLines(lines)
	bundleReqs := bundleRequests(mLines)
	dirs, e := ioutil.ReadDir(bundleDir)
	check(e)
	dirNames := make(map[string]bool)
	for _, dir := range dirs {
		dirNames[dir.Name()] = unverifiedDir
	}

	for _, line := range bundleReqs {
		dirNames[repoName(line)] = verifiedDir
		// add go-routines and cap worker pool at ... 20?
		processRepo(line)
	}

	deleteOldDirs(dirNames)
}

func deleteOldDirs(oldDirs map[string]bool) {
	deletes := make([]string, 0)
	for k, v := range oldDirs {
		if v == unverifiedDir {
			k = path.Join(bundleDir, k)
			k, err = filepath.Abs(k)
			check(err)
			response := askDelete(k)

			// TODO: add cmd
			// rm -rf each folder
			// syscall.Rmdir(k)
			if response {
				fmt.Printf("syscall.Rmdir(%s)", k)
			} else {
				fmt.Printf("noop(%s)", k)
			}

			// TODO: superfluous remove
			deletes = append(deletes, k)
		}
	}
}

func askDelete(line string) bool {

	fmt.Printf("Delete: %#v ? [y/N] => ", line)
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	m, err := regexp.MatchString("(?i)^([Y]+)", name)
	check(err)
	return m
}

func repoName(line string) string {
	arr := strings.SplitN(line, " ", 2)
	repo := strings.SplitN(arr[0], "/", 2)
	return repo[1]
}

func processRepo(line string) {
	// look for extra processing params
	arr := strings.SplitN(line, " ", 2)
	repo := arr[0]
	exists := dirExists(outputDir(repo))
	// fmt.Printf("Dir: %#v, %s\n", res, repo)
	if exists {
		// git pull
		output := outputDir(repo)
		os.Chdir(output)
		gitPull(repo)
		os.Chdir(startDir)
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
