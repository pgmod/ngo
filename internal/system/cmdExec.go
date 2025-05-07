package system

import "os/exec"

func Exec(cmd []string) (string, int) {
	if len(cmd) == 0 {
		return "", 1
	}
	command := exec.Command(cmd[0], cmd[1:]...)
	output, err := command.CombinedOutput()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return string(output), exitErr.ExitCode()
		}
		return string(output), 1
	}
	return string(output), 0
}
