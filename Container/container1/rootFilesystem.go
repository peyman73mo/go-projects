package main

// https://man7.org/linux/man-pages/man2/pivot_root.2.html
// using pivot_root syscall to set a new root filesystem for the calling process (mount namespace/CLONE_NEWNS).

func pivot_root(newroot string) error {

	return nil
}
