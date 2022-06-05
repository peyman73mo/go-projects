//go:build linux
// +build linux

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/docker/docker/pkg/reexec"
)

// init() will be detected after the reexec.Command("namespaceInit")
func init() {
	reexec.Register("namespaceInit", namespaceInit)
	if reexec.Init() {
		os.Exit(0)
	}
}

func namespaceInit() {
	fmt.Printf("\n *** Namespace Init *** \n")
	namespaceRun()
}

func namespaceRun() {
	fmt.Printf("\n *** Namespace Run *** \n")

	cmd := exec.Command("/bin/bash")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// set the child process's command prompt to "[new-namespace]- # "
	cmd.Env = []string{"PS1=[new-namespace]\\$ "}

	must(cmd.Run())

}

func main() {
	// run /proc/self/exe with os.Args[0] set to namespaceInit
	// https://github.com/moby/moby/tree/master/pkg/reexec
	cmd := reexec.Command("namespaceInit")

	// map the child process's stdin/stdout/stderr to os.Stdin/os.Stdout/os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// *** cmd.SysProcAttr raises error if the code is not run in Linux ***

	// Using Namespace Clone() to create a new process
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
	// cmd.SysProcAttr = &syscall.SysProcAttr{
	// 	Cloneflags: syscall.CLONE_NEWUTS |
	// 		syscall.CLONE_NEWUSER |
	// 		syscall.CLONE_NEWPID |
	// 		syscall.CLONE_NEWNS |
	// 		syscall.CLONE_NEWNET |
	// 		syscall.CLONE_NEWIPC,
	// 	UidMappings: []syscall.SysProcIDMap{ // map the child process's UID to the parent process's UID
	// 		{
	// 			ContainerID: 0,
	// 			HostID:      os.Getuid(),
	// 			Size:        1,
	// 		},
	// 	},
	// 	GidMappings: []syscall.SysProcIDMap{ // map the child process's GID to the parent process's GID in new user namespace
	// 		{
	// 			ContainerID: 0,
	// 			HostID:      os.Getgid(),
	// 			Size:        1,
	// 		},
	// 	},
	// }

	must(cmd.Run())

}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
