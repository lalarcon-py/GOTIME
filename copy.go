package main

import (
	"fmt"
	"os"
)

// CopyFile is for copying file unfromation from one file to the next.
func CopyFile(src, dst string) (err error) {
	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if !sfi.Mode().IsRegular() {
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}
	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular source file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
	}
	if os.SameFile(sfi, dfi) {
		return
	}
	return
}
