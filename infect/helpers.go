package infect

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"
    "regexp"
)

// func splitLine(s string) []string {
// 	arr := strings.SplitN(s, ":", 3)
// 	return arr
// }

func check(e error) {
	if e != nil {
		fmt.Printf("\n----\nError: %#v\n----\n", e)
		panic(e)
	}
}

func content(path string) string {
    c, e := ioutil.ReadFile(path)
    check(e)

    s := string(c)
    return s
}

func outputDir(repo string) string {
    r := strings.Split(repo, "/")
    return path.Join(bundleDir, r[1])
}

func split(s string) []string {
    return strings.Split(s, "\n")
}

func magicLines(lines []string) []string {

    r, _ := regexp.Compile("^\"=(.*)")
    results := make([]string, 0)
    for _, line := range lines {
        if r.MatchString(line) {
            cmdRepo := r.ReplaceAllString(line, "$1") // $1 is matched group
            results = append(results, cmdRepo)
        }
    }
    return results
}
