package main

import (
	"fmt"
	"os"
	"runtime"
	"syscall"

	"golang.org/x/sys/unix"

	"k8s.io/kubernetes/cmd/kubelet/app"
)

func init() {
	// main needs to be locked on one thread and no go routines
	runtime.LockOSThread()
}

func main() {
	fmt.Fprintln(os.Stderr, app.NewKubeletCommand)

	value := &syscall.Rlimit{
		Cur: uint64(1000000000),
		Max: uint64(1000000000),
	}
	err := syscall.Setrlimit(unix.RLIMIT_AS, value)
	if err != nil {
		os.Exit(1)
	}

	fmt.Fprintln(os.Stderr, "hi")
}
