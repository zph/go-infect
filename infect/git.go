package infect

import (
	"fmt"
	"os/exec"
	"strings"
)

func gitPull(repo string) bool {
	// output := outputDir(repo)
	// cd into dir && git pull -f
	cmd := git("pull", "-f")
	out := shell(cmd)
	fmt.Println(out)
	// TODO: actually pull
	return false
}

func gitClone(repo string) bool {
	output := outputDir(repo)
	remoteRepo := github(repo)
	args := fmt.Sprintf("%s %s", remoteRepo, output)
	cmd := git("clone", args)
	out := shell(cmd)
	fmt.Println(out)
	// TODO: actually clone
	return false
}

func github(repo string) string {
	return fmt.Sprintf("git@github.com:%s.git", repo)
}

func git(cmd string, args string) string {
	return fmt.Sprintf("git %s %s", cmd, args)
}

func shell(cmd string) string {

	cmds := strings.Split(cmd, " ")

	name := cmds[0]
	args := make([]string, 0)
	for _, arg := range cmds[1:] {
		args = append(args, arg)
	}

	cmdOut, err := exec.Command(name, args...).Output()
	check(err)

	return string(cmdOut)
}
