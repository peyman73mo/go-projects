// +buld linux
package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[0] {
	case "Run":
		Run()
	case "Child":
		Child()
	}

}

func Run() {
	fmt.Printf("[New-Container]\n")
	fmt.Printf("[Running %v]\n", os.Args[1])

	cmd := exec.Command("/proc/self/exe", append([]string{"Child"}, os.Args[1:]...)...)

	// map the child process's stdin/stdout/stderr to os.Stdin/os.Stdout/os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// *** cmd.SysProcAttr raises error if the code is not run in Linux ***
	/*
		***  https://man7.org/linux/man-pages/man7/namespaces.7.html
			" If the flags argument of the call specifies one or more of the
			CLONE_NEW* flags listed above, then new namespaces are
			created for each flag, and the child process is made a
			member of those namespaces."
		***
	*/
	// syscall.SysProcAttr allows us to set attributes on *exec.Cmd
	// each CLONE_NEW* adds new namespace to the process (UTS, PID, NS, USER, NET, IPC)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNS |
			syscall.CLONE_NEWUSER |
			syscall.CLONE_NEWNET |
			syscall.CLONE_NEWIPC,
	}

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func Child() {
	fmt.Printf("[Running %v]\n", os.Args[2])

	cmd := exec.Command(os.Args[2], os.Args[3:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// https://man7.org/linux/man-pages/man2/gethostname.2.html
	check(syscall.Sethostname([]byte("[new-container]$")))

	// https://man7.org/linux/man-pages/man2/chroot.2.html
	check(syscall.Chroot("ubuntu-fs/")) // mkdir /tmp/ubuntu-fs;mv ubuntu-fs /tmp/ubuntu-fs

	// https://man7.org/linux/man-pages/man2/chdir.2.html
	check(syscall.Chdir("/"))

	check(cmd.Run())
}

func check(err error) {
	panic(err)
}
