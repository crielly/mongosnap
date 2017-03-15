package lvmsnap

import (
	"fmt"
	"os/exec"

	"github.com/crielly/mongosnap/logger"
)

// LvmSnap triggers an LVM snapshot
func LvmSnap(size, name, path string) {
	cmd := fmt.Sprintf("lvcreate -L %sM -s -n %s %s", size, name, path)
	fmt.Println(cmd)
	run := exec.Command("bash", "-c", cmd)
	stdoutStderr, err := run.CombinedOutput()
	logger.LogError(err)
	fmt.Printf("%s\n", stdoutStderr)
}
