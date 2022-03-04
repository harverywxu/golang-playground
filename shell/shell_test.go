package shell

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestShellInVariable(t *testing.T) {
	ts := fmt.Sprintf(`
m=%s
v=$(uname -r)
if [[ "$v" == "20.4.0" ]]; then
	echo $m
else
	echo no
fi
`, "hello world")
	cmd := exec.Command(ts)
	out, err := cmd.CombinedOutput()
	t.Logf("out: %v, error: %v", string(out), err)
}
