package exec

import (
	"github.com/fs714/tunnel666/utils"
	"github.com/fs714/tunnel666/utils/log"
	"os/exec"
	"strings"
)

func ExecCommand(command string) (string, error) {
	var cmd *exec.Cmd
	if strings.Contains(command, "|") {
		cmd = exec.Command("bash", "-c", command)
	} else {
		cmd = exec.Command(utils.ToArgv(command)[0], utils.ToArgv(command)[1:]...)
	}
	log.Debug(strings.Join(cmd.Args, " "))
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Errorf("Failed to execute cmd: %s, out: %s, err: %v", strings.Join(cmd.Args, " "), string(out), err)
		return string(out), err
	}
	return string(out), nil
}
