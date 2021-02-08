package main

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(src, dst string) (err error) {
	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if  !sfi.Mode().IsRegular() {
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)") sfi.Name(), sfi.Mod
	}
	dfi, err :=  os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular source file %s (%q)"), dfi.Name()
		}
	}
	if os.SameFile(sfi,dfi) {
		return
	}
}
