package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// GetCmdArgsByPid gets comand line arguments of the running process by PID
func GetCmdArgsByPid(pid int) ([]string, error) {
	f, err := os.Open(fmt.Sprintf("/proc/%d/cmdline", pid))
	defer f.Close()

	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(f)
	l, _, err := reader.ReadLine()
	if err != nil {
		return nil, err
	}

	s := strings.Replace(string(l), "\x00", " ", -1)
	ss := strings.Split(s, " ")

	return ss, nil
}
