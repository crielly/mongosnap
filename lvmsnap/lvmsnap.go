package lvmsnap

import (
	"fmt"
	"os/exec"
)

func LvmSnap(size, name, path string) {
	cmd := fmt.Sprintf("lvcreate -L %sM -s -n %s %s", size, name, path)
	fmt.Println(cmd)
	run := exec.Command("bash", "-c", cmd)
	stdoutStderr, err := run.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", stdoutStderr)
}
