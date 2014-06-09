package infect

import (
    "fmt"
    "os/exec"
)

func gitPull(repo string) bool {
    output := outputDir(repo)
    // cd into dir && git pull -f
    cmd := git("pull", "-f")
    cmdFull := fmt.Sprintf("cd %s && %s", output, cmd)
    fmt.Println(cmdFull)
    // TODO: actually pull
    return false
}

func gitClone(repo string) bool {
    output := outputDir(repo)
    remoteRepo := github(repo)
    args := fmt.Sprintf("%s %s", remoteRepo, output)
    cmd := git("clone", args)
    fmt.Println(cmd)
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
    cmdOut, err := exec.Command(cmd).Output()
    check(err)

    return string(cmdOut)
}

