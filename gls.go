package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"syscall"
)

func main() {
	dir, _ := os.Getwd()

	fileinfos, _ := ioutil.ReadDir(dir)

	for _, fileinfo := range fileinfos {
		sys := fileinfo.Sys()
		stat, ok := sys.(*syscall.Stat_t)
		if !ok {
			fmt.Errorf("syscall failed\n")
		}

		mode := fileinfo.Mode()
		filename := fileinfo.Name()
		modTime := fileinfo.ModTime()
		size := fileinfo.Size()

		var ownerName string
		owner, err := user.LookupId(fmt.Sprintf("%d", stat.Uid))
		if err != nil {
			ownerName = "-"
		} else {
			ownerName = owner.Username
		}

		fmt.Println(mode, modTime, ownerName, size, filename)
	}
}
