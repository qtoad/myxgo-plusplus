package util

import (
	"fmt"

	"os/exec"
	"strconv"
	"strings"
)

/*
 * 运行命令
 *  */
func Command(cmd string, args ...string) (string, error) {
	command := exec.Command(cmd, args...)
	output, err := command.Output()
	return string(output), err
}

/*
 * 根据进程名称找到对应的pid
 *  */
func ProcessID(processName string) (int, error) {
	var pid int
	args := []string{
		"-ef",
		fmt.Sprintf("|grep %s", processName),
		"|grep -v grep",
		"|grep -v PPID",
		"|awk '{ print $2}'",
	}

	output, err := Command("ps", args...)
	if output != "" {
		output = strings.Replace(output, " ", "", -1)
		output = strings.Replace(output, "\n", "", -1)
		pid, err = strconv.Atoi(output)
	}

	return pid, err
}

/*
 * 运行指定的进程
 *  */
func ProcessStart(processName string) (string, error) {
	return Command("nohup", processName, "2>&1", "&")
}

/*
 * 杀死指定pid进程
 *  */
func ProcessStop(processName string) (string, error) {
	pid, err := ProcessID(processName)

	if err != nil {
		return "", err
	}

	return Command("kill", "-SIGTERM", fmt.Sprintf("%d", pid))
}
