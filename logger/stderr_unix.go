// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package logger

import (
	"os"
	"syscall"
)

// redirectStderr to the file passed in
func redirectStderr(f *os.File) {
	err := syscall.Dup2(int(f.Fd()), int(os.Stderr.Fd()))
	if err != nil {
		Errorf("Failed to redirect stderr to file: %v", err)
	}
}
